package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Duration    int      `json:"duration"`
	Genres      []Genre  `json:"genres" gorm:"many2many:movie_genres"`
	Artists     []Artist `json:"artists" gorm:"many2many:movie_artists"`
}

type Genre struct {
	gorm.Model
	Name string `json:"genre"`
}

type Artist struct {
	gorm.Model
	Name string `json:"name"`
}
