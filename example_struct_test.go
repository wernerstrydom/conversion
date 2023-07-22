package conversion

import "fmt"

type S1 struct {
	FirstName string
	LastName  string
}

type S2 struct {
	FirstName string
	LastName  string
}

type S3 struct {
	FirstName string
	LastName  string
	Age       int
}

type S4 struct {
	FirstName *string
	LastName  *string
}

// This example demonstrates how to map from a struct to a struct with the same
// number of fields.
func ExampleMapper_Map_identical() {
	source := S1{FirstName: "John", LastName: "Doe"}
	var destination S2
	cfg := NewConfiguration()
	mapper := cfg.BuildMapper()

	err := mapper.Map(source, &destination)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(destination.FirstName)
	fmt.Println(destination.LastName)
	// Output:
	// John
	// Doe
}

// This example demonstrates how to map from a struct to a struct with a
// different number of fields.
func ExampleMapper_Map_fieldMismatch() {
	source := S1{FirstName: "John", LastName: "Doe"}
	var destination S3
	cfg := NewConfiguration()
	mapper := cfg.BuildMapper()

	err := mapper.Map(source, &destination)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(destination.FirstName)
	fmt.Println(destination.LastName)
	fmt.Println(destination.Age)
	// Output:
	// John
	// Doe
	// 0
}

// This example demonstrates how to map from a struct to a struct with a
// different number of fields, and how any fields that are not mapped will
// have their original values.
func ExampleMapper_Map_fieldMismatchWithDefaults() {
	source := S1{FirstName: "John", LastName: "Doe"}
	destination := S3{Age: 42}
	cfg := NewConfiguration()
	mapper := cfg.BuildMapper()

	err := mapper.Map(source, &destination)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(destination.FirstName)
	fmt.Println(destination.LastName)
	fmt.Println(destination.Age)
	// Output:
	// John
	// Doe
	// 42
}

// This example demonstrates how to map from a struct to a struct with the same
// fields, except that the destination fields are pointers. This can be useful
// when making API calls where requests often have optional fields.
func ExampleMapper_Map_pointer() {
	source := S1{FirstName: "John", LastName: "Doe"}
	var destination S4
	cfg := NewConfiguration()
	mapper := cfg.BuildMapper()

	err := mapper.Map(source, &destination)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*destination.FirstName)
	fmt.Println(*destination.LastName)
	// Output:
	// John
	// Doe
}
