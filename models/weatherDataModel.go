package models

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
)

func parseDateTime(datetime, t string) time.Time {
	layout := "200601021504"

	switch t {
	case "timestamp":
		layout = "20060102150405"
	case "day":
		layout = "20060102"
	}

	dt, _ := time.Parse(layout, datetime)
	return dt
}

type WeatherModel struct {
	XMLName          xml.Name      `xml:"data"`
	Source           string        `xml:"source,attr"`
	ProductionCenter string        `xml:"productioncenter,attr"`
	Forecast         ForecastModel `xml:"forecast"`
}

type ForecastModel struct {
	XMLName xml.Name    `xml:"forecast"`
	Domain  string      `xml:"domain,attr"`
	Issue   IssueModel  `xml:"issue"`
	Areas   []AreaModel `xml:"area"`
}

type IssueModel struct {
	XMLName xml.Name `xml:"issue"`
	Year    string   `xml:"year"`
	Month   string   `xml:"month"`
	Day     string   `xml:"day"`
	Hour    string   `xml:"hour"`
	Minute  string   `xml:"minute"`
	Second  string   `xml:"second"`
}

func (i *IssueModel) Timestamp() time.Time {
	d := fmt.Sprintf("%s%s%s%s%s%s", i.Year, i.Month, i.Day, i.Hour, i.Minute, i.Second)
	return parseDateTime(d, "timestamp")
}

type AreaModel struct {
	XMLName     xml.Name         `xml:"area"`
	ID          string           `xml:"id,attr"`
	Latitude    string           `xml:"latitude,attr"`
	Longitude   string           `xml:"longitude,attr"`
	Type        string           `xml:"type,attr"`
	Region      string           `xml:"region,attr"`
	Level       string           `xml:"level,attr"`
	Description string           `xml:"description,attr"`
	Domain      string           `xml:"domain,attr"`
	Tags        string           `xml:"tags,attr"`
	Name        string           `xml:"name"`
	Parameters  []ParameterModel `xml:"parameter"`
}

type ParameterModel struct {
	XMLName     xml.Name         `xml:"parameter"`
	ID          string           `xml:"id,attr"`
	Description string           `xml:"description,attr"`
	Type        string           `xml:"type,attr"`
	Parameters  []TimerangeModel `xml:"timerange"`
}

type TimerangeModel struct {
	XMLName  xml.Name              `xml:"timerange"`
	Type     string                `xml:"type,attr"`
	H        string                `xml:"h,attr"`
	Day      string                `xml:"day,attr"`
	Datetime string                `xml:"datetime,attr"`
	Values   []TimerangeValueModel `xml:"value"`
}

type TimerangeValueModel struct {
	XMLName xml.Name `xml:"value"`
	Unit    string   `xml:"unit,attr"`
	Value   string   `xml:",chardata"`
}

func (a *AreaModel) getSpecificParameters(id string) ParameterModel {
	res := ParameterModel{}

	for _, v := range a.Parameters {
		if v.ID == id {
			res = v
			break
		}
	}

	return res
}

func (a *AreaModel) Humidity() Humidity {
	res := a.getSpecificParameters("hu")
	values := []HumidityValue{}

	for _, v := range res.Parameters {
		val, _ := strconv.ParseFloat(v.Values[0].Value, 64)

		values = append(values, HumidityValue{
			Type:     v.Type,
			H:        v.H,
			Datetime: parseDateTime(v.Datetime, "datetime"),
			Value:    val,
			Unit:     v.Values[0].Unit,
		})
	}

	return Humidity{
		ID:          res.ID,
		Description: res.Description,
		Type:        res.Type,
		Data:        values,
	}
}

func (a *AreaModel) HumidityMax() HumidityMinMax {
	res := a.getSpecificParameters("humax")
	values := []HumidityMinMaxValue{}

	for _, v := range res.Parameters {
		val, _ := strconv.ParseFloat(v.Values[0].Value, 64)

		values = append(values, HumidityMinMaxValue{
			Type:     v.Type,
			Day:      parseDateTime(v.Day, "day"),
			Datetime: parseDateTime(v.Datetime, "datetime"),
			Value:    val,
			Unit:     v.Values[0].Unit,
		})
	}

	return HumidityMinMax{
		ID:          res.ID,
		Description: res.Description,
		Type:        res.Type,
		Data:        values,
	}
}

func (a *AreaModel) HumidityMin() HumidityMinMax {
	res := a.getSpecificParameters("humin")
	values := []HumidityMinMaxValue{}

	for _, v := range res.Parameters {
		val, _ := strconv.ParseFloat(v.Values[0].Value, 64)

		values = append(values, HumidityMinMaxValue{
			Type:     v.Type,
			Day:      parseDateTime(v.Day, "day"),
			Datetime: parseDateTime(v.Datetime, "datetime"),
			Value:    val,
			Unit:     v.Values[0].Unit,
		})
	}

	return HumidityMinMax{
		ID:          res.ID,
		Description: res.Description,
		Type:        res.Type,
		Data:        values,
	}
}

