package cast

type Film struct {
	Title    string `json:"title"`
	Year     string `json:"year"`
	Runtime  string `json:"runtime"`
	Director string `json:"director"`
	Country  string `json:"country"`
	Genre    string `json:"genre"`
	Plot     string `json:"plot"`
	Poster   string `json:"poster"`
}

func NewFilm(title, year, runtime, director, country, genre, plot, poster string) *Film {
	return &Film{
		Title:    title,
		Year:     year,
		Runtime:  runtime,
		Director: director,
		Country:  country,
		Genre:    genre,
		Plot:     plot,
		Poster:   poster,
	}
}
