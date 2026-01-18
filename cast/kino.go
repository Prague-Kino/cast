package cast

type Kino struct {
	Name       string `json:"Name"`
	BaseDomain string `json:"BaseDomain"`
	ProgramURL string `json:"ProgramURL"`
}

func NewKino(name, baseDomain, programURL string) *Kino {
	return &Kino{
		Name:       name,
		BaseDomain: baseDomain,
		ProgramURL: programURL,
	}
}
