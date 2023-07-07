package conversion

import (
    "reflect"
)

func UsingTypeConverter(tc TypeConverter) {
    typeConverters = append(typeConverters, tc)
}

func ConvertUsing[TSource, TDestination any](f func(source TSource) (TDestination, error)) {
    canConvert := func(sourceType reflect.Type, destinationType reflect.Type) bool {
        var s TSource
        var d TDestination
        // check if the types are assignable
        return sourceType == reflect.TypeOf(s) && destinationType == reflect.TypeOf(d)
    }

    convert := func(source reflect.Value, destinationType reflect.Type) (reflect.Value, error) {
        s := source.Interface().(TSource)
        d, err := f(s)
        if err != nil {
            return reflect.Value{}, err
        }
        return reflect.ValueOf(d), nil
    }

    UsingTypeConverter(typeConverter{canConvert: canConvert, convert: convert})
}

// convert function
func convert(source reflect.Value, destinationType reflect.Type) (reflect.Value, error) {
    var dtype, stype reflect.Type
    var value reflect.Value
    var err error

    // if source is a pointer, then dereference it
    if source.Kind() == reflect.Ptr {
        source = source.Elem()
    }
    stype = source.Type()

    // if destinationType is a pointer, then use the element type
    if destinationType.Kind() == reflect.Ptr {
        dtype = destinationType.Elem()
    } else {
        dtype = destinationType
    }

    for _, converter := range typeConverters {
        if converter.CanConvert(stype, dtype) {
            value, err = converter.Convert(source, destinationType)
            if err != nil {
                return reflect.Value{}, err
            }
            if destinationType.Kind() == reflect.Ptr {
                ptr := reflect.New(destinationType.Elem())
                ptr.Elem().Set(value)
                return ptr, nil
            } else {
                return value, nil
            }
        }
    }

    if stype == dtype {
        if source.Kind() == reflect.Ptr {
            // dereference the pointer
            value = source.Elem()
        } else {
            value = source
        }

        if destinationType.Kind() == reflect.Ptr {
            ptr := reflect.New(destinationType.Elem())
            ptr.Elem().Set(value)
            return ptr, nil
        } else {
            return value, nil
        }
    } else if (dtype.Kind() == reflect.Slice || dtype.Kind() == reflect.Array) && (stype.Kind() == reflect.Slice || stype.Kind() == reflect.Array) {
        value, err = convertSlice(source, dtype)
    } else if stype.Kind() == reflect.Struct && dtype.Kind() == reflect.Struct {
        value, err = convertString(source, dtype)
    }

    if err != nil {
        return reflect.Value{}, err
    }

    if destinationType.Kind() == reflect.Ptr {
        ptr := reflect.New(destinationType.Elem())
        ptr.Elem().Set(value)
        return ptr, nil
    } else {
        return value, nil
    }
}

func convertSlice(source reflect.Value, destinationType reflect.Type) (reflect.Value, error) {
    var value reflect.Value
    var err error

    // if source is a pointer, then dereference it
    if source.Kind() == reflect.Ptr {
        source = source.Elem()
    }

    // if destinationType is a pointer, then use the element type
    if destinationType.Kind() == reflect.Ptr {
        destinationType = destinationType.Elem()
    }

    // create a new slice with the same length as the source slice
    value = reflect.MakeSlice(destinationType, source.Len(), source.Len())

    // iterate over the source slice, and convert each element
    for i := 0; i < source.Len(); i++ {
        convertedField, err := convert(source.Index(i), destinationType.Elem())
        if err != nil {
            return reflect.Value{}, err
        }

        // set the converted field
        value.Index(i).Set(convertedField)
    }
    return value, err
}

func convertString(source reflect.Value, dtype reflect.Type) (reflect.Value, error) {
    var value reflect.Value
    value = reflect.New(dtype)

    // iterate over the fields of the source struct, and convert fields with the same name in the destination struct
    for i := 0; i < source.NumField(); i++ {
        sf := source.Type().Field(i)
        df, ok := dtype.FieldByName(sf.Name)
        if ok {
            if !value.Elem().FieldByName(sf.Name).CanSet() {
                // we can't set the field, so let's skip it
                continue
            }
            // convert the field
            convertedField, err := convert(source.Field(i), df.Type)
            if err != nil {
                return reflect.Value{}, err
            }

            // set the converted field
            value.Elem().FieldByName(sf.Name).Set(convertedField)
        }
    }
    return value.Elem(), nil
}

func Convert[TSource, TDestination any](source TSource, destination *TDestination) error {
    // Obtain the reflect.Value of the source
    sourceValue := reflect.ValueOf(source)

    // Obtain the reflect.Type of the destination
    destinationType := reflect.TypeOf(destination).Elem()

    // Call the convert function
    convertedValue, err := convert(sourceValue, destinationType)
    if err != nil {
        return err
    }

    // Set the converted value to the destination pointer
    reflect.ValueOf(destination).Elem().Set(convertedValue)

    return nil
}
