package repository

import (
	"TestTask/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type VideosRepository struct {
	db *sqlx.DB
}

func NewVideosRepository(db *sqlx.DB) *VideosRepository {
	return &VideosRepository{db: db}
}

func (Video *VideosRepository) GetVideos() ([]models.Videos, error) {
	query := "SELECT id,title, link FROM Video"

	rows, err := Video.db.Query(query)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer rows.Close()

	var videos []models.Videos
	for rows.Next() {
		var video models.Videos
		err := rows.Scan(&video.Id, &video.Title, &video.Link)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}
		videos = append(videos, video)
	}

	if err = rows.Err(); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return videos, nil
}
