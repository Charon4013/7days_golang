package geecache

import (
	"fmt"
	"sync"
)

// Group is a cache namespace and associate data loaded spread over
type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

// 接口型函数

// Getter loads data for a key
type Getter interface {
	Get(key string) ([]byte, error)
}

// GetterFunc implement Getter with a function
type GetterFunc func(key string) ([]byte, error)

// Get implement Getter interface function
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// NewGroup create a new instance of Group
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter!")
	}
	mu.Lock()
	defer mu.Unlock()

	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

// GetGroup return the named group created with NewGroup
// return nil means no such a group
func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}

func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}

	// if key in mainCache, then return cache data
	if v, ok := g.mainCache.get(key); ok {
		return v, nil
	}

	// if not, call load
	return g.load(key)
}

// at distributed scene, load will get data from getFromPeer, if failed, it will get from getLocally
func (g *Group) load(key string) (value ByteView, err error) {
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (ByteView, error) {
	// get source data
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	// copy to the cache
	value := ByteView{b: cloneBytes(bytes)}
	g.populateCache(key, value)
	return value, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
