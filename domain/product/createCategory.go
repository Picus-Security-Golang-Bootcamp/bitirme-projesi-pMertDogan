package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
)

func CreateProduct(c *gin.Context) {

	var product Product
	var responseModel domain.ResponseModel

	if err := c.Bind(&product); err != nil {
		responseModel.ErrMsg = err.Error()
		responseModel.ErrDsc = "Error while binding request body"
		responseModel.ResponseCode = http.StatusBadRequest
		c.JSON(responseModel.ResponseCode, responseModel)
		return
	}

}
