package jwtUtils

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"net/http"
)

func JWTAuthMiddleware(secretKey string) gin.HandlerFunc {

	return func(c *gin.Context) {
		zap.L().Info("JWTAuthMiddleware is triggered")
		if c.GetHeader("Authorization") != "" {
			decodedClaims := VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodedClaims != nil {
				//Only admins can acces this route with acces token
				if decodedClaims.IsAdmin &&  decodedClaims.IsItAccesToken {
					c.Next()
					c.Abort()
					return
				}
			}

			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized!"})
		}
		c.Abort()
		return
	}
}
