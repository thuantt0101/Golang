package api

import (
	"encoding/json"
	"net/http"

	"github.com/k0kubun/pp"
)

//type to provide for api service
type ApiopenweathermapProvider struct {
	APIKey string
	URL    string
}

//data type that receive from servicce
type ApiopenweathermapData struct {
	Current struct {
		KelvinTemp float64 `json:"temp"`
		TempMax    float64 `json:"temp_max"`
		TempMin    float64 `json:"temp_min"`
	} `json:"main"`
}

//Implement interface WeatherInterface
func (p ApiopenweathermapProvider) GetTemperature(city string) (float64, error) {
	//1. Gọi đến Service để nhận một respond trả về và một err nếu có
	//2. Lấy kết quả trả về gáng cho một struct data và lấy data từ struct đó trả về cho thân hàm
	res, err := http.Get(p.URL + p.APIKey + "&q=" + city)

	if err != nil || res.StatusCode != 200 {
		pp.Println("Something error when calling openweathermap")
		return 0, err
	}
	defer res.Body.Close()
	data := &ApiopenweathermapData{}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		pp.Println("Something error when parsing data to data struct")
	}

	//if pass above
	pp.Println("TempMax openweathermap is:   ", data.Current.TempMax)
	pp.Println("TempMin openweathermap is:   ", data.Current.TempMin)
	pp.Println("KelvinTemp openweathermap is:   ", data.Current.KelvinTemp)

	tempC := data.Current.KelvinTemp - 273.15
	return tempC, err
}
