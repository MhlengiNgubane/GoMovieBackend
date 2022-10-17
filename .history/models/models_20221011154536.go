package models

import "time"

type Movies struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Year        int       `json:"year"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime     int       `json:"runtime"`
	Rating      int       `json:"rating"`
	MPAARating  string    `json:"mpaa_rating`
	
}

type Genre struct {
}

type MovieGenre struct {
}