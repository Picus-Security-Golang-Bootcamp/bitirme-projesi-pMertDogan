package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/config"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/jwt"

	"go.uber.org/zap"
)

func CategoryControllerDef(router *gin.Engine, cfg *config.Config) {

	c := router.Group("/category")
	//https://github.com/gin-gonic/gin#using-middleware
	//Use JWT verification middleware
	c.Use(jwtUtils.JWTAdminMiddleware(cfg.JWTConfig.SecretKey,cfg.JWTConfig.AccesTokenLifeMinute))
	{
		UploadCategorysFromCSV(router, c)

		// c.POST("/submit", submitEndpoint)
		// c.POST("/read", readEndpoint)
	}

}

// https://gin-gonic.com/docs/examples/upload-file/single-file/
func UploadCategorysFromCSV(router *gin.Engine, c *gin.RouterGroup) {

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	c.POST("/upload", func(c *gin.Context) {

		response := ResponseModel{}

		// single-file upload read categoryFile
		file, _ := c.FormFile("categoryFile")

		//IF user not provide file
		if file == nil {
			zap.L().Error("categoryFile is nil")
			response.ResponseCode = http.StatusBadRequest
			response.ErrMsg = "categoryFile is nil"
			c.JSON(http.StatusBadRequest, response)
			return
		}

		zap.L().Info("Uploaded succes", zap.String("fileName ", file.Filename))

		//send file to our converter function
		// https://stackoverflow.com/questions/40956103/how-to-convert-multipart-fileheader-file-type-to-os-file-in-golang
		//TLDR; we dont need :)
		csvCategories, err := CategoryFromCSV(file)
		if err != nil {
			response.ErrMsg = "check the csv  file. " + err.Error()
			response.ResponseCode = http.StatusBadRequest
			zap.L().Error("check the csv  file. " + err.Error())
			c.JSON(http.StatusBadRequest, response)
			return
		}

		//call category repo and save it to db
		Repo().CreateCategories(csvCategories)

		v, err := Repo().GetAllCategoriesWithLimit(50)
		if err != nil {
			zap.L().Error("error getting all categories", zap.Error(err))
			response.ResponseCode = http.StatusInternalServerError
			response.ErrMsg = "error getting all categories"
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response.Data = v
		c.JSON(http.StatusOK, response)
	})
}
