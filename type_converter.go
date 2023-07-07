package conversion

import "reflect"

type TypeConverter interface {
    CanConvert(sourceType reflect.Type, destinationType reflect.Type) bool
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
