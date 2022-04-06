package appAuth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
	user, err := user.Repo().CheckIsUserExistWithThisEmail(req.Email)

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
	jwtClaimsAccess := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"email":  user.Email,
		"iat":    time.Now().Unix(), //issued at current time
		// "iss":    os.Getenv("ENV"),
		"exp":            time.Now().Add(time.Duration(a.cfg.JWTConfig.AccesTokenLifeMinute) * time.Minute).Unix(), //expiration time is one hour
		"isAdmin":        user.IsAdmin,
		"isItAccesToken": true,
	})

	//generate JWT refresh token
	jwtClaimsRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"email":  user.Email,
		"iat":    time.Now().Unix(), //issued at current time
		// "iss":    os.Getenv("ENV"),
		"exp":            time.Now().Add(time.Duration(a.cfg.JWTConfig.RefreshTokenLifeMinute) * time.Hour).Unix(), //expiration time is one hour
		"isAdmin":        user.IsAdmin,
		"isItAccesToken": false,
	})

	//create JWT token

	accesToken := jwtUtils.GenerateToken(jwtClaimsAccess, a.cfg.JWTConfig.SecretKey)
	refreshToken := jwtUtils.GenerateToken(jwtClaimsRefresh, a.cfg.JWTConfig.SecretKey)
	res.AccesToken = accesToken
	res.RefreshToken = refreshToken
	res.ResponseCode = http.StatusOK
	c.JSON(http.StatusOK, res)
}
