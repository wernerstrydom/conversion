// Package terraform provides converters that converts between Terraform types and Go types.
//
// | Source | Destination | Comments |
// |--------------------|---------|-----|
// | `types.Bool` | `bool` | |
// | `types.Int` | `int` | |
// | `types.Int8` | `int8` | |
// | `types.Int16` | `int16` | |
// | `types.Int32` | `int32` | |
// | `types.Int64` | `int64` | |
// | `types.Uint` | `uint` | |
// | `types.Uint8` | `uint8` | |
// | `types.Uint16` | `uint16` | |
// | `types.Uint32` | `uint32` | |
// | `types.Uint64` | `uint64` | |
// | `types.Float32` | `float32` | |
// | `types.Float64` | `float64` | |
// | `types.String` | `string` | |
// | `types.Bool` | `*bool` | |
// | `types.Int` | `*int` | |
// | `types.Int8` | `*int8` | |
// | `types.Int16` | `*int16` | |
// | `types.Int32` | `*int32` | |
// | `types.Int64` | `*int64` | |
// | `types.Uint` | `*uint` | |
// | `types.Uint8` | `*uint8` | |
// | `types.Uint16` | `*uint16` | |
// | `types.Uint32` | `*uint32` | |
// | `types.Uint64` | `*uint64` | |
// | `types.Float32` | `*float32` | |
// | `types.Float64` | `*float64` | |
// | `types.String` | `*string` | |
// | `types.String` | `bool` | Returns `true` if the string is `1`, `t`, `T`, `TRUE`, `true`, `True`, or false when `0`, `f`, `F`, `FALSE`, `false`, `False`. |
// | `types.String` | `*bool` | Returns `true` if the string is `1`, `t`, `T`, `TRUE`, `true`, `True`, or false when `0`, `f`, `F`, `FALSE`, `false`, `False`. |
// | `types.String` | `int` | |
// | `types.String` | `*int` | |
// | `types.String` | `int8` | |
// | `types.String` | `*int8` | |
// | `types.String` | `int16` | |
// | `types.String` | `*int16` | |
// | `types.String` | `int32` | |
// | `types.String` | `*int32` | |
// | `types.String` | `int64` | |
// | `types.String` | `*int64` | |
// | `types.String` | `uint` | |
// | `types.String` | `*uint` | |
// | `types.String` | `uint8` | |
// | `types.String` | `*uint8` | |
// | `types.String` | `uint16` | |
// | `types.String` | `*uint16` | |
// | `types.String` | `uint32` | |
// | `types.String` | `*uint32` | |
// | `types.String` | `uint64` | |
// | `types.String` | `*uint64` | |
// | `types.String` | `float32` | |
// | `types.String` | `*float32` | |
// | `types.String` | `float64` | |
// | `types.String` | `*float64` | |
// | `types.String` | `string` | |
// | `types.String` | `*string` | |
package terraform
