package main

import (
	"ST/dal"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	dal.InitDB()
	config.InitGen()
	routers.InitRouter(r)
	r.Run()
}
