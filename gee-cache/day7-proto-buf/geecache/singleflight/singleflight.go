package singleflight

import "sync"

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex // mu用来保护m
	m  map[string]*call
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok { // 如果有key在Group中存在
		g.mu.Unlock()
		c.wg.Wait()         // 等待
		return c.val, c.err // 返回值
	}
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c // 将key添加进Group
	g.mu.Unlock()

	c.val, c.err = fn() // 调用fn, 发起请求
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key) // 更新g.m
	g.mu.Unlock()

	return c.val, c.err // 返回值
}
