package routes

import (
	"kasir-pintar-ionpaytest/controllers"

	"github.com/gin-gonic/gin"
)

func Registration(r *gin.Engine) *gin.Engine {
	v2 := r.Group("/api/v2")
	{
		registration := new(controllers.Registration)
		v2.GET("/registrations", registration.Index)
		v2.POST("/registrations", registration.Store)
	}

	return r
}
