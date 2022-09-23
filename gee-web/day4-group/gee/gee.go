package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	*RouterGroup // store all groups
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // support middleware
	router      *router
}

func New() *Engine {
	engine := &Engine{}
	engine.RouterGroup = &RouterGroup{router: newRouter()}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {

	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		router: group.router,
	}
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
