package models

import "time"

type Consumption struct {
	Heating    int       `json:"heating"`
	LastUpdate time.Time `json:"lastupdate"`
}
