package router

import (
	"github.com/gin-gonic/gin"
	"github.com/samridhi-sahu/cms/controller"
)

func RestRoutes(router *gin.Engine) {
	router.GET("/getCustomers", controller.Logging, controller.GetCustomers)
	router.GET("/getCustomer", controller.Logging, controller.GetACustomer)
	router.POST("/addCustomer", controller.Logging, controller.AddCustomer)
	router.PUT("/updateCustomer", controller.Logging, controller.UpdateACustomer)
	router.DELETE("/deleteCustomer", controller.Logging, controller.DeleteACustomer)
}
