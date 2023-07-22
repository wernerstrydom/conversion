package conversion

import (
	"reflect"
)

// TypeConverter is an interface for converting types of values from one type to another.
type TypeConverter interface {
	CanConvert(source, destination reflect.Type) bool
	Convert(source, destination reflect.Value, sourceType, destinationType reflect.Type, mapper Mapper) (
		reflect.Value,
		error,
	)
}

type typeConverter[TSource, TDestination any] struct {
	f func(source TSource) (TDestination, error)
}

func (tc typeConverter[TSource, TDestination]) CanConvert(source, destination reflect.Type) bool {
	var s TSource
	var d TDestination
	tSource := reflect.TypeOf(s)
	tDestination := reflect.TypeOf(d)
	canConvert := source == tSource && destination == tDestination
	// log.Printf("CanConvert[%s, %s]: %s  -> %s: %v", tSource, tDestination, source, destination, canConvert)
	return canConvert
}

func (tc typeConverter[TSource, TDestination]) Convert(
	source, destination reflect.Value,
	sourceType, destinationType reflect.Type,
	mapper Mapper,
) (reflect.Value, error) {
	s := source.Interface().(TSource)
	d, err := tc.f(s)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(d), nil
}

func NewTypeConverter[TSource, TDestination any](
	f func(source TSource) (TDestination, error),
) TypeConverter {
	return typeConverter[TSource, TDestination]{f: f}
}
