package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	wea "github.com/Golang/Proj/weather-services2/api"
	"github.com/gorilla/mux"
)

//1. Create List Provider Interface
//2. Insert data for this list
//3. define function that get datatemp from some api and calculate average
//3. loop over and go routine to to exec this func
//Provider List
type ProviderList []wea.WeatherProvider

type TemperatureData struct {
	CityName       string  `json:"city_name"`
	CelsiusTemp    float64 `json:"celsius_temp"`
	KelvinTemp     float64 `json:"kelvin_temp"`
	FahrenheitTemp float64 `json:"fahrenheit_temp"`
}

// Lấy dữ liệu nhiệt độ và tính trung bình
func (list ProviderList) temperature(city string) float64 {
	//1.Tao Channel de hung data tra ve ru goroutine
	//2. Tao cac routine de lay data tu 3 nguon
	//a. Open weather map
	//b. ApiXu
	//c. Weather bit
	//3. Doc du lieu tu channel ra
	chanTemp := make(chan float64)
	chanErr := make(chan error)

	for _, p := range list {

		go func(w wea.WeatherProvider) {
			temp, err := w.GetTemperature(city)
			if err != nil {
				chanErr <- err
				return
			}

			chanTemp <- temp
		}(p)
	}

	total := 0.0
	k := 0.0

	for i := 0; i < len(list); i++ {
		select {
		case temp := <-chanTemp:
			if temp > 0 {
				total += temp
				k++
			}
		case err := <-chanErr:
			if err != nil {
				panic(err)
			}
		}
	}

	return total / float64(k)

}

func main() {

	//1. Tao ra cac provider data
	apiopenweathemap := wea.ApiopenweathermapProvider{
		APIKey: "3d9d488fe5ec85a1d3545b6f8342a402",
		URL:    "https://api.openweathermap.org/data/2.5/weather?appid=",
	}

	apiweatherbit := wea.ApibitProvider{
		APIKey: "6fcef5a0eb824de2a8522cf80c6c0f54",
		URL:    "https://api.weatherbit.io/v2.0/current?key=",
	}

	apixu := wea.ApiXuProvider{
		APIKey: "e9ee939d97064b58bad51316190704",
		URL:    "https://api.apixu.com/v1/current.json?key=",
	}

	providerList := ProviderList{
		apiopenweathemap,
		apiweatherbit,
		apixu,
	}

	// Xử lý Rest API sử dụng thư viện Gorilla Mux
	r := mux.NewRouter()
	r.HandleFunc("/api/temperature/{city}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		city := vars["city"]

		//Lay nhiet do
		tempC := providerList.temperature(city)
		tempK := tempC + 273.15
		tempF := (tempC * 1.8) + 32

		data := TemperatureData{
			CityName:       city,
			CelsiusTemp:    tempC,
			KelvinTemp:     tempK,
			FahrenheitTemp: tempF,
		}

		fmt.Printf("Temperature of %s is %f Celsius, %f Kelvin, %f Fahrenheit\n\n", city, tempC, tempK, tempF)
		w.Header().Set("Content-Type", "application/json")

		//Endcode Json vao duong tuong data-->tra ket qua ve cho client
		json.NewEncoder(w).Encode(data)

	}).Methods("GET")

	port := 9000
	fmt.Printf("Server is listening at port: %d\n", port)
	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(port), r))
}
