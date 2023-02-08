package main

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	pprof.Register(r)
	bakend := r.Group("/test")
	{
		bakend.GET("a", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "This is a")
		})
		bakend.GET("b", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "This is b")
		})
	}

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
