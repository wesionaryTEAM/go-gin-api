package service

import (
	"go-gin-api/entity"
)

// VideoService is ...
type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

// create a constructor function to create new instances of VideoService ...
// Return implementation of interface

//New is ...
func New() VideoService {
	return &videoService{}
}

// Append new videos to slice...

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

// return slice of videos ...

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
