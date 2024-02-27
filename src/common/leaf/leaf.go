package leaf

import (
	"errors"
	"nichebox/common/leaf/model"
	"nichebox/common/leaf/model/mysql"
	"sync"

	"gorm.io/gorm"
)

type Leaf struct {
	bizTag      string
	maxId       int64
	cur         int64
	isAvailable bool
	isLoading   bool
	waitChan    chan error
	mutex       sync.Mutex
}

type Leaves struct {
	leafInterface model.LeafInterface
	leafMap       map[string]*Leaf
	cacheMap      map[string]*Leaf
	mutex         sync.Mutex
}

func NewLeavesMysql(database, username, password, host, port string, maxIdleConns, maxOpenConns, connMaxLifeTime int) (*Leaves, error) {
	leafInterface, error := mysql.NewMysqlInterface(database, username, password, host, port, maxIdleConns, maxOpenConns, connMaxLifeTime)
	if error != nil {
		return nil, error
	}
	return &Leaves{
		leafInterface: leafInterface,
		leafMap:       make(map[string]*Leaf),
		cacheMap:      make(map[string]*Leaf),
		mutex:         sync.Mutex{},
	}, nil
}

func (leaves *Leaves) createLeaf(bizTag string) error {
	leafModel := &model.Leaf{
		BizTag: bizTag,
		MaxId:  0,
	}
	if err := leaves.leafInterface.CreateLeafTX(leafModel); err != nil {
		return err
	}
	return nil
}

func (leaves *Leaves) Next(bizTag string, step int, factor float64) (int64, error) {
	if step <= 0 {
		return 0, errors.New("step less or equal than 0")
	}

	var leaf *Leaf
	getLeaf := func() error {
		leaves.mutex.Lock()
		defer leaves.mutex.Unlock()
		leaf = leaves.leafMap[bizTag]
		if leaf == nil {
			start, end, err := leaves.leafInterface.NextTX(bizTag, step)
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return err
				}
				if err = leaves.createLeaf(bizTag); err != nil {
					return err
				}
				start, end, err = leaves.leafInterface.NextTX(bizTag, step)
				if err != nil {
					return err
				}
			}
			leaf = &Leaf{
				bizTag:      bizTag,
				maxId:       end,
				cur:         start,
				isAvailable: true,
				isLoading:   false,
				waitChan:    make(chan error, 1),
				mutex:       sync.Mutex{},
			}
			leaves.leafMap[bizTag] = leaf
		}
		return nil
	}

	if err := getLeaf(); err != nil {
		return 0, err
	}

	leaf.mutex.Lock()

	for !leaf.isAvailable {
		leaf.mutex.Unlock()
		leaves.mutex.Lock()
		leaf = leaves.leafMap[bizTag]
		leaves.mutex.Unlock()
		leaf.mutex.Lock()
	}
	if leaf.cur > leaf.maxId {
		leaves.mutex.Lock()
		cache := leaves.cacheMap[bizTag]
		leaves.mutex.Unlock()

		if cache == nil {
			if !leaf.isLoading {
				leaf.isLoading = true
				go leaves.nextCahce(leaf, bizTag, step)
			}
			if err := <-leaf.waitChan; err != nil {
				leaf.mutex.Unlock()
				return 0, err
			}
			leaves.mutex.Lock()
			cache = leaves.cacheMap[bizTag]
			leaves.mutex.Unlock()
		}

		leaf.isAvailable = false

		leaves.mutex.Lock()
		leaf.mutex.Unlock()
		leaf = cache
		leaf.mutex.Lock()
		leaves.leafMap[bizTag] = leaf
		leaves.cacheMap[bizTag] = nil
		leaves.mutex.Unlock()
	} else if float64(leaf.maxId-leaf.cur)/float64(step) < float64(1-factor) && !leaf.isLoading {
		leaves.mutex.Lock()
		if leaves.cacheMap[bizTag] == nil {
			leaf.isLoading = true
			go leaves.nextCahce(leaf, bizTag, step)
		}
		leaves.mutex.Unlock()
	}

	id := leaf.cur
	leaf.cur++
	leaf.mutex.Unlock()

	return id, nil
}

func (leaves *Leaves) nextCahce(curLeaf *Leaf, bizTag string, step int) {
	defer func() {
		curLeaf.mutex.Lock()
		defer curLeaf.mutex.Unlock()
		curLeaf.isLoading = false
	}()

	start, end, err := leaves.leafInterface.NextTX(bizTag, step)
	if err != nil {
		curLeaf.waitChan <- err
		return
	}

	cache := &Leaf{
		bizTag:      bizTag,
		maxId:       end,
		cur:         start,
		isAvailable: true,
		isLoading:   false,
		waitChan:    make(chan error, 1),
		mutex:       sync.Mutex{},
	}

	leaves.mutex.Lock()
	leaves.cacheMap[bizTag] = cache
	leaves.mutex.Unlock()
	curLeaf.waitChan <- nil
}
