package basket

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product"
)

func AddToBasket(c *gin.Context) {

	response := domain.ResponseModel{}

	//get userID from url
	userID := c.Param("id")

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

	var req AddToBasketDTO
	//bind to dto
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ResponseCode = http.StatusBadRequest
		response.ErrMsg = "error binding json "
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//check if product has enought stock
	productQuantity, err := product.Repo().GetProductQuantityById(req.ProductID)

	if err != nil {
		response.ResponseCode = http.StatusBadRequest
		response.ErrMsg = "product not found"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if productQuantity < req.TotalQuantity {
		response.ResponseCode = http.StatusBadRequest
		response.ErrMsg = "not enought stock"
		response.ErrDsc = "product quantity is " + strconv.Itoa(productQuantity)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	//add to Basket
	err = Repo().CreateBasket(userIDInt, req.ProductID, req.TotalQuantity)

	if err != nil {
		response.ResponseCode = http.StatusBadRequest
		response.ErrMsg = "Unable add to basket"
		response.ErrDsc = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	v , err := Repo().GetBasketsByUserID(userIDInt)
	//return success
	response.ResponseCode = http.StatusOK
	response.Data = v
	c.JSON(http.StatusOK, response)



}
