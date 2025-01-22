package models

import "time"

type Consumption struct {
	HotWater   int       `json:"hotwater"`
	ColdWater  int       `json:"coldwater"`
	Heating    int       `json:"heating"`
	Cooling    int       `json:"cooling"`
	LastUpdate time.Time `json:"lastupdate"`
	Year       int       `json:"year"`
	Month      int       `json:"month"`
}

type Action string

const (
	Heating   Action = "heating"
	HotWater  Action = "hotwater"
	ColdWater Action = "coldwater"
)

type Filter string

const (
	Day   Filter = "day"
	Week  Filter = "week"
	Month Filter = "month"
	Year  Filter = "year"
)
