package conversion

import (
	"fmt"
	"reflect"

	"github.com/wernerstrydom/conversion/types"
)

func ExampleMapper_Map() {
	source := "123"
	var destination int
	cfg := NewConfiguration()
	mapper := cfg.BuildMapper()

	err := mapper.Map(source, &destination)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(destination)
	// Output:
	// 123
}

func ExampleMapper_Map_pointerToValue() {
	value := "123"
	source := &value
	var destination int
	cfg := NewConfiguration()
	mapper := cfg.BuildMapper()

	err := mapper.Map(source, &destination)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(destination)
	// Output:
	// 123
}

func ExampleMapper_Map_valueToPointer() {
	source := "123"
	var destination *int
	cfg := NewConfiguration()
	mapper := cfg.BuildMapper()

	err := mapper.Map(source, &destination)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*destination)
	// Output:
	// 123
}

func ExampleMapper_Map_pointerToPointer() {
	value := "123"
	source := &value
	var destination *int
	cfg := NewConfiguration()
	mapper := cfg.BuildMapper()

	err := mapper.Map(source, &destination)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*destination)
	// Output:
	// 123
}

func ExampleMapper_MapTo_stringToInt() {
	source := "123"
	cfg := NewConfiguration()
	mapper := cfg.BuildMapper()

	destination, err := mapper.MapTo(source, reflect.TypeOf(0))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(destination.(int))
	// Output:
	// 123
}

func ExampleMapper_MapTo_parseError() {
	source := "123.45"
	cfg := NewConfiguration()
	mapper := cfg.BuildMapper()

	_, err := mapper.MapTo(source, reflect.TypeOf(0))
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// strconv.Atoi: parsing "123.45": invalid syntax
}

func ExampleMapper_MapTo() {
	source := true
	cfg := NewConfiguration()
	mapper := cfg.BuildMapper()

	destination, err := mapper.MapTo(source, types.Int)
	if err != nil {
		fmt.Println(err)
	}

	// cast destination to int
	fmt.Println(destination.(int))
	// Output:
	// 1
}
