package routes

import (
	"kasir-pintar-ionpaytest/controllers"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) *gin.Engine {
	v2 := r.Group("/api/v2")
	{
		registration := new(controllers.Registration)
		v2.POST("/registrations", registration.Store)
	}
	{
		payment := new(controllers.Payment)
		v2.POST("/payment", payment.Show)
	}
	{
		inquiry := new(controllers.Inquiry)
		v2.POST("/inquiry", inquiry.Show)
	}
	return r
}
