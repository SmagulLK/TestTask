package service

import (
	"TestTask/models"
	"TestTask/repository"
)

type Service struct {
	VideoServiceIn
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		VideoServiceIn: NewVideoService(repository.VideoReposIn),
	}
}

type VideoServiceIn interface {
	GetVideos() ([]models.Videos, error)
}
