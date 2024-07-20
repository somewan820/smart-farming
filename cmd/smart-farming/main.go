package main

import (
	"ST/biz/config"
	"ST/biz/dal"
	"ST/biz/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	dal.InitDB()
	config.InitGen()
	routers.InitRouter(r)
	r.Run()
}
