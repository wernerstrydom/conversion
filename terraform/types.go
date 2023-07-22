package terraform

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/wernerstrydom/conversion"
)

func ConvertStringToTerraformString(source string) (types.String, error) {
	return types.StringValue(source), nil
}

func ConvertStringPtrToTerraformString(source *string) (types.String, error) {
	return types.StringPointerValue(source), nil
}

func ConvertTerraformStringToString(value types.String) (string, error) {
	return value.ValueString(), nil
}

func ConvertTerraformStringToStringPtr(value types.String) (*string, error) {
	return value.ValueStringPointer(), nil
}

func ConvertBoolToTerraformBool(source bool) (types.Bool, error) {
	return types.BoolValue(source), nil
}

func ConvertBoolPtrToTerraformBool(source *bool) (types.Bool, error) {
	return types.BoolPointerValue(source), nil
}

func ConvertTerraformBoolToBool(value types.Bool) (bool, error) {
	return value.ValueBool(), nil
}

func ConvertTerraformBoolToBoolPtr(value types.Bool) (*bool, error) {
	return value.ValueBoolPointer(), nil
}

// float64
func ConvertFloat64ToTerraformFloat64(source float64) (types.Float64, error) {
	return types.Float64Value(source), nil
}

func ConvertFloat64PtrToTerraformFloat64(source *float64) (types.Float64, error) {
	return types.Float64PointerValue(source), nil
}

func ConvertTerraformFloat64ToFloat64(value types.Float64) (float64, error) {
	return value.ValueFloat64(), nil
}

func ConvertTerraformFloat64ToFloat64Ptr(value types.Float64) (*float64, error) {
	return value.ValueFloat64Pointer(), nil
}

func ConvertInt64ToTerraformInt64(source int64) (types.Int64, error) {
	return types.Int64Value(source), nil
}

func ConvertInt64PtrToTerraformInt64(source *int64) (types.Int64, error) {
	return types.Int64PointerValue(source), nil
}

func ConvertTerraformInt64ToInt64(value types.Int64) (int64, error) {
	return value.ValueInt64(), nil
}

func ConvertTerraformInt64ToInt64Ptr(value types.Int64) (*int64, error) {
	return value.ValueInt64Pointer(), nil
}

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
	}
}
