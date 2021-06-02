package models

import "time"

type Weather struct {
	ID          string
	Type        string
	Description string
	Data        []WeatherValue
}

type WeatherValue struct {
	Type     string
	H        string
	Datetime time.Time
	Value    float64
	Unit     string
}

func (w *WeatherValue) ValueText(lang string) string {
	m := map[float64]map[string]string{
		0: {
			"id": "Cerah",
			"en": "Clear Skies",
		},
		1: {
			"id": "Cerah Berawan",
			"en": "Partly Cloudy",
		},
		2: {
			"id": "Partly Cloudy",
			"en": "",
		},
		3: {
			"id": "Berawan",
			"en": "Mostly Cloudy",
		},
		4: {
			"id": "Berawan Tebal",
			"en": "Overcast",
		},
		5: {
			"id": "Udara Kabur",
			"en": "Haze",
		},
		10: {
			"id": "Asap",
			"en": "Smoke",
		},
		45: {
			"id": "Kabut",
			"en": "Fog",
		},
		60: {
			"id": "Hujan Ringan",
			"en": "Light Rain",
		},
		61: {
			"id": "Hujan Sedang",
			"en": "Rain",
		},
		63: {
			"id": "Hujan Lebat",
			"en": "Heavy Rain",
		},
		80: {
			"id": "Hujan Lokal",
			"en": "Isolated Shower",
		},
		95: {
			"id": "Hujan Petir",
			"en": "Severe Thunderstorm",
		},
		97: {
			"id": "Hujan Petir",
			"en": "Severe Thunderstorm",
		},
	}

	return m[w.Value][lang]
}
