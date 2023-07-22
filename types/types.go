package types

import "reflect"

var String = reflect.TypeOf("")
var Int = reflect.TypeOf(0)
var Int8 = reflect.TypeOf(int8(0))
var Int16 = reflect.TypeOf(int16(0))
var Int32 = reflect.TypeOf(int32(0))
var Int64 = reflect.TypeOf(int64(0))
var Uint = reflect.TypeOf(uint(0))
var Uint8 = reflect.TypeOf(uint8(0))
var Uint16 = reflect.TypeOf(uint16(0))
var Uint32 = reflect.TypeOf(uint32(0))
var Uint64 = reflect.TypeOf(uint64(0))
var Float32 = reflect.TypeOf(float32(0))
var Float64 = reflect.TypeOf(float64(0))
var Bool = reflect.TypeOf(false)
var Byte = reflect.TypeOf(byte(0))
var Rune = reflect.TypeOf(rune(0))
var Complex64 = reflect.TypeOf(complex64(0))
var Complex128 = reflect.TypeOf(complex128(0))
var Uintptr = reflect.TypeOf(uintptr(0))
