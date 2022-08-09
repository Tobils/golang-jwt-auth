package service

import (
	"fmt"
	"go-auth-jwt/models"
)

type AlbumService interface {
	Find() []models.Album
}

type logInformation struct {
	message string
}

func AlbumHandler() AlbumService {
	return &logInformation{
		message: "find all album",
	}
}

func (f *logInformation) Find() []models.Album {
	fmt.Println(f.message)
	var albums = []models.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrame", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan", Artist: "Sarah Vaughan", Price: 39.99},
	}

	fmt.Println(albums)
	return albums
}
