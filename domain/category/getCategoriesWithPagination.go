package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategoriesWithPagination(c *gin.Context) {

	response := ResponseModel{}
	//get query params
	pageSize := c.DefaultQuery("pageSize", "10")
	pageNo := c.DefaultQuery("pageNo", "1")

	//get all categories with pagination
	v, err := Repo().GetAllCategoriesWithPagination(pageNo, pageSize)

	if err != nil {
		response.ResponseCode = http.StatusInternalServerError
		response.ErrMsg = "error getting  categories with pagination"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.ResponseCode = http.StatusOK
	response.Data = v
	c.JSON(http.StatusOK, response)

}
