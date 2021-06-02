package models

import "time"

type Temperature struct {
	ID          string
	Type        string
	Description string
	Data        []TemperatureValue
}

type TemperatureValue struct {
	Type       string
	H          string
	Datetime   time.Time
	Celcius    float64
	Fahrenheit float64
}

type TemperatureMinMax struct {
	ID          string
	Type        string
	Description string
	Data        []TemperatureMinMaxValue
}

type TemperatureMinMaxValue struct {
	Type       string
	Day        time.Time
	Datetime   time.Time
	Celcius    float64
	Fahrenheit float64
}
