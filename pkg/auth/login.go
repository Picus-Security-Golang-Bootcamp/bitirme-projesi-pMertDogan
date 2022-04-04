package appAuth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gin-gonic/gin"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/user"
	customCrypto "github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/crypto"
	jwtUtils "github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/jwt"
)

func (a *authHandler) login(c *gin.Context) {
	var req user.LoginRequestDTO
	var res user.ResponseModel
	//extract user from request
	if err := c.Bind(&req); err != nil {
		res.ErrMsg = "Your request body is not valid. Please check your request body. In Body email and password are required."
		res.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, res)
	}
	//Verify is user exist with the same hash
	user ,err := user.Repo().CheckIsUserExistWithThisEmail(req.Email)

	if err != nil {
		res.ErrMsg = "Something went wrong. Please try again later."
		res.ResponseCode = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if user == nil {
		res.ErrMsg = "This email is not registered."
		res.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, res)
		return
	}
	
	//Verify is password correct
	isItCorrect := customCrypto.CheckPasswordHash(req.Password, user.Password)

	//if password is not correct return error
	if !isItCorrect {
		res.ErrMsg = "Password is not correct."
		res.ResponseCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, res)
		return
	}

	//create JWT claims
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"email":  user.Email,
		"iat":    time.Now().Unix(),
		// "iss":    os.Getenv("ENV"),
		"exp":    time.Now().Add(1 * time.Hour).Unix(),
		"isAdmin":  user.IsAdmin,
	})
	//create JWT token
	token := jwtUtils.GenerateToken(jwtClaims, a.cfg.JWTConfig.SecretKey)
	res.AccesToken = token
	c.JSON(http.StatusOK, res)
}
