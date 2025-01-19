package models

import "time"

type Consumption struct {
	HotWater   int       `json:"hotwater"`
	ColdWater  int       `json:"coldwater"`
	Heating    int       `json:"heating"`
	Cooling    int       `json:"cooling"`
	LastUpdate time.Time `json:"lastupdate"`
}

type Action string

const (
	Heating   Action = "heating"
	HotWater  Action = "hotwater"
	ColdWater Action = "coldwater"
)