func (a *AreaModel) Temperature() Temperature {
	res := a.getSpecificParameters("t")
	values := []TemperatureValue{}

	for _, v := range res.Parameters {
		c, _ := strconv.ParseFloat(v.Values[0].Value, 64)
		f, _ := strconv.ParseFloat(v.Values[1].Value, 64)

		values = append(values, TemperatureValue{
			Type:       v.Type,
			H:          v.H,
			Datetime:   parseDateTime(v.Datetime, "datetime"),
			Celcius:    c,
			Fahrenheit: f,
		})
	}

	return Temperature{
		ID:          res.ID,
		Description: res.Description,
		Type:        res.Type,
		Data:        values,
	}
}

func (a *AreaModel) TemperatureMax() TemperatureMinMax {
	res := a.getSpecificParameters("tmax")
	values := []TemperatureMinMaxValue{}

	for _, v := range res.Parameters {
		c, _ := strconv.ParseFloat(v.Values[0].Value, 64)
		f, _ := strconv.ParseFloat(v.Values[1].Value, 64)

		values = append(values, TemperatureMinMaxValue{
			Type:       v.Type,
			Day:        parseDateTime(v.Day, "day"),
			Datetime:   parseDateTime(v.Datetime, "datetime"),
			Celcius:    c,
			Fahrenheit: f,
		})
	}

	return TemperatureMinMax{
		ID:          res.ID,
		Description: res.Description,
		Type:        res.Type,
		Data:        values,
	}
}

func (a *AreaModel) TemperatureMin() TemperatureMinMax {
	res := a.getSpecificParameters("tmin")
	values := []TemperatureMinMaxValue{}

	for _, v := range res.Parameters {
		c, _ := strconv.ParseFloat(v.Values[0].Value, 64)
		f, _ := strconv.ParseFloat(v.Values[1].Value, 64)

		values = append(values, TemperatureMinMaxValue{
			Type:       v.Type,
			Day:        parseDateTime(v.Day, "day"),
			Datetime:   parseDateTime(v.Datetime, "datetime"),
			Celcius:    c,
			Fahrenheit: f,
		})
	}

	return TemperatureMinMax{
		ID:          res.ID,
		Description: res.Description,
		Type:        res.Type,
		Data:        values,
	}
}

func (a *AreaModel) Weather() Weather {
	res := a.getSpecificParameters("weather")
	values := []WeatherValue{}

	for _, v := range res.Parameters {
		val, _ := strconv.ParseFloat(v.Values[0].Value, 64)

		values = append(values, WeatherValue{
			Type:     v.Type,
			H:        v.H,
			Datetime: parseDateTime(v.Datetime, "datetime"),
			Value:    val,
			Unit:     v.Values[0].Unit,
		})
	}

	return Weather{
		ID:          res.ID,
		Description: res.Description,
		Type:        res.Type,
		Data:        values,
	}
}

func (a *AreaModel) WindDirection() WindDirection {
	res := a.getSpecificParameters("wd")
	values := []WindDirectionValue{}

	for _, v := range res.Parameters {
		deg, _ := strconv.ParseFloat(v.Values[0].Value, 64)
		card := v.Values[1].Value
		sexa, _ := strconv.ParseFloat(v.Values[2].Value, 64)

		values = append(values, WindDirectionValue{
			Type:     v.Type,
			H:        v.H,
			Datetime: parseDateTime(v.Datetime, "datetime"),
			Degree: struct {
				Unit  string
				Value float64
			}{
				Unit:  "deg",
				Value: deg,
			},
			Card: WindCardValue{
				Unit:  "CARD",
				Value: card,
			},
			Sexa: struct {
				Unit  string
				Value float64
			}{
				Unit:  "SEXA",
				Value: sexa,
			},
		})
	}

	return WindDirection{
		ID:          res.ID,
		Description: res.Description,
		Type:        res.Type,
		Data:        values,
	}
}

func (a *AreaModel) WindSpeed() WindSpeed {
	res := a.getSpecificParameters("ws")
	values := []WindSpeedValue{}

	for _, v := range res.Parameters {
		knot, _ := strconv.ParseFloat(v.Values[0].Value, 64)
		mph, _ := strconv.ParseFloat(v.Values[1].Value, 64)
		kph, _ := strconv.ParseFloat(v.Values[2].Value, 64)
		ms, _ := strconv.ParseFloat(v.Values[3].Value, 64)

		values = append(values, WindSpeedValue{
			Type:     v.Type,
			H:        v.H,
			Datetime: parseDateTime(v.Datetime, "datetime"),
			Knot: struct {
				Unit  string
				Value float64
			}{
				Unit:  "Kt",
				Value: knot,
			},
			Mph: struct {
				Unit  string
				Value float64
			}{
				Unit:  "MPH",
				Value: mph,
			},
			Kph: struct {
				Unit  string
				Value float64
			}{
				Unit:  "KPH",
				Value: kph,
			},
			Ms: struct {
				Unit  string
				Value float64
			}{
				Unit:  "MS",
				Value: ms,
			},
		})
	}

	return WindSpeed{
		ID:          res.ID,
		Description: res.Description,
		Type:        res.Type,
		Data:        values,
	}
}
