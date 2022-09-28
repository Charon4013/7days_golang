package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.GET("/", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "Hello Pan\n")
	})

	// 主动触发数组越界来测试Recovery()
	r.GET("/panic", func(ctx *gee.Context) {
		names := []string{"Pan"}
		ctx.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
