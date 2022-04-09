package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"go.uber.org/zap"
)

func Search(c *gin.Context) {
	//Json
	c.Header("Content-Type", "application/json")

	//get query params
		//get query params
		pageSize := c.DefaultQuery("pageSize", "10")
		page := c.DefaultQuery("page", "1")
	

	var responseModel domain.ResponseModel

	//get query params
	searchText := c.Query("search")

	if searchText == "" {
		responseModel.ErrMsg = "search text is empty"
		responseModel.ResponseCode = http.StatusBadRequest
		c.JSON(responseModel.ResponseCode, responseModel)
		zap.L().Debug("Search", zap.String("searchText", searchText))
		return
	}

	v, err := Repo().SearchProducts(searchText,page,pageSize)

	if err != nil {
		responseModel.ErrMsg = "error getting products"
		responseModel.ErrDsc = err.Error()
		responseModel.ResponseCode = http.StatusInternalServerError
		c.JSON(responseModel.ResponseCode, responseModel)
		zap.L().Debug("Getting products error ", zap.String("searchText", searchText))
		return
	}

	responseModel.Data = v
	responseModel.ResponseCode = http.StatusOK
	responseModel.PageSize = strconv.Itoa(len(v))
	c.JSON(responseModel.ResponseCode, responseModel)

}
