package conversion

import "fmt"

type Fahrenheit float64
type Celsius float64

// FahrenheitToCelsius is a function that converts a Fahrenheit temperature to
// Celsius. It is used in the example below.
func FahrenheitToCelsius(f Fahrenheit) (Celsius, error) {
	// Conversion logic from Fahrenheit to Celsius
	return Celsius((f - 32) * 5 / 9), nil
}

// CelsiusToFahrenheit is a function that converts a Celsius temperature to
// Fahrenheit. It is used in the example below.
func CelsiusToFahrenheit(c Celsius) (Fahrenheit, error) {
	// Conversion logic from Celsius to Fahrenheit
	return Fahrenheit((c * 9 / 5) + 32), nil
}

// Example_temperature demonstrates how to convert between two types using a
// function. In this example, we convert from Fahrenheit to Celsius.
func Example_temperature() {

	cfg := NewConfiguration()
	cfg.RegisterTypeConverter(NewTypeConverter(FahrenheitToCelsius))
	cfg.RegisterTypeConverter(NewTypeConverter(CelsiusToFahrenheit))
	mapper := cfg.BuildMapper()

	var err error
	var fahrenheit Fahrenheit = 100
	var celsius Celsius

	err = mapper.Map(fahrenheit, &celsius)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%.2f°F is %.2f°C\n", fahrenheit, celsius)
	}

	err = mapper.Map(celsius, &fahrenheit)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%.2f°C is %.2f°F\n", celsius, fahrenheit)
	}
	// Output:
	// 100.00°F is 37.78°C
	// 37.78°C is 100.00°F
}
