package models


type Movies struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Year int `json:"year"`
	ReleaseDate time.Time `json:"re"`
}

type Genre struct {

}

type MovieGenre struct {

}