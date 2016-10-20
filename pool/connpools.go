package pool

import (
	"sync"
)

type ConnPools struct {
	sync.RWMutex
	pools map[string]*ConnPool
}

func (cps *ConnPools) Has(instance string) bool {
	cps.RLock()
	_, has := cps.pools[instance]
	cps.RUnlock()
	return has
}

func (cps *ConnPools) Get(instance string) (*ConnPool, bool) {
	cps.RLock()
	p, has := cps.pools[instance]
	cps.RUnlock()
	return p, has
}

func (cps *ConnPools) Put(connPool *ConnPool) {
	cps.Lock()
	cps.pools[connPool.Addr] = connPool
	cps.Unlock()
}

func (cps *ConnPools) Size() int {
	cps.RLock()
	l := len(cps.pools)
	cps.RUnlock()
	return l
}

func (cps *ConnPools) Keys() []string {
	i := 0
	cps.RLock()
	keys := make([]string, len(cps.pools))
	for key := range cps.pools {
		keys[i] = key
	}
	cps.RUnlock()
	return keys
}

func (cps *ConnPools) Evict(instance string) {
	cps.Lock()
	p, has := cps.pools[instance]
	if !has {
		cps.Unlock()
		return
	}
	delete(cps.pools, instance)
	cps.Unlock()
	p.Clean()
}
