package product

import (
	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/config"
	jwtUtils "github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/jwt"
)

func ProductControllerDef(router *gin.Engine, cfg *config.Config) {

	//https://github.com/gin-gonic/gin#using-middleware
	//Use JWT verification middleware

	product := router.Group("/product")

	product.POST("/", jwtUtils.JWTAdminMiddleware(cfg.JWTConfig.SecretKey, cfg.JWTConfig.AccesTokenLifeMinute), CreateProduct)
	// cat.GET("/", GetAllCategoriesWithPagination)
	
}