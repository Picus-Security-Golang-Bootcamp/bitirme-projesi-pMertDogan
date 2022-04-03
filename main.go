package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/database"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/category"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/utils"
	"github.com/spf13/viper"
)

func init() {

	//Init Zap logger
	utils.InitializeLogger()
	defer utils.Logger.Sync() // flushes buffer, if any

	//https://github.com/spf13/viper
	viper.SetConfigName("config") // name of config file (without extension)
	// viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	// viper.AddConfigPath("/etc/appname/")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			utils.Logger.Fatal("Config file not found")

		} else {
			// Config file was found but another error was produced
			utils.Logger.Fatal("Config file found but another error was produced")
		}
	}
	// Config file found and successfully parsed
}

func main() {
	//Init Gin
	router := gin.Default()

	//init database
	db, err := database.ConnectPostgresDB()
	if err != nil {
		utils.Logger.Fatal("cannot connect to database")
	}

	//init Repos
	category.CategoryRepoInit(db)

	//Migrate Structure
	category.Repo().Migrations()

	category.CategoryControllerDef(router)

	router.Run(":8080")
}
