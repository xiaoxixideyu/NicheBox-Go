package routes

import (
	"nichebox/service/long-connection/rpc/internal/access/connection"
	"strconv"
	"strings"
	"sync"
)

const Plus = "+"

type Router struct {
	mUidToConns   map[int64][]connection.LongConn
	mConnIDToConn map[string]connection.LongConn
	mAddrToConnID map[string]string

	mu sync.Mutex
}

var RouterIns Router
var RouterOnce sync.Once

func GetRouter() *Router {
	RouterOnce.Do(func() {
		RouterIns = Router{
			mUidToConns:   map[int64][]connection.LongConn{},
			mConnIDToConn: map[string]connection.LongConn{},
			mAddrToConnID: map[string]string{},
			mu:            sync.Mutex{},
		}
	})
	return &RouterIns
}

func (r *Router) GetConnsByUid(uid int64) ([]connection.LongConn, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	conns, ok := r.mUidToConns[uid]
	return conns, ok
}

func (r *Router) RegisterAddr(addr string, uid int64, userAgent string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	connID := strconv.FormatInt(uid, 10) + Plus + userAgent
	r.mAddrToConnID[addr] = connID
}

func (r *Router) GetUidAndUAByAddr(addr string) (int64, string, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	connID, ok := r.mAddrToConnID[addr]
	if !ok {
		return 0, "", false
	}
	split := strings.Split(connID, Plus)
	uid, _ := strconv.ParseInt(split[0], 10, 64)
	ua := split[1]
	return uid, ua, true
}

func (r *Router) RegisterConn(uid int64, userAgent string, conn connection.LongConn) {
	r.mu.Lock()
	defer r.mu.Unlock()

	connID := strconv.FormatInt(uid, 10) + Plus + userAgent
	oldConn, needDelete := r.mConnIDToConn[connID]
	r.mConnIDToConn[connID] = conn

	conns, ok := r.mUidToConns[uid]
	if !ok {
		s := make([]connection.LongConn, 0, 5)
		s = append(s, conn)
		r.mUidToConns[uid] = s

	} else {
		if needDelete {
			for i, c := range conns {
				if c == oldConn {
					conns = append(conns[:i], conns[i+1:]...)
					break
				}
			}
		}
		conns = append(conns, conn)
		r.mUidToConns[uid] = conns

	}
}

func (r *Router) UnregisterConn(conn connection.LongConn) {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := strconv.FormatInt(conn.GetUid(), 10) + Plus + conn.GetUserAgent()
	delete(r.mConnIDToConn, key)

	conns := r.mUidToConns[conn.GetUid()]
	for i, c := range conns {
		if c == conn {
			conns = append(conns[:i], conns[i+1:]...)
			break
		}
	}
}
