package Pool

import (
	"github.com/zeromicro/go-zero/core/logx"
	"nichebox/service/long-connection/rpc/internal/access/connection"
	"nichebox/service/long-connection/rpc/internal/access/protocol"
	"sync"
	"sync/atomic"
)

const (
	cap = 100
)

type Pool struct {
	cap         int32 // 容量
	workerCount int32 // worker数量

	taskHead *task      // task Pool 头指针
	taskTail *task      // task Pool 尾指针
	taskLock sync.Mutex // task lock
}

func NewPool() *Pool {
	return &Pool{
		cap:         cap,
		workerCount: 0,
	}
}

type worker struct {
	pool *Pool // 指向Pool，需要获取task
}

type task struct {
	f      func(packet *protocol.Packet, conn connection.LongConn, logger logx.Logger) // 要执行的函数
	packet *protocol.Packet
	conn   connection.LongConn
	logger logx.Logger

	next *task // 指向下一个task
}

func (p *Pool) Go(f func(packet *protocol.Packet, conn connection.LongConn, logger logx.Logger), packet *protocol.Packet, conn connection.LongConn, logger logx.Logger) {
	t := &task{f: f, packet: packet, conn: conn, logger: logger}
	p.taskLock.Lock()
	if p.taskHead == nil {
		p.taskHead = t
		p.taskTail = t
	} else {
		p.taskTail.next = t
		p.taskTail = t
	}
	p.taskLock.Unlock()

	// 创建新worker
	if atomic.LoadInt32(&p.workerCount) < p.cap {
		atomic.AddInt32(&p.workerCount, 1)
		w := &worker{pool: p}
		w.run() // run 待实现
	}
}

func (w *worker) run() {
	go func() {
		defer func() {
			if r := recover(); r != nil {

			}
		}()

		for {
			var t *task
			w.pool.taskLock.Lock()

			// 从 task Pool 获取 task
			if w.pool.taskHead != nil {
				t = w.pool.taskHead
				w.pool.taskHead = t.next
			}

			// 如果没有任何task，则关闭该worker
			if t == nil {
				atomic.AddInt32(&w.pool.workerCount, -1)
				w.pool.taskLock.Unlock()
				return
			}
			w.pool.taskLock.Unlock()

			// 执行函数
			t.f(t.packet, t.conn, t.logger)
		}
	}()
}
