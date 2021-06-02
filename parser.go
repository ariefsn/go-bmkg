package gobmkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetXML(province string) ([]byte, error) {
	url := fmt.Sprintf("https://data.bmkg.go.id/DataMKG/MEWS/DigitalForecast/DigitalForecast-%s.xml", province)

	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("read body: %v", err)
	}

	return data, nil
}
