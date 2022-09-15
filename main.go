package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"work-hour-warn/conf"
	"work-hour-warn/controller"
	_ "work-hour-warn/docs" //执行swag init生成的docs文件夹路径 _引包表示只执行init函数
)

// @title           工时预警系统API
// @version         1.0
// @description     这是一个工时预警系统
// @host      localhost:8080
func main() {
	conf.InitMysql()

	r := gin.Default()

	r.GET("/lazyBoys", controller.Warn)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
