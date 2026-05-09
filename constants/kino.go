package constants

import "github.com/Prague-Kino/cast/cast"

const AerokinaProgramURL = "https://www.kinoaero.cz/en?cinema=1%2C2%2C9%2C7%2C3&sort=sort-by-data&hall=34%2C35%2C1%2C2%2C3%2C24&english-friendly=1"

var (
	KinoAero = cast.Kino{
		Name:        "Kino Aero",
		Address:     "Biskupcova 31, Praha 3",
		Description: "The cult repertory cinema with a single screen but a lot going on it.",
		BaseDomain:  "https://www.kinoaero.cz/en",
		ProgramURL:  AerokinaProgramURL,
	}

	KinoSvetozor = cast.Kino{
		Name:        "Kino Světozor",
		Address:     "Vodičkova 41, Praha 1",
		Description: "A three screen arthouse cinema in the very heart of Prague.",
		BaseDomain:  "https://www.kinosvetozor.cz/en",
		ProgramURL:  AerokinaProgramURL,
	}

	KinoLucerna = cast.Kino{
		Name:        "Kino Lucerna",
		Address:     "Vodičkova 36, Prague 1",
		Description: "A legendary cinema in Prague with a grand and festive interior.",
		BaseDomain:  "https://www.kinolucerna.cz/en",
		ProgramURL:  AerokinaProgramURL,
	}

	KinoPritomnost = cast.Kino{
		Name:        "Kino Přítomnost",
		Address:     "Cinema Siwiecova 1, Praha 3",
		Description: "Prague's first boutique cinema. Where the film experience meets the bar experience.",
		BaseDomain:  "https://www.kinopritomnost.cz/en",
		ProgramURL:  AerokinaProgramURL,
	}

	EdisonFilmhub = cast.Kino{
		Name:        "Edison Filmhub",
		BaseDomain:  "https://edisonfilmhub.cz/en/",
		ProgramURL:  "https://edisonfilmhub.cz/en/programme",
		Address:     "Jeruzalémská 1321/2, 110 00 Nové Město, Praha 1",
		Description: "You can find us in the Edison Transformation station, one of Prague's most significant functionalist buildings. The building is part of Prague's heritage reservation and is protected by UNESCO.",
	}

	KinoPonrepo = cast.Kino{
		Name:        "Ponrepo",
		BaseDomain:  "nfa.cz",
		ProgramURL:  "https://nfa.cz/en/ponrepo-cinema/program/program",
		Address:     "Bartolomějská 291/11, 110 00 Staré Město",
		Description: "We are a memory institution. Our mission is to care for film heritage, to facilitate its discovery and to help the development of the Czech audiovisual industry and film culture.",
	}
)

var AllKinos = []cast.Kino{
	KinoAero, KinoSvetozor, KinoLucerna, KinoPritomnost, EdisonFilmhub, KinoPonrepo,
}
