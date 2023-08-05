package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	"github.com/spf13/viper"
)

func Db() *gorm.DB {
	// Enable VIPER to read Environment Variables
	// viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	// To get the value from the config file using key
	// viper package read .env
	viper_user := viper.Get("PG_USER")
	viper_password := viper.Get("PG_PASSWORD")
	viper_db := viper.Get("PG_DB_NAME")
	viper_host := viper.Get("PG_HOST")
	viper_port := viper.Get("PG_PORT")
	fmt.Println(viper_db)
	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viper_host, viper_port, viper_user, viper_db, viper_password)
	// fmt.Println("conname is\t\t", prosgret_conname)
	db, err := gorm.Open("postgres", prosgret_conname)
	if err != nil {
		panic("Failed to connect to database!")
	}

	// db.AutoMigrate(&Registration{})
	return db
}
