package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/convert", convertHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	temperature := r.FormValue("temperature")
	unit := r.FormValue("unit")

	convertedTemperatures, err := convertTemperature(temperature, unit)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Original temperature: %s\n", temperature)
	fmt.Fprintf(w, "Converted temperatures:\n")
	fmt.Fprintf(w, "Fahrenheit: %f\n", convertedTemperatures.Fahrenheit)
	fmt.Fprintf(w, "Celsius: %f\n", convertedTemperatures.Celsius)
	fmt.Fprintf(w, "Kelvin: %f\n", convertedTemperatures.Kelvin)
}

func convertTemperature(temperature string, unit string) (Temperature, error) {
	temp, err := strconv.ParseFloat(temperature, 64)
	if err != nil {
		return Temperature{}, err
	}

	switch unit {
	case "Fahrenheit":
		return Temperature{
			Fahrenheit: temp,
			Celsius:    fahrenheitToCelsius(temp),
			Kelvin:     celsiusToKelvin(fahrenheitToCelsius(temp)),
		}, nil
	case "Celsius":
		return Temperature{
			Fahrenheit: celsiusToFahrenheit(temp),
			Celsius:    temp,
			Kelvin:     celsiusToKelvin(temp),
		}, nil
	case "Kelvin":
		return Temperature{
			Fahrenheit: celsiusToFahrenheit(kelvinToCelsius(temp)),
			Celsius:    kelvinToCelsius(temp),
			Kelvin:     temp,
		}, nil
	default:
		return Temperature{}, fmt.Errorf("Unrecognized unit: %s", unit)
	}
}

type Temperature struct {
	Fahrenheit float64
	Celsius    float64
	Kelvin     float64
}
