package conversion

import (
    "net"
    "net/url"
    "reflect"
    "testing"
    "time"
)

type S1 struct {
    StringField string
    IntField    int
}

func (s S1) String() string {
    return s.StringField
}

type S2 struct {
    StringField string
    IntField    int
}

type S3 struct {
    StringField string
    IntField    *int
}

type S4 struct {
    StringField string
}

type S5 struct {
    StringField string
    S4          *S4
}

type S6 struct {
    StringField string
    S4          S4
}

// shorthand for unicode character U+2192 'RIGHTWARDS ARROW'
var arrow = "\u2192"

// convertHelper1 helps to test the Convert function for the following cases:
//
//	T -> T
//	*T -> *T
//	*T -> T
//	T -> *T
func convertHelper1[T any](value T, expected T, wantErr bool) func(t *testing.T) {
    return func(t *testing.T) {
        // get the type name
        typeName := typeName(value)

        t.Run(
            typeName+arrow+typeName, func(t *testing.T) {
                var destination T

                err := Convert(value, &destination)

                if wantErr {
                    if err == nil {
                        t.Errorf("Convert(%v, &destination) = %v, want error", value, destination)
                    }
                } else {
                    if err != nil {
                        t.Errorf("Convert(%v, &destination) = %v, want %v", value, destination, expected)
                    } else if reflect.DeepEqual(destination, expected) == false {
                        t.Errorf("Convert(%v, &destination) = %v, want %v", value, destination, expected)
                    }
                }
            },
        )

        t.Run(
            typeName+arrow+"*"+typeName, func(t *testing.T) {
                var destination *T

                err := Convert(value, &destination)

                if wantErr {
                    if err == nil {
                        t.Errorf("Convert(%v, &destination) = %v, want error", value, destination)
                    }
                } else {
                    if err != nil {
                        t.Errorf("Convert(%v, &destination) returned error: %v", value, err)
                    } else if reflect.DeepEqual(*destination, expected) == false {
                        t.Errorf("Convert(%v, &destination) = %v, want %v", value, destination, expected)
                    }
                }
            },
        )

        t.Run(
            "*"+typeName+arrow+typeName, func(t *testing.T) {
                var destination T

                err := Convert(&value, &destination)

                if wantErr {
                    if err == nil {
                        t.Errorf("Convert(%v, &destination) = %v, want error", value, destination)
                    }
                } else {
                    if err != nil {
                        t.Errorf("Convert(%v, &destination) return error: %v", value, err)
                    } else if reflect.DeepEqual(destination, expected) == false {
                        t.Errorf("Convert(%v, &destination) = %v, want %v", value, destination, expected)
                    }
                }
            },
        )

        t.Run(
            "*"+typeName+arrow+"*"+typeName, func(t *testing.T) {
                var destination *T

                err := Convert(&value, &destination)

                if wantErr {
                    if err == nil {
                        t.Errorf("Convert(%v, &destination) = %v, want error", value, destination)
                    }
                } else {
                    if err != nil {
                        t.Errorf("Convert(%v, &destination) = %v, want %v", value, destination, expected)
                    }

                    if reflect.DeepEqual(*destination, expected) == false {
                        t.Errorf("Convert(%v, &destination) = %v, want %v", value, destination, expected)
                    }
                }
            },
        )
    }
}

func typeName(i any) string {
    t := reflect.TypeOf(i)
    var name string
    for {
        name = t.String() + name
        if t.Kind() != reflect.Ptr {
            break
        }
        t = t.Elem()
    }
    return name
}

