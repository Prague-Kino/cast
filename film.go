package main

type Film struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Plot     string `json:"Plot"`
	Poster   string `json:"Poster"`
}

func NewFilm(title, year, rated, released, runtime, genre, director, plot, poster string) *Film {
	return &Film{
		Title:    title,
		Year:     year,
		Rated:    rated,
		Released: released,
		Runtime:  runtime,
		Genre:    genre,
		Director: director,
		Plot:     plot,
		Poster:   poster,
	}
}
