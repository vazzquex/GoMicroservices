package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/convert", convertHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Mehotd not allowed", http.StatusMethodNotAllowed)
		return
	}

	temperature := r.FormValue("temperature")
	unit := r.FormValue("unit")

	convertedTemperatures := convertTemperature(temperature, unit)

	fmt.Fprintf(w, "Original temperature: %s\n", temperature)
	fmt.Fprintf(w, "Converted temperatures:\n")
	fmt.Fprintf(w, "Fahrenheit: %f\n", convertedTemperatures.Fahrenheit)
	fmt.Fprintf(w, "Celsius: %f\n", convertedTemperatures.Celsius)
	fmt.Fprintf(w, "Kelvin: %f\n", convertedTemperatures.Kelvin)

}

func convertTemperature(temperature string, unit string) Temperature {

	// Retorna una estructura Temperature con los valores convertidos.
	return Temperature{
		Fahrenheit: 0.0,
		Celsius:    0.0,
		Kelvin:     0.0,
	}
}

type Temperature struct {
	Fahrenheit float64
	Celsius    float64
	Kelvin     float64
}
