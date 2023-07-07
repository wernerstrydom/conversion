package conversion

import "fmt"

// SourceStruct is a struct with exported fields.
type SourceStruct struct {
    Name  string
    Value int
}

// DestinationStruct is a struct with exported fields that are pointers.
type DestinationStruct struct {
    Name  *string
    Value *int
}

// ExampleConvert_struct demonstrates how to convert a struct. In this example,
// we convert from a struct with exported fields to a struct with exported
// fields that are pointers.
func ExampleConvert_struct() {

    var source = SourceStruct{
        Name:  "John",
        Value: 42,
    }

    var destination DestinationStruct
    err := Convert(source, &destination)
    if err != nil {
        // Handle error
    }
    fmt.Println(*destination.Name, *destination.Value)
    // Output: John 42
}
