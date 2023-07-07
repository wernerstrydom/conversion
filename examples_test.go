package conversion

import (
    "fmt"
    "net"
    "net/url"
)

// ExampleConvert_float demonstrates how to convert a string to a float64.
func ExampleConvert_float() {
    var result float64
    err := Convert("3.14", &result)
    if err != nil {
        // Handle error
    }
    fmt.Println(result)
    // Output: 3.14
}

// ExampleConvert_int64 demonstrates how to convert a string to an int64.
func ExampleConvert_int64() {
    var result int64
    err := Convert("42", &result)
    if err != nil {
        // Handle error
    }
    fmt.Println(result)
    // Output: 42
}

// ExampleConvert_ip_address demonstrates how to convert a string to a net.IP.
func ExampleConvert_ip_address() {
    var result net.IP
    err := Convert("172.16.2.1", &result)
    if err != nil {
        // Handle error
    }
    fmt.Println(result.String())
    // Output: 172.16.2.1
}

// ExampleConvert_ipnet demonstrates how to convert a string to a net.IPNet.
func ExampleConvert_ipnet() {
    var result net.IPNet
    err := Convert("192.168.0.0/24", &result)
    if err != nil {
        // Handle error
    }
    fmt.Println(result.String())
    // Output: 192.168.0.0/24
}

// ExampleConvert_hardwareAddr demonstrates how to convert a string to a
// net.HardwareAddr.
func ExampleConvert_hardwareAddr() {
    var result net.HardwareAddr
    err := Convert("00:11:22:33:44:55", &result)
    if err != nil {
        // Handle error
    }
    fmt.Println(result.String())
    // Output: 00:11:22:33:44:55
}

// ExampleConvert_url demonstrates how to convert a string to a url.URL.
func ExampleConvert_url() {
    var result url.URL
    err := Convert("https://www.example.com", &result)
    if err != nil {
        // Handle error
    }
    fmt.Println(result.String())
    // Output: https://www.example.com
}
