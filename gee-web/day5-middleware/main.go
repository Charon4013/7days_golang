package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello World!</h1>\n")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(ctx *gee.Context) {
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
		})
	}
	r.Run(":9999")
}

func onlyForV2() gee.HandlerFunc {
	return func(ctx *gee.Context) {
		t := time.Now()
		ctx.Fail(500, "Internal Serve Error")
		log.Printf("[%d] %s in %v for group v2", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}
