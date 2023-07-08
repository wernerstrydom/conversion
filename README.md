# Conversion: A Go Module for Type Conversions

![License](https://img.shields.io/badge/license-MIT-blue.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/wernerstrydom/conversion.svg)](https://pkg.go.dev/github.com/wernerstrydom/conversion)

Conversion is a powerful and flexible Go module for converting one data type to another. This library offers a convenient and efficient solution for mapping types in Go, heavily influenced by AutoMapper. It provides out-of-the-box conversions for basic types and allows for registration of custom conversions as per your application needs. The Conversion library is available under the MIT license and open to contributions.

## Table of Contents

- [Installation](#installation)
- [Predefined Conversions](#predefined-conversions)
- [Usage](#usage)
    - [Built-in Conversions](#built-in-conversions)
    - [Registering Custom Conversions](#registering-custom-conversions)
- [Contributing](#contributing)
- [License](#license)

## Installation

To include Conversion in your project, use the following command:

```bash
go get github.com/wernerstrydom/conversion
```

Then, you can import it in your Go files as follows:

```go
import "github.com/wernerstrydom/conversion"
```

Register the default converters by calling the `RegisterDefaultConverters` function in your `init` function:

```go
func init() {
    conversion.RegisterDefaultConverters()
}
```

## Usage

### Built-in Conversions

Conversion provides default converters for basic data types. The `Convert` function is the primary method for converting data types. It converts a source value to a destination value. The destination value must be a pointer, while the source value can be either a pointer or a value.

Here is a basic example:

```go
package main

import (
    "fmt"
    "github.com/wernerstrydom/conversion"
)

func main() {
    var source int = 10
    var destination float64

    err := conversion.Convert(source, &destination)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(destination) // Prints: 10.0
}
```


The following table lists the built-in conversions provided by the `conversion` library:

| Source Type      | Destination Type | Description                             |
|------------------|------------------|-----------------------------------------|
| string           | int              | Converts a string to an int             |
| string           | int64            | Converts a string to an int64           |
| string           | int32            | Converts a string to an int32           |
| string           | int16            | Converts a string to an int16           |
| string           | int8             | Converts a string to an int8            |
| string           | uint             | Converts a string to a uint             |
| string           | uint64           | Converts a string to a uint64           |
| string           | uint32           | Converts a string to a uint32           |
| string           | uint16           | Converts a string to a uint16           |
| string           | uint8            | Converts a string to a uint8            |
| string           | float64          | Converts a string to a float64          |
| string           | float32          | Converts a string to a float32          |
| string           | complex64        | Converts a string to a complex64        |
| string           | complex128       | Converts a string to a complex128       |
| string           | bool             | Converts a string to a bool             |
| string           | time.Time        | Converts a string to a time.Time        |
| string           | time.Duration    | Converts a string to a time.Duration    |
| string           | url.URL          | Converts a string to a url.URL          |
| string           | net.IP           | Converts a string to a net.IP           |
| string           | net.IPNet        | Converts a string to a net.IPNet        |
| string           | net.HardwareAddr | Converts a string to a net.HardwareAddr |
| uint             | string           | Converts a uint to a string             |
| uint8            | string           | Converts a uint8 to a string            |
| uint16           | string           | Converts a uint16 to a string           |
| uint32           | string           | Converts a uint32 to a string           |
| uint64           | string           | Converts a uint64 to a string           |
| int              | string           | Converts an int to a string             |
| int8             | string           | Converts an int8 to a string            |
| int16            | string           | Converts an int16 to a string           |
| int32            | string           | Converts an int32 to a string           |
| int64            | string           | Converts an int64 to a string           |
| float32          | string           | Converts a float32 to a string          |
| float64          | string           | Converts a float64 to a string          |
| bool             | string           | Converts a bool to a string             |
| complex64        | string           | Converts a complex64 to a string        |
| complex128       | string           | Converts a complex128 to a string       |
| time.Duration    | string           | Converts a time.Duration to a string    |
| time.Time        | string           | Converts a time.Time to a string        |
| url.URL          | string           | Converts a url.URL to a string          |
| net.IP           | string           | Converts a net.IP to a string           |
| net.IPNet        | string           | Converts a net.IPNet to a string        |
| net.HardwareAddr | string           | Converts a net.HardwareAddr to a string |
| string           | []byte           | Converts a string to a byte slice       |
| []byte           | string           | Converts a byte slice to a string       |
| []rune           | string           | Converts a rune slice to a string       |
| string           | []rune           | Converts a string to a rune slice       |

### Registering Custom Conversions

You can register custom conversions by defining functions for type conversion and using `ConvertUsing` to register them.

```go
package main

import (
    "fmt"
    "github.com/wernerstrydom/conversion"
)

type MySource struct {
    Value int
}

type MyDestination struct {
    Value string
}

func main() {
    conversion.ConvertUsing(func(source MySource) (MyDestination, error) {
        return MyDestination{Value: fmt.Sprintf("Converted: %d", source.Value)}, nil
    })

    var source = MySource{Value: 10}
    var destination MyDestination

    err := conversion.Convert(source, &destination)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(destination.Value) // Prints: Converted: 10
}
```

## Contributing

Contributions are always welcome! You can contribute in many ways:

- Report issues
- Propose new features
- Improve/fix documentation
- Submit bug fixes or features

Before contributing, please read the [Contributing Guide](CONTRIBUTING.md).

## License

Conversion is licensed under the MIT license. Please see [License File](LICENSE.md) for more information.
