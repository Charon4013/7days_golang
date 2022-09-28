package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func trace(message string) string {
	var pcs [32]uintptr

	// Callers返回调用栈的程序计数器, 第0个是callers本身, 第1个是上层trace, 第2个是再上一层的defer func
	// 跳过前三个是为了简化日志
	n := runtime.Callers(3, pcs[:]) // callers本身

	var str strings.Builder
	str.WriteString(message + "\nTraceback: ")
	for _, pc := range pcs[:n] {
		// 获取对应的函数
		fn := runtime.FuncForPC(pc)
		// 获取函数的文件名和行号
		file, line := fn.FileLine(pc)
		// 打印到日志
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recovery() HandlerFunc {
	return func(ctx *Context) {
		// 再上一层的defer func
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				// 上层trace
				log.Printf("%s\n\n", trace(message))
				ctx.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		ctx.Next()
	}
}
