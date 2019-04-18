package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	wea "github.com/Golang/Proj/weather-services1/weatherapi"
	"github.com/gorilla/mux"
	"github.com/k0kubun/pp"
)

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

	//1. Tao channel de hung du data va err tra ve tu routine
	chanTemp := make(chan float64)
	chanErr := make(chan error)

	//2. Tao cac routine de lay data tu 3 nguon
	//a. Open weather map
	//b. ApiXu
	//c. Weather bit
	for _, p := range list {
		//run routine
		go func(w wea.WeatherProvider) {
			temp, err := w.GetTemperature(city)
			if err != nil {
				chanErr <- err
				return
			}
			//lay du lieu nhiet do dua vao channel
			chanTemp <- temp

		}(p)
	}
	total := 0.0
	k := 0.0
	//lay du lieu tu cac channel neu co
	for i := 0; i < len(list); i++ {
		select {
		case temp := <-chanTemp:
			if temp > 0 {
				total += temp
				k++
			}
		case err := <-chanErr:
			panic(err)
		}
	}
	//Sau do tinh trung binh nhiet do va tra ket qua
	return total / float64(k)
}

func main() {
	pp.Println("Welcome to Weather Average API Program")

	//Tao Provider de goi api openweathermap.org
	openWeatherMap := wea.OpenWeatherMapProvider{
		APIKey: "3d9d488fe5ec85a1d3545b6f8342a402",
		URL:    "https://api.openweathermap.org/data/2.5/weather?appid=",
	}

	// Tạo provider để gọi api apixu.com
	apiXu := wea.ApiXuProvider{
		APIKey: "e9ee939d97064b58bad51316190704",
		URL:    "https://api.apixu.com/v1/current.json?key=",
	}

	// Tạo provider để gọi api weatherbit.io
	weatherBit := wea.WeatherBitProvider{
		APIKey: "6fcef5a0eb824de2a8522cf80c6c0f54",
		URL:    "https://api.weatherbit.io/v2.0/current?key=",
	}

	// Danh sách chứa các service
	providerList := ProviderList{
		openWeatherMap,
		apiXu,
		weatherBit,
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
