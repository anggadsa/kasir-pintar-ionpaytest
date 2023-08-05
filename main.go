package main

import (
	// "kasir-pintar-ionpaytest/database"
	"kasir-pintar-ionpaytest/routes"
	// "net/http"
	// "github.com/gin-gonic/gin"
	// "github.com/rs/zerolog/log"
	// "github.com/spf13/viper"
)

func main() {
	// database.Migrate()
	// database.Seeder()
	routes.Run()

	// log.Info().Msg("Started Server!")
	// r := gin.Default()
	// viper.SetConfigFile(".env")
	// viper.ReadInConfig()

	// db := database.Migrate()

	// // Provide db variable to controllers
	// r.Use(func(c *gin.Context) {
	// 	c.Set("db", db)
	// 	c.Next()
	// })

	// r.GET("", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, "Hallo")
	// })

	// r.Run()
}
