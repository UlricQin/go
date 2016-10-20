package pool

import (
	"errors"
	"io"
	"sync"
)

// ErrMaxConn Maximum connections reached
var ErrMaxConn = errors.New("maximum connections reached")

// ConnPool manages the life cycle of connections
type ConnPool struct {
	sync.RWMutex

	// NewConn is used to create a new connection if necessary.
	NewConn func() (io.Closer, error)

	Addr string
	Max  int
	Idle int

	active int
	free   []io.Closer
	conns  map[io.Closer]struct{}
}

// NewConnPool create a connection pool
func NewConnPool(addr string, max int, idle int) *ConnPool {
	return &ConnPool{
		Addr:  addr,
		Max:   max,
		Idle:  idle,
		conns: make(map[io.Closer]struct{}),
	}
}

// Get get a connection
func (cp *ConnPool) Get() (conn io.Closer, err error) {
	conn = cp.tryFree()
	if conn != nil {
		return
	}

	if cp.overMax() {
		return nil, ErrMaxConn
	}

	conn, err = cp.NewConn()
	if err == nil {
		cp.inc()
		cp.conns[conn] = struct{}{}
	}

	return
}

// Clean close all connections
func (cp *ConnPool) Clean() {
	cp.Lock()
	for conn := range cp.conns {
		if conn != nil {
			conn.Close()
		}
	}
	cp.active = 0
	cp.conns = nil
	cp.Unlock()
}

// Put recycle the connection
func (cp *ConnPool) Put(conn io.Closer) {
	if cp.overIdle() {
		cp.Close(conn)
	} else {
		cp.Lock()
		cp.free = append(cp.free, conn)
		cp.Unlock()
	}
}

// Close close the connection
func (cp *ConnPool) Close(conn io.Closer) {
	cp.dec()
	if conn != nil {
		conn.Close()
	}
}

func (cp *ConnPool) tryFree() io.Closer {
	cp.Lock()
	if len(cp.free) == 0 {
		cp.Unlock()
		return nil
	}

	conn := cp.free[0]
	cp.free = cp.free[1:]

	cp.Unlock()
	return conn
}

func (cp *ConnPool) overMax() bool {
	cp.RLock()
	over := cp.active >= cp.Max
	cp.RUnlock()
	return over
}

func (cp *ConnPool) overIdle() bool {
	cp.RLock()
	over := len(cp.free) >= cp.Idle
	cp.RUnlock()
	return over
}

func (cp *ConnPool) inc() {
	cp.Lock()
	cp.active++
	cp.Unlock()
}

func (cp *ConnPool) dec() {
	cp.Lock()
	cp.active--
	cp.Unlock()
}
