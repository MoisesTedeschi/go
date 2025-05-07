package main

import "fmt"

const pontoECelsius = 100
const pontoEFahrenheit = 212
const pontoEKelvin = 373

func main() {
	fmt.Println("Conversão de Escalas Termométricas")

	fmt.Println("Celsius para Fahrenheit: ", (pontoECelsius*9/5)+32)
	fmt.Println("Fahrenheit para Celsius: ", (pontoEFahrenheit-32)*5/9)
	fmt.Println("Celsius para Kelvin: ", (pontoECelsius + 273.15))
	fmt.Println("Kelvin para Celsius: ", (pontoEKelvin - 273.15))
}