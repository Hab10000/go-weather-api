package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "city query param required", http.StatusBadRequest)
		return
	}

	// Check Redis cache
	if data, found := GetFromCache(city); found {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(data))
		return
	}

	// Call OpenWeather API
	apiKey := os.Getenv("WEATHER_API_KEY")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		http.Error(w, "Failed to fetch weather", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	SetCache(city, string(body)) // Save to cache

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