func convertHelper2[S, D any](value S, expected D, wantErr bool) func(t *testing.T) {
    return func(t *testing.T) {
        // get the type name
        sName := typeName(value)
        dName := typeName(expected)

        t.Run(
            sName+arrow+dName, func(t *testing.T) {
                var destination D

                err := Convert(value, &destination)

                if wantErr {
                    if err == nil {
                        t.Errorf("Convert(%v, &destination) = %v, want error", value, destination)
                    }
                } else {
                    if err != nil {
                        t.Errorf("Convert(%v, &destination) returned error: %v", value, err)
                    } else if reflect.DeepEqual(destination, expected) == false {
                        t.Errorf("Convert(%v, &destination) = %v, want %v", value, destination, expected)
                    }
                }
            },
        )

        t.Run(
            sName+arrow+"*"+dName, func(t *testing.T) {
                var destination *D

                err := Convert(value, &destination)

                if wantErr {
                    if err == nil {
                        t.Errorf("Convert(%v, &destination) = %v, want error", value, destination)
                    }
                } else {
                    if err != nil {
                        t.Errorf("Convert(%v, &destination) returned error: %v", value, err)
                    } else if reflect.DeepEqual(*destination, expected) == false {
                        t.Errorf("Convert(%v, &destination) = %v, want %v", value, *destination, expected)
                    }
                }
            },
        )

        t.Run(
            "*"+sName+arrow+dName, func(t *testing.T) {
                var destination D

                err := Convert(&value, &destination)

                if wantErr {
                    if err == nil {
                        t.Errorf("Convert(%v, &destination) = %v, want error", value, destination)
                    }
                } else {
                    if err != nil {
                        t.Errorf("Convert(%v, &destination) returned error: %v", value, err)
                    } else if reflect.DeepEqual(destination, expected) == false {
                        t.Errorf("Convert(%v, &destination) = %v, want %v", value, destination, expected)
                    }
                }
            },
        )

        t.Run(
            "*"+sName+arrow+"*"+dName, func(t *testing.T) {
                var destination *D

                err := Convert(&value, &destination)

                if wantErr {
                    if err == nil {
                        t.Errorf("Convert(%v, &destination) = %v, want error", value, destination)
                    }
                } else {
                    if err != nil {
                        t.Errorf("Convert(%v, &destination) returned error: %v", value, err)
                    } else if reflect.DeepEqual(*destination, expected) == false {
                        t.Errorf("Convert(%v, &destination) = %v, want %v", value, *destination, expected)
                    }
                }
            },
        )
    }
}

