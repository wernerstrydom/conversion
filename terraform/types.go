package terraform

import (
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/wernerstrydom/conversion"
)

// ConvertStringToTerraformString converts a string to Terraform plugin framework string
func ConvertStringToTerraformString(source string) (types.String, error) {
	return types.StringValue(source), nil
}

// ConvertSignedNumberToTerraformString converts a signed number to Terraform plugin framework string
func ConvertSignedNumberToTerraformString[T conversion.SignedNumber](value T) (types.String, error) {
	s := strconv.FormatInt(int64(value), 10)
	return types.StringValue(s), nil
}

// ConvertUnsignedNumberToTerraformString converts an unsigned number to Terraform plugin framework string
func ConvertUnsignedNumberToTerraformString[T conversion.UnsignedNumber](value T) (types.String, error) {
	s := strconv.FormatUint(uint64(value), 10)
	return types.StringValue(s), nil
}

// ConvertSignedNumberPtrToTerraformString converts a signed number pointer to Terraform plugin framework string
func ConvertSignedNumberPtrToTerraformString[T conversion.SignedNumber](value *T) (types.String, error) {
	if value == nil {
		return types.StringPointerValue(nil), nil
	}
	s := strconv.FormatInt(int64(*value), 10)
	return types.StringValue(s), nil
}

// ConvertUnsignedNumberPtrToTerraformString converts an unsigned number pointer to Terraform plugin framework string
func ConvertUnsignedNumberPtrToTerraformString[T conversion.UnsignedNumber](value *T) (types.String, error) {
	if value == nil {
		return types.StringPointerValue(nil), nil
	}
	s := strconv.FormatUint(uint64(*value), 10)
	return types.StringValue(s), nil
}

// ConvertFloat32ToTerraformString converts a float32 to Terraform plugin framework string
func ConvertFloat32ToTerraformString(value float32) (types.String, error) {
	s := strconv.FormatFloat(float64(value), 'f', -1, 32)
	return types.StringValue(s), nil
}

// ConvertFloat64ToTerraformString converts a float64 to Terraform plugin framework string
func ConvertFloat64ToTerraformString(value float64) (types.String, error) {
	s := strconv.FormatFloat(value, 'f', -1, 64)
	return types.StringValue(s), nil
}

// ConvertFloat64PtrToTerraformString converts a float32 pointer to Terraform plugin framework string
func ConvertFloat64PtrToTerraformString(value *float64) (types.String, error) {
	if value == nil {
		return types.StringPointerValue(nil), nil
	}
	s := strconv.FormatFloat(*value, 'f', -1, 64)
	return types.StringValue(s), nil
}

// ConvertFloat32PtrToTerraformString converts a float32 pointer to Terraform plugin framework string
func ConvertFloat32PtrToTerraformString(value *float32) (types.String, error) {
	if value == nil {
		return types.StringPointerValue(nil), nil
	}
	s := strconv.FormatFloat(float64(*value), 'f', -1, 32)
	return types.StringValue(s), nil
}

// ConvertStringPtrToTerraformString converts a string pointer to Terraform plugin framework string
func ConvertStringPtrToTerraformString(source *string) (types.String, error) {
	return types.StringPointerValue(source), nil
}

// ConvertTerraformStringToString converts a Terraform plugin framework string to string
func ConvertTerraformStringToString(value types.String) (string, error) {
	return value.ValueString(), nil
}

// ConvertTerraformStringToStringPtr converts a Terraform plugin framework string to string pointer
func ConvertTerraformStringToStringPtr(value types.String) (*string, error) {
	return value.ValueStringPointer(), nil
}

// ConvertBoolToTerraformBool converts a bool to Terraform plugin framework bool
func ConvertBoolToTerraformBool(source bool) (types.Bool, error) {
	return types.BoolValue(source), nil
}

// ConvertBoolPtrToTerraformBool converts a bool pointer to Terraform plugin framework bool
func ConvertBoolPtrToTerraformBool(source *bool) (types.Bool, error) {
	return types.BoolPointerValue(source), nil
}

// ConvertTerraformBoolToBool converts a Terraform plugin framework bool to bool
func ConvertTerraformBoolToBool(value types.Bool) (bool, error) {
	return value.ValueBool(), nil
}

// ConvertTerraformBoolToBoolPtr converts a Terraform plugin framework bool to bool pointer
func ConvertTerraformBoolToBoolPtr(value types.Bool) (*bool, error) {
	return value.ValueBoolPointer(), nil
}

// ConvertFloat64ToTerraformFloat64 converts a float64 to Terraform plugin framework float64
func ConvertFloat64ToTerraformFloat64(source float64) (types.Float64, error) {
	return types.Float64Value(source), nil
}

