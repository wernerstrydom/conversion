package conversion

import "reflect"

var (
	_                   TypeConverter = structTypeConverter{}
	StructTypeConverter TypeConverter
)

type structTypeConverter struct{}

func (tc structTypeConverter) CanConvert(source, destination reflect.Type) bool {
	canConvert := source.Kind() == reflect.Struct && destination.Kind() == reflect.Struct
	// log.Printf("CanConvert[%s, %s]: %s  -> %s: %v", tSource, tDestination, source, destination, canConvert)
	return canConvert
}

func (tc structTypeConverter) Convert(
	source, destination reflect.Value,
	sourceType, destinationType reflect.Type,
	mapper Mapper,
) (reflect.Value, error) {

	// if source and destination types are the same, just return the source
	if sourceType == destinationType {
		return source, nil
	}

	// if destination is nil, create a new instance of the destination type
	if !destination.IsValid() {
		destination = reflect.New(destinationType).Elem()
	}

	// copy the fields from source to destination when the field names match
	for i := 0; i < source.NumField(); i++ {
		sourceField := source.Field(i)
		sourceFieldName := sourceType.Field(i).Name
		destinationField := destination.FieldByName(sourceFieldName)
		if destinationField.IsValid() {
			// if destinationField and sourceField are the same type, just set the destinationField to the sourceField
			if destinationField.Type() == sourceField.Type() {
				destinationField.Set(sourceField)
				continue
			}

			// if they are not the same type, use the map context to map the sourceField to the destinationField
			err := mapper.Map(sourceField.Interface(), destinationField.Addr().Interface())
			if err != nil {
				return reflect.Value{}, err
			}
		}
	}
	return destination, nil
}

func newStructTypeConverter() TypeConverter {
	return structTypeConverter{}
}

func init() {
	StructTypeConverter = newStructTypeConverter()
}
