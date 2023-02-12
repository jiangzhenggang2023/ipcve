package routers

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	pprof.Register(router)

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "HelloWorld,这是后端模块")
	})

	// API
	// 生成token
	// admin/gentoken

	// 下载修复功能文件
	// -H {"token":""} -X GET admin/getUpdateScript?autoAction=False

	// 远程触发修复
	//

	return router
}
