package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/database"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/category"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/check"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/config"
	logger "github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/logging"
	"go.uber.org/zap"
)

func main() {

	//Load Config with depency injection
	cfg, err := config.LoadConfig("config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	// Set global logger
	logger.NewLogger(cfg)
	defer logger.Close()

	//init database
	db := database.ConnectPostgresDB(cfg)
	if err != nil {
		zap.L().Fatal("cannot connect to database")
	}

	//Init Gin
	router := gin.Default()
	//init Repos
	category.CategoryRepoInit(db)

	//Migrate Structure
	category.Repo().Migrations()

	category.CategoryControllerDef(router)
	check.CheckControllerDef(router)
	router.Run(":8080")
}
