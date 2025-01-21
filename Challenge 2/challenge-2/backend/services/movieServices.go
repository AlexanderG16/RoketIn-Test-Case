package services

import (
	"challenge-2/backend/models"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func CreateMovie(db *gorm.DB, movie *models.Movie) error {
	return db.Create(movie).Error
}

func UpdateMovie(db *gorm.DB, movie *models.Movie, id uint) error {
	movie.ID = id
	return db.Save(movie).Error
}

func GetAllMovies(db *gorm.DB, page int) ([]models.Movie, error) {
	var movies []models.Movie
	err := db.Model(&models.Movie{}).Preload("Genres").Preload("Artists").Limit(10).Offset((page - 1) * 10).Find(&movies).Error
	return movies, err
}

func SearchMovies(db *gorm.DB, title []string, description []string, artists_names []string, genres_names []string) ([]models.Movie, error) {
	var movies []models.Movie
	var artists []string
	var genres []string

	if len(artists_names) > 0 {
		artists = artists_names
	} else {
		db.Table("artists").Select("name").Find(&artists)
	}
	if len(genres_names) > 0 {
		genres = genres_names
	} else {
		db.Table("genres").Select("name").Find(&genres)
	}

	fmt.Println(artists)
	fmt.Println(genres)

	err := db.Preload("Genres", "name IN ?", genres).Preload("Artists", "name IN ?", artists).Where("title LIKE ?", ("%"+strings.Join(title[:], "")+"%")).Where("description LIKE ?", ("%" + strings.Join(description[:], "") + "%")).Find(&movies).Error

	return movies, err
}

// func SearchMoviesbyTitle(db *gorm.DB, title []string) ([]models.Movie, error) {
// 	var movies []models.Movie
// 	err := db.mo
// }
