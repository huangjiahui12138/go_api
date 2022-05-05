package main

import (
	"github.com/gin-gonic/gin"
	// "jzapi/configs"
	"jzapi/routes"
)

func main() {
	//gin.SetMode(gin.ReleaseMode) // 默认为 debug 模式，设置为发布模式
	engine := gin.Default()
	router.InitRouter(engine)
	engine.Run(":8000")
}