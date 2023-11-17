package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherResponse struct {
	CurrentLocation string      `json:"current_location"`
	Weather         WeatherData `json:"weather"`
}

type WeatherData struct {
	Description string  `json:"description"`
	Temperature float64 `json:"temperature"`
	Humidity    int     `json:"humidity"`
}

func getWeatherData(w http.ResponseWriter, r *http.Request) {
	// Replace "YOUR_API_KEY" with your actual OpenWeatherMap API key
	apiKey := "c8583cf5fe4691cfb61a1f199a5508c3"
	apiURL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=Toronto&appid=%s", apiKey)

	// Make a request to the OpenWeatherMap API
	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Error fetching weather data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Parse the API response
	var weatherResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		http.Error(w, "Error decoding weather data", http.StatusInternalServerError)
		return
	}

	// Extract relevant weather data from the response
	description := weatherResponse["weather"].([]interface{})[0].(map[string]interface{})["description"].(string)
	temperature := weatherResponse["main"].(map[string]interface{})["temp"].(float64)
	humidity := int(weatherResponse["main"].(map[string]interface{})["humidity"].(float64))

	// Create a WeatherData struct
	weatherData := WeatherData{
		Description: description,
		Temperature: temperature,
		Humidity:    humidity,
	}

	// Create a response struct
	response := WeatherResponse{
		CurrentLocation: "Toronto",
		Weather:         weatherData,
	}

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the struct to JSON and write the response
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/weather", getWeatherData)
	fmt.Println("Server is running on port 9091")
	// Start the server
	http.ListenAndServe(":9091", nil)
}