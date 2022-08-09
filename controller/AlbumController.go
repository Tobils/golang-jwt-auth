package controller

import (
	"go-auth-jwt/models"
	"go-auth-jwt/service"

	"github.com/gin-gonic/gin"
)

type AlbumController interface {
	Find(ctx *gin.Context) []models.Album
}

type albumController struct {
	albumService service.AlbumService
}

func AlbumHandler(albumService service.AlbumService) AlbumController {
	return &albumController{
		albumService: albumService,
	}
}

func (controller *albumController) Find(ctx *gin.Context) []models.Album {
	return controller.albumService.Find()
}
