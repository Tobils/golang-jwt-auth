package main

import (
	"fmt"
	"go-auth-jwt/controller"
	"go-auth-jwt/middleware"
	"go-auth-jwt/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var authService service.AuthService = service.StaticAuthService()
	var jwtService service.JWTService = service.JWTAuthService()
	var authController controller.AuthController = controller.AuthHandler(authService, jwtService)

	var albumService service.AlbumService = service.AlbumHandler()
	var albumController controller.AlbumController = controller.AlbumHandler(albumService)

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	auth := router.Group("/auth")
	{
		auth.POST("/login", func(ctx *gin.Context) {
			token := authController.Login(ctx)
			if token != "" {
				ctx.JSON(http.StatusOK, gin.H{
					"token": token,
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, nil)
			}
		})
	}

	authorized := router.Group("/")
	authorized.Use(middleware.AuthorizeJWT())
	{
		authorized.GET("/album", func(ctx *gin.Context) {
			album := albumController.Find(ctx)
			fmt.Println(album)
			ctx.JSON(http.StatusOK, gin.H{
				"album": album,
			})
		})
	}

	port := "8080"
	router.Run(":" + port)
}
