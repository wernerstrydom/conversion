package conversion

import (
    "fmt"
    "reflect"
)

type Pound float64
type Kilogram float64

// WeightTypeConverter is a type converter that can be used to convert a pounds
// value to a kilograms value.
type WeightTypeConverter struct{}

// CanConvert returns true if the sourceType is a Pound and the destinationType
// is a Kilogram.
func (tc WeightTypeConverter) CanConvert(sourceType reflect.Type, destinationType reflect.Type) bool {
    var p Pound
    var k Kilogram
    // check if the types are assignable
    return sourceType == reflect.TypeOf(p) && destinationType == reflect.TypeOf(k)
}

// Convert converts pounds to kilograms.
func (tc WeightTypeConverter) Convert(source reflect.Value, destinationType reflect.Type) (reflect.Value, error) {
    p := source.Interface().(Pound)
    k := Kilogram(p * 0.45359237)
    return reflect.ValueOf(k), nil
}

func init() {
    UsingTypeConverter(WeightTypeConverter{})
}

// Example_weight demonstrates how to convert between two types using a type
// converter. In this example, we convert from pounds to kilograms.
func Example_weight() {

    var pound Pound = 100
    var kilogram Kilogram

    err := Convert(pound, &kilogram)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("%.2f lbs is %.2f kg\n", pound, kilogram)
    }
    // Output: 100.00 lbs is 45.36 kg
}
