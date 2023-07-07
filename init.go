package conversion

import (
    "fmt"
    "net"
    "net/url"
    "strconv"
    "time"
)

var typeConverters []TypeConverter

func init() {
    // register default type converters
    ConvertUsing(
        func(source string) (int, error) {
            i, err := strconv.ParseInt(source, 10, 64)
            if err != nil {
                return 0, err
            }
            return int(i), nil
        },
    )

    ConvertUsing(
        func(source string) (int64, error) {
            i, err := strconv.ParseInt(source, 10, 64)
            if err != nil {
                return 0, err
            }
            return int64(i), nil
        },
    )

    ConvertUsing(
        func(source string) (int32, error) {
            i, err := strconv.ParseInt(source, 10, 32)
            if err != nil {
                return 0, err
            }
            return int32(i), nil
        },
    )

    ConvertUsing(
        func(source string) (int16, error) {
            i, err := strconv.ParseInt(source, 10, 16)
            if err != nil {
                return 0, err
            }
            return int16(i), nil
        },
    )

    ConvertUsing(
        func(source string) (int8, error) {
            i, err := strconv.ParseInt(source, 10, 8)
            if err != nil {
                return 0, err
            }
            return int8(i), nil
        },
    )
    ConvertUsing(
        func(source string) (uint, error) {
            i, err := strconv.ParseUint(source, 10, 64)
            if err != nil {
                return 0, err
            }
            return uint(i), nil
        },
    )

    ConvertUsing(
        func(source string) (uint64, error) {
            i, err := strconv.ParseUint(source, 10, 64)
            if err != nil {
                return 0, err
            }
            return uint64(i), nil
        },
    )

    ConvertUsing(
        func(source string) (uint32, error) {
            i, err := strconv.ParseUint(source, 10, 32)
            if err != nil {
                return 0, err
            }
            return uint32(i), nil
        },
    )

    ConvertUsing(
        func(source string) (uint16, error) {
            i, err := strconv.ParseUint(source, 10, 16)
            if err != nil {
                return 0, err
            }
            return uint16(i), nil
        },
    )

    ConvertUsing(
        func(source string) (uint8, error) {
            i, err := strconv.ParseUint(source, 10, 8)
            if err != nil {
                return 0, err
            }
            return uint8(i), nil
        },
    )

    ConvertUsing(
        func(source string) (float64, error) {
            i, err := strconv.ParseFloat(source, 64)
            if err != nil {
                return 0, err
            }
            return i, nil
        },
    )

    ConvertUsing(
        func(source string) (float32, error) {
            i, err := strconv.ParseFloat(source, 32)
            if err != nil {
                return 0, err
            }
            return float32(i), nil
        },
    )

    ConvertUsing(
        func(source string) (complex64, error) {
            i, err := strconv.ParseComplex(source, 64)
            if err != nil {
                return 0, err
            }
            return complex64(i), nil
        },
    )

    ConvertUsing(
        func(source string) (complex128, error) {
            i, err := strconv.ParseComplex(source, 128)
            if err != nil {
                return 0, err
            }
            return i, nil
        },
    )

    ConvertUsing(
        func(source string) (bool, error) {
            i, err := strconv.ParseBool(source)
            if err != nil {
                return false, err
            }
            return i, nil
        },
    )

    ConvertUsing(
        func(source string) (time.Time, error) {
            i, err := time.Parse(time.RFC3339, source)
            if err != nil {
                return time.Time{}, err
            }
            return i, nil
        },
    )

    ConvertUsing(
        func(source string) (time.Duration, error) {
            i, err := time.ParseDuration(source)
            if err != nil {
                return time.Duration(0), err
            }
            return i, nil
        },
    )

    ConvertUsing(
        func(source string) (url.URL, error) {
            i, err := url.Parse(source)
            if err != nil {
                return url.URL{}, err
            }
            return *i, nil
        },
    )

    ConvertUsing(
        func(source string) (net.IP, error) {
            i := net.ParseIP(source)
            if i == nil {
                return nil, fmt.Errorf("invalid IP address: %s", source)
            }
            return i, nil
        },
    )

    ConvertUsing(
        func(source string) (net.IPNet, error) {
            _, i, err := net.ParseCIDR(source)
            if err != nil {
                return net.IPNet{}, err
            }
            return *i, nil
        },
    )

    ConvertUsing(
        func(source string) (net.HardwareAddr, error) {
            i, err := net.ParseMAC(source)
            if err != nil {
                return nil, err
            }
            return i, nil
        },
    )

    ConvertUsing(
        func(value uint) (string, error) {
            return strconv.FormatUint(uint64(value), 10), nil
        },
    )
    ConvertUsing(
        func(value uint8) (string, error) {
            return strconv.FormatUint(uint64(value), 10), nil
        },
    )
    ConvertUsing(
        func(value uint16) (string, error) {
            return strconv.FormatUint(uint64(value), 10), nil
        },
    )
    ConvertUsing(
        func(value uint32) (string, error) {
            return strconv.FormatUint(uint64(value), 10), nil
        },
    )

    ConvertUsing(
        func(value uint64) (string, error) {
            return strconv.FormatUint(uint64(value), 10), nil
        },
    )

    ConvertUsing(
        func(value int) (string, error) {
            return strconv.FormatInt(int64(value), 10), nil
        },
    )

    ConvertUsing(
        func(value int8) (string, error) {
            return strconv.FormatInt(int64(value), 10), nil
        },
    )

    ConvertUsing(
        func(value int16) (string, error) {
            return strconv.FormatInt(int64(value), 10), nil
        },
    )

    ConvertUsing(
        func(value int32) (string, error) {
            return strconv.FormatInt(int64(value), 10), nil
        },
    )

    ConvertUsing(
        func(value int64) (string, error) {
            return strconv.FormatInt(int64(value), 10), nil
        },
    )

    ConvertUsing(
        func(value float32) (string, error) {
            return strconv.FormatFloat(float64(value), 'f', -1, 32), nil
        },
    )

    ConvertUsing(
        func(value float64) (string, error) {
            return strconv.FormatFloat(float64(value), 'f', -1, 64), nil
        },
    )

    ConvertUsing(
        func(value bool) (string, error) {
            return strconv.FormatBool(value), nil
        },
    )

    ConvertUsing(
        func(value complex64) (string, error) {
            return strconv.FormatComplex(complex128(value), 'f', -1, 64), nil
        },
    )

    ConvertUsing(
        func(value complex128) (string, error) {
            return strconv.FormatComplex(value, 'f', -1, 128), nil
        },
    )

    ConvertUsing(
        func(duration time.Duration) (string, error) {
            return duration.String(), nil
        },
    )

    ConvertUsing(
        func(value time.Time) (string, error) {
            return value.Format(time.RFC3339), nil
        },
    )

    ConvertUsing(
        func(value url.URL) (string, error) {
            return value.String(), nil
        },
    )

    ConvertUsing(
        func(value net.IP) (string, error) {
            return value.String(), nil
        },
    )

    ConvertUsing(
        func(value net.IPNet) (string, error) {
            return value.String(), nil
        },
    )

    ConvertUsing(
        func(value net.HardwareAddr) (string, error) {
            return value.String(), nil
        },
    )

    ConvertUsing(
        func(value string) ([]byte, error) {
            return []byte(value), nil
        },
    )

    ConvertUsing(
        func(value []byte) (string, error) {
            return string(value), nil
        },
    )

    ConvertUsing(
        func(value []rune) (string, error) {
            return string(value), nil
        },
    )

    ConvertUsing(
        func(value string) ([]rune, error) {
            return []rune(value), nil
        },
    )
}
