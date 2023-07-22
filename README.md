# Conversion: A Go Module for Type Conversions

![License](https://img.shields.io/badge/license-MIT-blue.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/wernerstrydom/conversion.svg)](https://pkg.go.dev/github.com/wernerstrydom/conversion)
# Conversion
)
[![GoDoc](https://godoc.org/github.com/wernerstrydom/conversion?status.svg)](https://godoc.org/github.com/wernerstrydom/conversion)

Conversion is a Go library inspired by .NET's AutoMapper, designed to streamline and simplify the process of converting one type to another.

This library's primary goal is to abstract the boilerplate code of converting disparate types, allowing developers to focus on their application's core logic.

## Features

- Simplified type conversion.
- Decouples conversion logic from application business logic.
- Support for custom conversion functions.
- Configurable and extensible.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

The project requires:

- [Go](https://golang.org/doc/install) (version 1.18 or higher)

### Installing

To start using Conversion, install Go and run `go get`:

```sh
$ go get github.com/wernerstrydom/conversion
```

This will retrieve the library.

## Usage

Here is a simple example of how to use the Conversion library to convert temperatures between Fahrenheit and Celsius:

```go
package main

import (
  "fmt"

  "github.com/wernerstrydom/conversion"
)

type Fahrenheit float64
type Celsius float64

func FahrenheitToCelsius(f Fahrenheit) (Celsius, error) {
  return Celsius((f - 32) * 5 / 9), nil
}

func CelsiusToFahrenheit(c Celsius) (Fahrenheit, error) {
  return Fahrenheit((c * 9 / 5) + 32), nil
}

func main() {

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
}
```

## Documentation

More detailed documentation is available on [godoc](https://godoc.org/github.com/wernerstrydom/conversion).

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

- AutoMapper in .NET, which inspired this project.
