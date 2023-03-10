package main

import (
	"gim/initialize"
	"gim/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello world.",
	})
}

func main() {
	logConf := initialize.DefaultLogger()
	//初始化日志
	initialize.InitLogger(logConf)
	//初始化配置
	//initialize.InitConfig()
	//初始化数据库
	initialize.InitDB()
	initialize.InitRedis()

	router := router.Router()
	router.Run(":8080")
}
