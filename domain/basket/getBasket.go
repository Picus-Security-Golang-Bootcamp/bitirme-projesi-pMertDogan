package basket

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
)

func GetBasket(c *gin.Context) {

	response := domain.ResponseModel{}

	//get userID from url
	userID := c.Param("id")

	//checked on MW side but we can add addtionale check here
	if userID == "" {
		response.ErrMsg = "userID is required"
		response.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//userID to int
	userIDInt, err := strconv.Atoi(userID)

	if err != nil {
		response.ErrMsg = "userID must be integer"
		response.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, response)
		return

	}

	//add to Basket
	v,err := Repo().GetBasketsByUserID(userIDInt)

	if err != nil {
		response.ErrMsg = "error getting basket"
		response.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//return success
	response.ResponseCode = http.StatusOK
	response.Data = v
	c.JSON(http.StatusOK, response)

}
