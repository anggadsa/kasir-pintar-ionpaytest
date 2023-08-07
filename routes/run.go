package routes

import (
	"kasir-pintar-ionpaytest/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	database.Migrate() //Migrate Database Models
	// database.Seeder() //Seed Data Dummy to Database

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"v2": "http://localhost:8080/api/v2",
		})
	})

	Route(r)

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hallo")
	})
	return r
}

func REGISTRATION(r *gin.Engine) {
	panic("unimplemented")
}

func Run() *gin.Engine {
	r := SetupRouter()
	r.Run()

	return r
}
