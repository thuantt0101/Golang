package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/k0kubun/pp"
)

//Provider cho apixu.com
//Current Weather
//HTTP: http://api.apixu.com/v1/current.json?key=e9ee939d97064b58bad51316190704&q=Paris
type ApiXuProvider struct {
	APIKey string
	URL    string
}

//ApiXuData: for getting data from service apixu.com

type ApiXuData struct {
	Current struct {
		CelsiusTemp      float64 `json:"temp_c"`
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
	} `json:"current"`
}

//Implement WeatherInterface
func (p ApiXuProvider) GetTemperature(city string) (float64, error) {
	//1. Goi Service của apixu.com cung cấp
	//2. Lấy độ c và nếu có thì là error và trả về
	res, err := http.Get(p.URL + p.APIKey + "&q=" + city)
	if err != nil || res.StatusCode != 200 {
		pp.Println("Something error in ApiXu")
		return 0, err
	}
	defer res.Body.Close()
	data := &ApiXuData{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("Something error when parsing from json result to data struct")
		return 0, err
	}

	fmt.Println("LastUpdatedEpoch is: ", data.Current.LastUpdatedEpoch)
	fmt.Println("Celsius from ApiXu is: ", data.Current.CelsiusTemp)
	return data.Current.CelsiusTemp, err
}
