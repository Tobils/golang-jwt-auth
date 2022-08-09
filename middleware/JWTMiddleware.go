package middleware

import (
	"fmt"
	"go-auth-jwt/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		autHHeader := c.GetHeader("Authorization")
		tokenString := autHHeader[len(BEARER_SCHEMA):]
		// Clean up the value since it could contain line breaks
		tokenString = strings.TrimSpace(tokenString)
		token, err := service.JWTAuthService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
			fmt.Println(claims["email"])
			fmt.Printf("Type claims %T\n", claims)
			// c.Set("email", claims["email"])
			// c.Set("user", "suhada")
			// c.Next()
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
