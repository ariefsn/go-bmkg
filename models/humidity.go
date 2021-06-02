package models

import "time"

type Humidity struct {
	ID          string
	Type        string
	Description string
	Data        []HumidityValue
}

type HumidityValue struct {
	Type     string
	H        string
	Datetime time.Time
	Value    float64
	Unit     string
}

type HumidityMinMax struct {
	ID          string
	Type        string
	Description string
	Data        []HumidityMinMaxValue
}

type HumidityMinMaxValue struct {
	Type     string
	Day      time.Time
	Datetime time.Time
	Value    float64
	Unit     string
}
