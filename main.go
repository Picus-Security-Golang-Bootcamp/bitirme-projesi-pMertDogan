package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/database"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/category"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/check"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/user"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/auth"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/config"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/graceful"
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
	user.UserRepoInit(db)
	//Migrate Structure
	category.Repo().Migrations()
	user.Repo().Migrations()
	user.Repo().CreateAdminIfNotExist(cfg)

	category.CategoryControllerDef(router,cfg)
	check.CheckControllerDef(router)
	appAuth.AuthHandler(router,cfg)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.ServerConfig.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.ServerConfig.ReadTimeoutSecs * int64(time.Second)),
		WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeoutSecs * int64(time.Second)),
	}

	//Graceful Shutdown & Restart
	//https://grisha.org/blog/2014/06/03/graceful-restart-in-golang/
	//https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	
	log.Println("Basket Service started")
	graceful.ShutdownGin(srv, time.Duration(cfg.ServerConfig.TimeoutSecs*int64(time.Second)))
}
