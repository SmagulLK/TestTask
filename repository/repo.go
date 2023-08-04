package repository

import (
	"TestTask/models"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	VideoReposIn
}

//var _ VideoReposIn = (*Repository)(nil)

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		VideoReposIn: NewVideosRepository(db),
	}
}

type VideoReposIn interface {
	GetVideos() ([]models.Videos, error)
}