// ConvertFloat64PtrToTerraformFloat64 converts a float64 pointer to Terraform plugin framework float64
func ConvertFloat64PtrToTerraformFloat64(source *float64) (types.Float64, error) {
	return types.Float64PointerValue(source), nil
}

// ConvertTerraformFloat64ToFloat64 converts a Terraform plugin framework float64 to float64
func ConvertTerraformFloat64ToFloat64(value types.Float64) (float64, error) {
	return value.ValueFloat64(), nil
}

// ConvertTerraformFloat64ToFloat64Ptr converts a Terraform plugin framework float64 to float64 pointer
func ConvertTerraformFloat64ToFloat64Ptr(value types.Float64) (*float64, error) {
	return value.ValueFloat64Pointer(), nil
}

// ConvertInt64ToTerraformInt64 converts a int64 to Terraform plugin framework int64
func ConvertInt64ToTerraformInt64(source int64) (types.Int64, error) {
	return types.Int64Value(source), nil
}

// ConvertInt64PtrToTerraformInt64 converts a int64 pointer to Terraform plugin framework int64
func ConvertInt64PtrToTerraformInt64(source *int64) (types.Int64, error) {
	return types.Int64PointerValue(source), nil
}

// ConvertTerraformInt64ToInt64 converts a Terraform plugin framework int64 to int64
func ConvertTerraformInt64ToInt64(value types.Int64) (int64, error) {
	return value.ValueInt64(), nil
}

// ConvertTerraformInt64ToInt64Ptr converts a Terraform plugin framework int64 to int64 pointer
func ConvertTerraformInt64ToInt64Ptr(value types.Int64) (*int64, error) {
	return value.ValueInt64Pointer(), nil
}

// BitSize returns the bit size of a signed or unsigned number
func BitSize[T conversion.SignedNumber | conversion.UnsignedNumber]() int {
	var value T
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(t.Size()) * 8
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int(t.Size()) * 8
	default:
		panic("Unsupported type")
	}
}

// ConvertTerraformStringToUnsignedNumber converts a Terraform plugin framework string to an unsigned number
func ConvertTerraformStringToUnsignedNumber[T conversion.UnsignedNumber](value types.String) (T, error) {
	s := value.ValueString()
	i, err := strconv.ParseUint(s, 10, BitSize[T]())
	if err != nil {
		return 0, err
	}
	return T(i), nil
}

// ConvertTerraformStringToSignedNumber converts a Terraform plugin framework string to a signed number
func ConvertTerraformStringToSignedNumber[T conversion.SignedNumber](value types.String) (T, error) {
	s := value.ValueString()
	i, err := strconv.ParseInt(s, 10, BitSize[T]())
	if err != nil {
		return 0, err
	}
	return T(i), nil
}

// ConvertTerraformStringToSignedNumberPtr converts a Terraform plugin framework string to a signed number pointer
func ConvertTerraformStringToSignedNumberPtr[T conversion.SignedNumber](value types.String) (*T, error) {
	if value.IsNull() {
		return nil, nil
	}
	s := value.ValueString()
	i, err := strconv.ParseInt(s, 10, BitSize[T]())
	if err != nil {
		return nil, err
	}
	result := T(i)
	return &result, nil

}

// ConvertTerraformStringUnsignedNumberPtr converts a Terraform plugin framework string to an unsigned number pointer
func ConvertTerraformStringUnsignedNumberPtr[T conversion.UnsignedNumber](value types.String) (*T, error) {
	if value.IsNull() {
		return nil, nil
	}
	s := value.ValueString()
	i, err := strconv.ParseUint(s, 10, BitSize[T]())
	if err != nil {
		return nil, err
	}
	result := T(i)
	return &result, nil
}

func ConvertBoolToTerraformString(source bool) (types.String, error) {
	return types.StringValue(strconv.FormatBool(source)), nil
}

func ConvertBoolPtrToTerraformString(source *bool) (types.String, error) {
	if source == nil {
		return types.StringPointerValue(nil), nil
	}
	return types.StringValue(strconv.FormatBool(*source)), nil
}

func ConvertTerraformStringToBool(value types.String) (bool, error) {
	return strconv.ParseBool(value.ValueString())
}

