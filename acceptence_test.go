package conversion

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

type Struct1 struct {
	Field1 string
	Field2 int
}

type Struct2 struct {
	Field1 string
	Field2 int
}

type Struct3 struct {
	Field1 *string
	Field2 *int
}

func testHelper[TSource, TDestination any](suite *AcceptanceTests, valid TSource, expected TDestination) {
	var err error
	var s TSource
	var d TDestination
	var tDestination reflect.Type = reflect.TypeOf(d)
	var tSource reflect.Type = reflect.TypeOf(s)

	suite.Run(
		fmt.Sprintf("%v", tDestination), func() {

			suite.Run(
				fmt.Sprintf("%v\u2192%v", tSource, tDestination), func() {

					source := valid
					var destination TDestination

					cfg := NewConfiguration()
					mapper := cfg.BuildMapper()
					err = mapper.Map(source, &destination)
					suite.NoError(err)
					if err == nil {
						suite.Equal(expected, destination)
					}
				},
			)

			suite.Run(
				fmt.Sprintf("%v\u2192*%v", tSource, tDestination), func() {

					source := valid
					var destination *TDestination

					cfg := NewConfiguration()
					mapper := cfg.BuildMapper()
					err = mapper.Map(source, &destination)
					suite.NoError(err)
					if err == nil {
						suite.Equal(expected, *destination)
					}
				},
			)

			suite.Run(
				fmt.Sprintf("*%v\u2192*%v", tSource, tDestination), func() {
					value := valid
					source := &value
					var destination *TDestination

					cfg := NewConfiguration()
					mapper := cfg.BuildMapper()
					err = mapper.Map(source, &destination)
					suite.NoError(err)
					if err == nil {
						suite.Equal(expected, *destination)
					}
				},
			)

			suite.Run(
				fmt.Sprintf("*%v\u2192%v", tSource, tDestination), func() {
					value := valid
					source := &value
					var destination TDestination

					cfg := NewConfiguration()
					mapper := cfg.BuildMapper()
					err = mapper.Map(source, &destination)
					suite.NoError(err)
					if err == nil {
						suite.Equal(expected, destination)
					}
				},
			)
		},
	)

}

type AcceptanceTests struct {
	suite.Suite
}

func (suite *AcceptanceTests) SetupTest() {
	fmt.Printf("SetupTest\n")
}

func (suite *AcceptanceTests) TearDownTest() {
	fmt.Printf("TearDownTest\n")
}

func TestAcceptanceTests(t *testing.T) {
	suite.Run(t, new(AcceptanceTests))
}

func (suite *AcceptanceTests) TestMap() {
	suite.Run(
		"int", func() {
			testHelper(suite, 123, 123)
			testHelper(suite, 123, "123")
			testHelper(suite, 123, 123.0)
		},
	)
	suite.Run(
		"string", func() {
			testHelper(suite, "123", "123")
		},
	)
	suite.Run(
		"bool", func() {
			testHelper(suite, true, true)
		},
	)
	suite.Run(
		"Struct", func() {
			s := "123"
			i := 123
			s1 := Struct1{"123", 123}
			s2 := Struct2{"123", 123}
			s3 := Struct3{Field1: &s, Field2: &i}
			testHelper(suite, s1, s1)
			testHelper(suite, s1, s2)
			testHelper(suite, s1, s3)
		},
	)
}

func (suite *AcceptanceTests) TestMap_1() {
	var err error
	source := "123"
	var destination *string

	cfg := NewConfiguration()
	mapper := cfg.BuildMapper()
	err = mapper.Map(source, &destination)
	suite.Nil(err)
	suite.Equal("123", *destination)
}

func (suite *AcceptanceTests) TestMap_2() {
	var err error
	source := "123"
	var destination string

	cfg := NewConfiguration()
	mapper := cfg.BuildMapper()
	err = mapper.Map(source, &destination)
	suite.Nil(err)
	suite.Equal("123", destination)
}

func (suite *AcceptanceTests) TestMap_3() {
	var err error
	value := "123"
	source := &value
	var destination *string

	cfg := NewConfiguration()
	mapper := cfg.BuildMapper()
	err = mapper.Map(source, &destination)
	suite.Nil(err)
	suite.Equal("123", *destination)
}

func (suite *AcceptanceTests) TestMap_4() {
	var err error
	value := "123"
	source := &value
	var destination string

	cfg := NewConfiguration()
	mapper := cfg.BuildMapper()
	err = mapper.Map(source, &destination)
	suite.Nil(err)
	suite.Equal("123", destination)
}
