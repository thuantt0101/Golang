package weatherapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OpenWeatherMapProvider struct {
	APIKey string
	URL    string
}

//Struct data tra ve cua Service
type OpenWeatherMapData struct {
	Current struct {
		KelvinTemp float64 `json:"temp"`
	} `json:"main"`
}

// Implement hàm GetTemperature của WeatherProvider Interface
func (p OpenWeatherMapProvider) GetTemperature(city string) (float64, error) {
	res, err := http.Get(p.URL + p.APIKey + "&q=" + city)
	if err != nil || res.StatusCode != 200 {
		fmt.Println("OpenWeatherMapProvider is error")
		return 0, err
	}
	defer res.Body.Close()
	data := OpenWeatherMapData{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	//Tinh lai theo do c
	tempC := data.Current.KelvinTemp - 273.15
	fmt.Println("openweathermap", tempC)

	return tempC, err

}
