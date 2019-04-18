package weatherapi

// WeatherProvider là interface để 3 provider kia implement hàm GetTemperature
type WeatherProvider interface {
	GetTemperature(city string) (float64, error)
}
