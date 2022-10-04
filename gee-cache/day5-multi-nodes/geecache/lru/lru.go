package lru

import (
	"container/list"
)

// Cache is a LRU cache. It is not safe for concurrent access
type Cache struct {
	// max memory
	maxBytes int64
	// current used memory
	nbytes int64
	// double linked list
	ll    *list.List
	cache map[string]*list.Element
	// callback function(when someone removed)
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

// Value use Len() to count how many bytes it takes
type Value interface {
	Len() int
}

// New is the constructor of Cache
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	// if the key of node exist, move to front, then return values
	if ele, ok := c.cache[key]; ok {
		// assume the front node is the last list node
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOldest removes the oldest item
func (c *Cache) RemoveOldest() {
	// get the first node, then remove it
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		// remove the node's mapping relation
		delete(c.cache, kv.key)
		// update current used memory capacity
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			// callback function if OnEvicted is not nil
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Add adds a value to the Cache
func (c *Cache) Add(key string, value Value) {
	// if the key exist
	if ele, ok := c.cache[key]; ok {
		// move the node to list end
		c.ll.MoveToFront(ele)
		// get old kv position
		kv := ele.Value.(*entry)
		// update memory
		c.nbytes += int64(len(kv.key)) - int64(value.Len())
		// update key's value
		kv.value = value
	} else {
		// add new node to list end
		ele := c.ll.PushFront(&entry{key, value})
		// update cache
		c.cache[key] = ele
		// update memory
		c.nbytes += int64(len(key)) + int64(value.Len())
		//fmt.Println("-current memory: ", c.Len(), "<= int64key: ", int64(len(key)), "int64value: ", int64(value.Len()))
	}
	// remove lru nodes if current memory is larger than maxBytes
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
