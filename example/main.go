package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"

	. "github.com/ariefsn/go-bmkg"
	"github.com/ariefsn/go-bmkg/models"
)

func prettyPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "  ")
	fmt.Println(string(s))
}

func main() {
	if xmlBytes, err := GetXML("Aceh"); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		var result models.WeatherModel
		xml.Unmarshal(xmlBytes, &result)
		// do what you want with result
		prettyPrint(result)
	}
}
