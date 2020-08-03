package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"searcher/controller"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/search", controller.GetESRes) //搜索接口
	r.POST("/health", controller.PostApi) //健康检查接口

	//r.Use(GlobalMiddleware)
	_ = r.Run(":9366")
}
