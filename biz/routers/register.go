package routers

import (
	"ST/biz/routers/employees"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	employees.Register(r)
}
