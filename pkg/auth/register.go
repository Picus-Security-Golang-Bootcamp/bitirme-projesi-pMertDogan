package appAuth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/user"
)

func (a *authHandler) register(c *gin.Context) {

	var req user.RegisterRequestDTO
	var res user.ResponseModel
	//extract user from request
	if err := c.Bind(&req); err != nil {
		res.ErrMsg = "Your request body is not valid. Please check your request body. In Body email and password are required."
		res.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, res)
		return
	}

	//check is user exist with the same email adress
	 userR, err := user.Repo().CheckIsUserExistWithThisEmail(req.Email)
	//isUserExist return true if user exist on DB :)
	if err != nil {
		res.ErrMsg = "Something went wrong. Please try again later."
		res.ResponseCode = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if userR != nil {
		res.ErrMsg = "This email is already registered."
		res.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//we can register user
	err = user.Repo().RegisterUser(req.Email, req.Password)

	if err != nil {
		res.ErrMsg = "Something went wrong. Please try again later."
		res.ResponseCode = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res.ResponseCode = http.StatusCreated

	c.JSON(http.StatusOK, res)

}
