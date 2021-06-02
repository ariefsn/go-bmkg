package models

import (
	"time"
)

type WindDirection struct {
	ID          string
	Type        string
	Description string
	Data        []WindDirectionValue
}

type WindDirectionValue struct {
	Type     string
	H        string
	Datetime time.Time
	Degree   struct {
		Unit  string
		Value float64
	}
	Card WindCardValue
	Sexa struct {
		Unit  string
		Value float64
	}
}

type WindCardValue struct {
	Unit  string
	Value string
}

func (c *WindCardValue) ValueText() string {
	m := map[string]string{
		"N":        "North",
		"NNE":      "North-Northeast",
		"NE":       "Northeast",
		"ENE":      "East-Northeast",
		"E":        "East",
		"ESE":      "East-Southeast",
		"SE":       "Southeast",
		"SSE":      "South-Southeast",
		"S":        "South",
		"SSW":      "South-Southwest",
		"SW":       "Southwest",
		"WSW":      "West-Southwest",
		"W":        "West",
		"WNW":      "West-Northwest",
		"NW":       "Northwest",
		"NNW":      "North-Northwest",
		"VARIABLE": "VARIABLE",
	}

	return m[c.Value]
}
