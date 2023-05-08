package api

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"sync"
)

type Cache struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CacheInMemory struct {
	caches map[int]Cache
	mutex  sync.Mutex
	count  int
}

func NewCacheInMemory() CacheInMemory {
	return CacheInMemory{
		caches: caches,
		count:  len(caches),
	}
}

var (
	caches       = make(map[int]Cache)
	defaultCache CacheInMemory
)

func init() {
	for i := 1; i < 100; i++ {
		uid := uuid.New()
		caches[i] = Cache{
			Id:   i,
			Name: fmt.Sprintf("%s_%d", uid.String(), i),
		}
	}
	defaultCache = NewCacheInMemory()
}

func (c *CacheInMemory) GetCache(id int) (Cache, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	v, ok := caches[id]
	if ok {
		return v, nil
	}
	return Cache{}, errors.New(fmt.Sprintf("get cache fail %d", id))
}

func (c *CacheInMemory) DeleteCache(id int) (bool, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, ok := caches[id]
	if ok {
		delete(caches, id)
		return true, nil
	}
	return false, errors.New(fmt.Sprintf("delete cache fail %d", id))
}

func (c *CacheInMemory) PutCache(id int, name string) (Cache, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	v, ok := caches[id]
	if !ok {
		return Cache{}, errors.New(fmt.Sprintf("put cache fail %d", id))
	}
	v.Name = name
	caches[id] = v
	return v, nil
}

func (c *CacheInMemory) PostCache(name string) (Cache, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
	c.caches[c.count] = Cache{
		Id:   c.count,
		Name: name,
	}
	return c.caches[c.count], nil
}
