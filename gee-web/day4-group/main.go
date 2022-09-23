package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello World!</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(ctx *gee.Context) {
			ctx.HTML(http.StatusOK, "<h1>Hello v1</h1>")
		})

		v1.GET("/hello", func(ctx *gee.Context) {
			// expect /hello?name=helloGeektutu
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(ctx *gee.Context) {
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
		})

		v2.POST("/login", func(ctx *gee.Context) {
			// expect /hello?name=helloGeektutu
			ctx.JSON(http.StatusOK, gee.H{
				"username": ctx.PostForm("username"),
				"password": ctx.PostForm("password"),
			})
		})
	}

	r.Run(":9999")
}