func TestConvert(t *testing.T) {
    // source and destination is the same type
    t.Run(
        "string", func(t *testing.T) {
            t.Run("string", convertHelper1("hello", "hello", false))
            t.Run("[]byte", convertHelper2("hello", []byte("hello"), false))
            t.Run("[]rune", convertHelper2("hello", []rune("hello"), false))
            t.Run("int", convertHelper2("123", int(123), false))
            t.Run("int8", convertHelper2("123", int8(123), false))
            t.Run("int16", convertHelper2("123", int16(123), false))
            t.Run("int32", convertHelper2("123", int32(123), false))
            t.Run("int64", convertHelper2("123", int64(123), false))
            t.Run("uint", convertHelper2("123", uint(123), false))
            t.Run("uint8", convertHelper2("123", uint8(123), false))
            t.Run("uint16", convertHelper2("123", uint16(123), false))
            t.Run("uint32", convertHelper2("123", uint32(123), false))
            t.Run("uint64", convertHelper2("123", uint64(123), false))
            t.Run("float32", convertHelper2("123.456", float32(123.456), false))
            t.Run("float64", convertHelper2("123.456", float64(123.456), false))
            t.Run("bool", convertHelper2("true", true, false))
            // complex64 and complex128 are supported
            t.Run("complex64", convertHelper2("(123+789i)", complex64(123+789i), false))
            t.Run("complex128", convertHelper2("(123+789i)", complex128(123+789i), false))

            t.Run(
                "net.IP", func(t *testing.T) {
                    t.Run("IPv4", convertHelper2("172.16.0.1", net.IPv4(172, 16, 0, 1), false))
                    t.Run("IPv6", convertHelper2("2001:db8::68", net.ParseIP("2001:db8::68"), false))
                },
            )

            t.Run("URL", convertHelper2("http://example.com", url.URL{Scheme: "http", Host: "example.com"}, false))

            t.Run("Duration", convertHelper2("1h", time.Hour, false))

            t.Run(
                "net.IPNet", func(t *testing.T) {
                    t.Run(
                        "IPv4", func(t *testing.T) {
                            _, ipNet, _ := net.ParseCIDR("172.16.1.0/24")
                            convertHelper2("172.16.1.0/24", ipNet, false)
                        },
                    )
                    t.Run(
                        "IPv6", func(t *testing.T) {
                            _, ipNet, _ := net.ParseCIDR("2001:db8::68/64")
                            convertHelper2("2001:db8::68/64", ipNet, false)
                        },
                    )
                },
            )

            t.Run(
                "net.HardwareAddr",
                convertHelper2("01:23:45:67:89:ab", net.HardwareAddr{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}, false),
            )

            t.Run(
                "time.Time", func(t *testing.T) {
                    t.Run(
                        "RFC3339", func(t *testing.T) {
                            convertHelper2("2018-04-01T07:00:00Z", time.Date(2018, 4, 1, 7, 0, 0, 0, time.UTC), false)
                        },
                    )
                    t.Run(
                        "Unix", func(t *testing.T) {
                            convertHelper2("1522566000", time.Date(2018, 4, 1, 7, 0, 0, 0, time.UTC), false)
                        },
                    )
                },
            )

        },
    )

    t.Run(
        "[]byte", func(t *testing.T) {
            t.Run("[]byte", convertHelper1([]byte("hello"), []byte("hello"), false))
            t.Run("string", convertHelper2([]byte("hello"), "hello", false))
        },
    )

    t.Run(
        "[]rune", func(t *testing.T) {
            t.Run("[]rune", convertHelper1([]rune("hello"), []rune("hello"), false))
            t.Run("string", convertHelper2([]rune("hello"), "hello", false))
        },
    )

    t.Run(
        "int8", func(t *testing.T) {
            t.Run("int8", convertHelper1(int8(123), 123, false))
            t.Run("string", convertHelper2(int8(123), "123", false))
        },
    )

    t.Run(
        "int16", func(t *testing.T) {
            t.Run("int16", convertHelper1(int16(123), 123, false))
            t.Run("string", convertHelper2(int16(123), "123", false))
        },
    )

    t.Run(
        "int32", func(t *testing.T) {
            t.Run("int32", convertHelper1(int32(123), 123, false))
            t.Run("string", convertHelper2(int32(123), "123", false))
        },
    )

    t.Run(
        "int64", func(t *testing.T) {
            t.Run("int64", convertHelper1(int64(123), 123, false))
            t.Run("string", convertHelper2(int64(123), "123", false))
        },
    )

    t.Run(
        "uint8", func(t *testing.T) {
            t.Run("uint8", convertHelper1(uint8(123), 123, false))
            t.Run("string", convertHelper2(uint8(123), "123", false))

        },
    )

    t.Run(
        "uint16", func(t *testing.T) {
            t.Run("uint16", convertHelper1(uint16(123), 123, false))
            t.Run("string", convertHelper2(uint16(123), "123", false))
        },
    )

    t.Run(
        "uint32", func(t *testing.T) {
            t.Run("uint32", convertHelper1(uint32(123), 123, false))
            t.Run("string", convertHelper2(uint32(123), "123", false))
        },
    )

    t.Run(
        "uint64", func(t *testing.T) {
            t.Run("uint64", convertHelper1(uint64(123), 123, false))
            t.Run("string", convertHelper2(uint64(123), "123", false))
        },
    )

    t.Run(
        "float32", func(t *testing.T) {
            t.Run("float32", convertHelper1(float32(123.0), 123.0, false))
            t.Run("string", convertHelper2(float32(123.0), "123", false))
        },
    )

    t.Run(
        "float64", func(t *testing.T) {
            t.Run("float64", convertHelper1(float64(123.0), 123.0, false))
            t.Run("string", convertHelper2(float64(123.0), "123", false))
        },
    )

    t.Run(
        "bool", func(t *testing.T) {
            t.Run("bool", convertHelper1(true, true, false))
            t.Run("string", convertHelper2(true, "true", false))
        },
    )

    // complex64 and complex128 are supported
    t.Run(
        "complex64", func(t *testing.T) {
            t.Run("complex64", convertHelper1(complex64(123+789i), complex64(123+789i), false))
            t.Run("string", convertHelper2(complex64(123+789i), "(123+789i)", false))
        },
    )

    t.Run(
        "complex128", func(t *testing.T) {
            t.Run("complex128", convertHelper1(complex128(123+789i), complex128(123+789i), false))
            t.Run("string", convertHelper2(complex128(123+789i), "(123+789i)", false))
        },
    )

    // source and destination is different type
    t.Run(
        "complex64", func(t *testing.T) {
            t.Run("complex64", convertHelper2(complex64(123+789i), complex64(123+789i), false))
            t.Run("string", convertHelper2(complex64(123+789i), "(123+789i)", false))
        },
    )

    t.Run(
        "complex128", func(t *testing.T) {
            t.Run("complex128", convertHelper2(complex128(123+789i), complex128(123+789i), false))
            t.Run("string", convertHelper2(complex128(123+789i), "(123+789i)", false))
        },
    )

    t.Run(
        "struct", func(t *testing.T) {
            var s1 = S1{
                StringField: "hello",
                IntField:    123,
            }

            var s2 = S2{
                StringField: "hello",
                IntField:    123,
            }
            i := 123
            var s3 = S3{
                StringField: "hello",
                IntField:    &i,
            }

            var s4 = S4{
                StringField: "hello",
            }

            var s5 = S5{
                StringField: "hello",
                S4:          nil,
            }

            var s5_2 = S5{
                StringField: "hello",
                S4:          &s4,
            }

            t.Run("same type", convertHelper1(s1, s1, false))
            t.Run("same layout", convertHelper2(s1, s2, false))
            t.Run("different fields", convertHelper2(s1, s3, false))
            t.Run("different fields", convertHelper2(s3, s1, false))
            t.Run("destination has less fields", convertHelper2(s1, s4, false))
            t.Run(
                "nested struct", func(t *testing.T) {
                    t.Run("missing from source", convertHelper2(s4, s5, false))
                    t.Run("missing from destination", convertHelper2(s5, s4, false))
                    t.Run("field is nil", convertHelper2(s5, s5, false))
                    t.Run("field is not nil", convertHelper2(s5_2, s5_2, false))
                },
            )
        },
    )

    t.Run(
        "[]string", func(t *testing.T) {
            t.Run("[]string", convertHelper1([]string{"A", "B", "C"}, []string{"A", "B", "C"}, false))
            t.Run("[]int", convertHelper2([]string{"1", "2", "3"}, []int{1, 2, 3}, false))
        },
    )

    t.Run(
        "URL", func(t *testing.T) {
            t.Run("URL", convertHelper1(URL("http://example.com"), URL("http://example.com"), false))
            t.Run("string", convertHelper2(URL("http://example.com"), "http://example.com", false))
        },
    )

    t.Run(
        "time.Time", func(t *testing.T) {
            t.Run(
                "time.Time",
                convertHelper1(
                    time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
                    time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
                    false,
                ),
            )
            t.Run("string", convertHelper2(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), "2019-01-01T00:00:00Z", false))
        },
    )

    t.Run(
        "time.Duration", func(t *testing.T) {
            t.Run("time.Duration", convertHelper1(time.Duration(123), time.Duration(123), false))
            t.Run("string", convertHelper2(time.Duration(123), "123ns", false))
        },
    )

    t.Run(
        "net.IP", func(t *testing.T) {
            t.Run("net.IP", convertHelper1(net.IPv4(127, 0, 0, 1), net.IPv4(127, 0, 0, 1), false))
            t.Run("string", convertHelper2(net.IPv4(127, 0, 0, 1), "127.0.0.1", false))
        },
    )

    t.Run(
        "net.IPNet", func(t *testing.T) {
            t.Run(
                "net.IPNet",
                convertHelper1(
                    net.IPNet{IP: net.IPv4(127, 0, 0, 1), Mask: net.CIDRMask(24, 32)},
                    net.IPNet{IP: net.IPv4(127, 0, 0, 1), Mask: net.CIDRMask(24, 32)},
                    false,
                ),
            )
            t.Run(
                "string",
                convertHelper2(
                    net.IPNet{IP: net.IPv4(127, 0, 0, 1), Mask: net.CIDRMask(24, 32)},
                    "127.0.0.1/24",
                    false,
                ),
            )
        },
    )

    t.Run(
        "net.HardwareAddr", func(t *testing.T) {
            t.Run(
                "net.HardwareAddr",
                convertHelper1(
                    net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
                    net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
                    false,
                ),
            )
            t.Run(
                "string",
                convertHelper2(
                    net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
                    "00:00:00:00:00:00",
                    false,
                ),
            )
        },
    )
}

func URL(s string) url.URL {
    u, _ := url.Parse(s)
    return *u
}
