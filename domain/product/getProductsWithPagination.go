package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
)

func GetAllProductWithPagination(c *gin.Context) {

	response := domain.ResponseModel{}
	//get query params
	pageSize := c.DefaultQuery("pageSize", "10")
	pageNo := c.DefaultQuery("pageNo", "1")

	

	//get all categories with pagination
	v, err := Repo().GetAllWithPagination(pageNo, pageSize)
	
	if err != nil {
		response.ResponseCode = http.StatusInternalServerError
		response.ErrMsg = "error getting  categories with pagination"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//No data found
	if len(v) == 0 {
		response.ResponseCode = http.StatusNotFound
		response.ErrMsg = "no data found"
		c.JSON(http.StatusNotFound, response)
		return
	}

	response.ResponseCode = http.StatusOK
	response.Data = v
	response.PageNo = pageNo
	response.PageSize = pageSize
	
	c.JSON(http.StatusOK, response)

}


