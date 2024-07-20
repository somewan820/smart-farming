package employees

import (
	"ST/biz/handler/employees"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	root := r.Group("/employees")
	{
		//root.POST("/login", employees.Login)
		//root.POST("/register", employees.Register)
		//root.GET("/verify/code", employees.GetVerifyCode)
		affiliation := root.Group("/affiliation")
		{
			//affiliation.PUT("", employees.EditAffiliationInfo)
			affiliation.GET("", employees.GetAffiliationInfo)
			//affiliation.POST("", employees.CreateAffiliation)
		}
		//root.PUT("/basic", employees.EditBasicInfo)
		//root.GET("", employees.GetEmployeeInfo)
	}
}
