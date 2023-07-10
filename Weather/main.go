package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Lat     float64 `json:"latitude"`
	Lon     float64 `json:"longitude"`
	Current struct {
		Temp          float64 `json:"temperature"`
		WindSpeed     float64 `json:"windspeed"`
		WindDirection float64 `json:"winddirection"`
		IsDay         int     `json:"is_day"`
		Time          string  `json:"time"`
	} `json:"current_weather"`
}

func main() {
	res, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=37.63&longitude=22.73&hourly=temperature_2m,relativehumidity_2m,apparent_temperature,rain,windspeed_10m,winddirection_10m,windgusts_10m,is_day&daily=temperature_2m_max,temperature_2m_min,apparent_temperature_max,apparent_temperature_min&current_weather=true&forecast_days=1&timezone=auto")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	fmt.Println(weather)
}
