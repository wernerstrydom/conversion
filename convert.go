package conversion

import "reflect"

// DefaultConfiguration is the default configuration for the conversion package
var DefaultConfiguration = NewConfiguration()

// Convert converts the source to the destination
func Convert[TSource, TDestination any](source TSource, destination *TDestination) error {
	mapper := DefaultConfiguration.BuildMapper()
	return mapper.Map(source, destination)
}

// ConvertTo converts the value to the destination type
func ConvertTo[TDestination any](value any) (TDestination, error) {
	var destination TDestination
	mapper := DefaultConfiguration.BuildMapper()
	destinationType := reflect.TypeOf(destination)
	to, err := mapper.MapTo(value, destinationType)
	if err != nil {
		return destination, err
	}
	return to.(TDestination), nil
}
