package api

import (
	"encoding/json"
	"net/http"

	"github.com/k0kubun/pp"
)

//DataProvider
type ApibitProvider struct {
	APIKey string
	URL    string
}

//DataReceive struct
type ApibitData struct {
	Current []struct {
		CelsiusTemp float64 `json:"temp"`
		Station     string  `json:"station"`
	} `json:"data"`
}

//implement interface WeatherInterface

func (p ApibitProvider) GetTemperature(city string) (float64, error) {
	//1. Gọi đến service và nhận một respond
	//2. Parse Json và lấy dữ liệu và trả về cho thân ham

	res, err := http.Get(p.URL + p.APIKey + "&city=" + city)

	if err != nil || res.StatusCode != 200 {
		pp.Println("Something error when calling apibitData")
		return 0, err
	}
	defer res.Body.Close()
	data := &ApibitData{}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		pp.Println("Something error when to struct data from json data")
	}

	//Pass obove
	pp.Println("CelsiusTemp is : ", data.Current[0].CelsiusTemp)
	pp.Println("Station is : ", data.Current[0].Station)

	return data.Current[0].CelsiusTemp, err
}
