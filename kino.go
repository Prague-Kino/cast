package main

type Kino struct {
	Name        string `json:"name"`
	BaseDomain  string `json:"website"`
	ProgramURL  string `json:"ProgramURL"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

func NewKino(name, baseDomain, programURL string) *Kino {
	return &Kino{
		Name:       name,
		BaseDomain: baseDomain,
		ProgramURL: programURL,
	}
}