func ConvertTerraformStringToBoolPtr(value types.String) (*bool, error) {
	if value.IsNull() {
		return nil, nil
	}
	result, err := strconv.ParseBool(value.ValueString())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Converters returns the type converters for transforming between Terraform plugin framework types and Go types
func Converters() []conversion.TypeConverter {
	return []conversion.TypeConverter{
		conversion.NewTypeConverter(ConvertStringToTerraformString),
		conversion.NewTypeConverter(ConvertStringPtrToTerraformString),
		conversion.NewTypeConverter(ConvertTerraformStringToString),
		conversion.NewTypeConverter(ConvertTerraformStringToStringPtr),
		conversion.NewTypeConverter(ConvertBoolToTerraformBool),
		conversion.NewTypeConverter(ConvertBoolPtrToTerraformBool),
		conversion.NewTypeConverter(ConvertTerraformBoolToBool),
		conversion.NewTypeConverter(ConvertTerraformBoolToBoolPtr),
		conversion.NewTypeConverter(ConvertFloat64ToTerraformFloat64),
		conversion.NewTypeConverter(ConvertFloat64PtrToTerraformFloat64),
		conversion.NewTypeConverter(ConvertTerraformFloat64ToFloat64),
		conversion.NewTypeConverter(ConvertTerraformFloat64ToFloat64Ptr),
		conversion.NewTypeConverter(ConvertInt64ToTerraformInt64),
		conversion.NewTypeConverter(ConvertInt64PtrToTerraformInt64),
		conversion.NewTypeConverter(ConvertTerraformInt64ToInt64),
		conversion.NewTypeConverter(ConvertTerraformInt64ToInt64Ptr),

		conversion.NewTypeConverter(ConvertSignedNumberToTerraformString[int8]),
		conversion.NewTypeConverter(ConvertSignedNumberToTerraformString[int16]),
		conversion.NewTypeConverter(ConvertSignedNumberToTerraformString[int32]),
		conversion.NewTypeConverter(ConvertSignedNumberToTerraformString[int64]),
		conversion.NewTypeConverter(ConvertUnsignedNumberToTerraformString[uint8]),
		conversion.NewTypeConverter(ConvertUnsignedNumberToTerraformString[uint16]),
		conversion.NewTypeConverter(ConvertUnsignedNumberToTerraformString[uint32]),
		conversion.NewTypeConverter(ConvertUnsignedNumberToTerraformString[uint64]),
		conversion.NewTypeConverter(ConvertFloat32ToTerraformString),
		conversion.NewTypeConverter(ConvertFloat64ToTerraformString),

		conversion.NewTypeConverter(ConvertSignedNumberPtrToTerraformString[int8]),
		conversion.NewTypeConverter(ConvertSignedNumberPtrToTerraformString[int16]),
		conversion.NewTypeConverter(ConvertSignedNumberPtrToTerraformString[int32]),
		conversion.NewTypeConverter(ConvertSignedNumberPtrToTerraformString[int64]),
		conversion.NewTypeConverter(ConvertUnsignedNumberPtrToTerraformString[uint8]),
		conversion.NewTypeConverter(ConvertUnsignedNumberPtrToTerraformString[uint16]),
		conversion.NewTypeConverter(ConvertUnsignedNumberPtrToTerraformString[uint32]),
		conversion.NewTypeConverter(ConvertUnsignedNumberPtrToTerraformString[uint64]),

		conversion.NewTypeConverter(ConvertFloat32PtrToTerraformString),
		conversion.NewTypeConverter(ConvertFloat64PtrToTerraformString),

		conversion.NewTypeConverter(ConvertTerraformStringToUnsignedNumber[uint8]),
		conversion.NewTypeConverter(ConvertTerraformStringToUnsignedNumber[uint16]),
		conversion.NewTypeConverter(ConvertTerraformStringToUnsignedNumber[uint32]),
		conversion.NewTypeConverter(ConvertTerraformStringToUnsignedNumber[uint64]),
		conversion.NewTypeConverter(ConvertTerraformStringToSignedNumber[int8]),
		conversion.NewTypeConverter(ConvertTerraformStringToSignedNumber[int16]),
		conversion.NewTypeConverter(ConvertTerraformStringToSignedNumber[int32]),
		conversion.NewTypeConverter(ConvertTerraformStringToSignedNumber[int64]),

		conversion.NewTypeConverter(ConvertTerraformStringUnsignedNumberPtr[uint8]),
		conversion.NewTypeConverter(ConvertTerraformStringUnsignedNumberPtr[uint16]),
		conversion.NewTypeConverter(ConvertTerraformStringUnsignedNumberPtr[uint32]),
		conversion.NewTypeConverter(ConvertTerraformStringUnsignedNumberPtr[uint64]),
		conversion.NewTypeConverter(ConvertTerraformStringToSignedNumberPtr[int8]),
		conversion.NewTypeConverter(ConvertTerraformStringToSignedNumberPtr[int16]),
		conversion.NewTypeConverter(ConvertTerraformStringToSignedNumberPtr[int32]),
		conversion.NewTypeConverter(ConvertTerraformStringToSignedNumberPtr[int64]),

		conversion.NewTypeConverter(ConvertBoolToTerraformString),
		conversion.NewTypeConverter(ConvertBoolPtrToTerraformString),
		conversion.NewTypeConverter(ConvertTerraformStringToBool),
		conversion.NewTypeConverter(ConvertTerraformStringToBoolPtr),
	}
}
