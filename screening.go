package main

import (
	"fmt"
	"time"
)

type Screening struct {
	Film  Film      `json:"Film"`
	Kino  string    `json:"Kino"`
	Date  time.Time `json:"Date"`
	Time  string    `json:"Time"`
	Price int       `json:"Price"`
}

func NewScreening(film Film, kino string, date time.Time, time string, price int) Screening {
	return Screening{
		Film:  film,
		Kino:  kino,
		Date:  date,
		Time:  time,
		Price: price,
	}
}

func (s Screening) String() string {
	return fmt.Sprintf(
		"{ %-10s %5s | %-25s @ %-6s for %4d Kč }\n",
		s.Kino,
		s.Date.Format("02/01"),
		s.Film.Title,
		s.Time,
		s.Price,
	)
}
