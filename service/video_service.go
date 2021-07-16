package service

import "apigolang/entity"

type VideoServiceInterface interface {
	FindAll() []entity.Video
	Save(entity.Video) entity.Video
}

type VideoService struct {
	videos []entity.Video
}

func (service *VideoService) FindAll() []entity.Video {
	return service.videos
}

func (service *VideoService) Save(entidad entity.Video) entity.Video {
	service.videos = append(service.videos, entidad)
	return entidad
}
