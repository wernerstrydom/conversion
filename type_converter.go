package conversion

import "reflect"

// TypeConverter is an interface that defines a type converter. A type converter
// is responsible for converting a value from one type to another.
type TypeConverter interface {

    // CanConvert returns true if the type converter can convert a value from
    // sourceType to destinationType.
    CanConvert(sourceType reflect.Type, destinationType reflect.Type) bool

    // Convert converts a value from source to destinationType. If the
    // conversion is not possible, the function will return an error.
    Convert(source reflect.Value, destinationType reflect.Type) (reflect.Value, error)
}

type typeConverter struct {
    canConvert func(sourceType reflect.Type, destinationType reflect.Type) bool
    convert    func(source reflect.Value, destinationType reflect.Type) (reflect.Value, error)
}

func (tc typeConverter) CanConvert(sourceType reflect.Type, destinationType reflect.Type) bool {
    return tc.canConvert(sourceType, destinationType)
}

func (tc typeConverter) Convert(source reflect.Value, destinationType reflect.Type) (reflect.Value, error) {
    return tc.convert(source, destinationType)
}
