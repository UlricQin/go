package pool

import (
	"io"
	"net"
	"net/rpc"
	"time"
)

func NewGobRPCPool(addr string, ct time.Duration, max, idle int) *ConnPool {
	p := NewConnPool(addr, max, idle)

	p.NewConn = func() (io.Closer, error) {
		conn, err := net.DialTimeout("tcp", p.Addr, ct)
		if err != nil {
			return nil, err
		}
		return rpc.NewClient(conn), nil
	}

	return p
}
