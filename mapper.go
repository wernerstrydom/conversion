package conversion

import (
	"reflect"
)

// Mapper maps values from one type to another
type Mapper interface {
	// Map maps the value to the destination. The destination must be a pointer to a value.
	Map(value, destination any) error

	// MapTo maps the value to a new instance of the destination type
	MapTo(value any, destinationType reflect.Type) (any, error)
}

type mapper struct {
	typeConverters []TypeConverter
}

// Map maps the value to the destination. The destination must be a pointer to a value.
func (m *mapper) Map(source, destination any) error {
	sourceValue := reflect.ValueOf(source)
	destinationValue := reflect.ValueOf(destination)

	if destinationValue.Kind() != reflect.Ptr {
		return NewMappingError(source, sourceValue.Type(), destinationValue.Type())
	}
	destinationValue = destinationValue.Elem()

	if !sourceValue.IsValid() {
		destinationValue.Set(reflect.Zero(destinationValue.Type()))
		return nil
	}

	sourceType := sourceValue.Type()
	destinationType := destinationValue.Type()

	if sourceType == destinationType {
		destinationValue.Set(sourceValue)
		return nil
	}

	if sourceType.Kind() == reflect.Ptr && sourceType.Elem() == destinationType {
		destinationValue.Set(sourceValue.Elem())
		return nil
	}

	if sourceType.Kind() == reflect.Ptr && destinationType.Kind() == reflect.Ptr && sourceType.Elem() == destinationType.Elem() {
		destinationValue.Set(sourceValue)
		return nil
	}

	if sourceType.Kind() != reflect.Ptr && destinationType.Kind() == reflect.Ptr && sourceType == destinationType.Elem() {
		destinationValue.Set(reflect.New(sourceType))
		destinationValue.Elem().Set(sourceValue)
		return nil
	}

	st := sourceType
	dt := destinationType

	if sourceType.Kind() == reflect.Ptr {
		st = sourceType.Elem()
	}
	if destinationType.Kind() == reflect.Ptr {
		dt = destinationType.Elem()
	}
	sv := sourceValue
	dv := destinationValue
	if sv.Kind() == reflect.Ptr {
		sv = sv.Elem()
	}
	// if dv is not nil, and it's a pointer, set dv to the pointer's value
	if dv.IsValid() && dv.Kind() == reflect.Ptr {
		dv = dv.Elem()
	}
	for _, converter := range m.typeConverters {
		if converter.CanConvert(st, dt) {
			value, err := converter.Convert(sv, dv, st, dt, m)
			if err != nil {
				return err
			}
			if destinationType.Kind() == reflect.Ptr {
				destinationValue.Set(reflect.New(destinationType.Elem()))
				destinationValue.Elem().Set(value)
			} else {
				destinationValue.Set(value)
			}
			return nil
		}
	}

	return NewMappingError(source, sourceType, destinationType)
}

// MapTo maps the value to a new instance of the destination type
func (m *mapper) MapTo(value any, destinationType reflect.Type) (any, error) {
	// if value is nil, return zero value of destination type
	if value == nil {
		return reflect.Zero(destinationType).Interface(), nil
	}

	destinationValue := reflect.New(destinationType).Elem()
	err := m.Map(value, destinationValue.Addr().Interface())
	if err != nil {
		return nil, err
	}
	return destinationValue.Interface(), nil
}
