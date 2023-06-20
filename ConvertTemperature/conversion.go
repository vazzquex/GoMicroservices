package main

// Funciones de conversion de temperaturas

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}
