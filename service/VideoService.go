package service

import (
	"TestTask/models"
	"TestTask/repository"
)

type VideoService struct {
	repo repository.VideoReposIn
}

func NewVideoService(repo repository.VideoReposIn) *VideoService {
	return &VideoService{repo: repo}
}
func (s *VideoService) GetVideos() ([]models.Videos, error) {
	videos, err := s.repo.GetVideos()
	if err != nil {
		return nil, err
	}
	return videos, nil
}
