package models

import "time"

type WindSpeed struct {
	ID          string
	Type        string
	Description string
	Data        []WindSpeedValue
}

type WindSpeedValue struct {
	Type     string
	H        string
	Datetime time.Time
	Knot     struct {
		Unit  string
		Value float64
	}
	Mph struct {
		Unit  string
		Value float64
	}
	Kph struct {
		Unit  string
		Value float64
	}
	Ms struct {
		Unit  string
		Value float64
	}
}
