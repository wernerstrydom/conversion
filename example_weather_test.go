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

func init() {
    // Register the FahrenheitToCelsius function at startup
    ConvertUsing(FahrenheitToCelsius)
}

// Example_temperature demonstrates how to convert between two types using a
// function. In this example, we convert from Fahrenheit to Celsius.
func Example_temperature() {

    var fahrenheit Fahrenheit = 100
    var celsius Celsius

    err := Convert(fahrenheit, &celsius)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("%.2f°F is %.2f°C\n", fahrenheit, celsius)
    }
    // Output: 100.00°F is 37.78°C
}
