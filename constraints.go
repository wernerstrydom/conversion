package conversion

import "reflect"

type SignedNumber interface {
	int | int8 | int16 | int32 | int64
}

type UnsignedNumber interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type FloatNumber interface {
	float32 | float64
}

type Number interface {
	SignedNumber | UnsignedNumber | FloatNumber
}

func IsSignedNumber(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}

func IsUnsignedNumber(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	default:
		return false
	}
}

func IsFloatNumber(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func IsNumber(t reflect.Type) bool {
	return IsSignedNumber(t) || IsUnsignedNumber(t) || IsFloatNumber(t)
}

func IsPrimitiveType(destination reflect.Type) bool {
	switch destination.Kind() {
	case reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.String:
		return true
	}
	return false
}
