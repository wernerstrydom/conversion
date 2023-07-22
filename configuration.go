package conversion

type Configuration struct {
	typeConverters []TypeConverter
}

func NewConfiguration() *Configuration {
	c := &Configuration{}
	c.RegisterDefaultConverters()
	return c
}

func (c *Configuration) RegisterTypeConverter(tc TypeConverter) {
	c.typeConverters = append(c.typeConverters, tc)
}

func (c *Configuration) RegisterTypeConverters(tcs []TypeConverter) {
	c.typeConverters = append(c.typeConverters, tcs...)
}

func (c *Configuration) DeregisterTypeConverter(tc TypeConverter) {
	for i, t := range c.typeConverters {
		if t == tc {
			c.typeConverters = append(c.typeConverters[:i], c.typeConverters[i+1:]...)
			return
		}
	}
}

func (c *Configuration) Clear() {
	c.typeConverters = []TypeConverter{}
}

func (c *Configuration) RegisterDefaultConverters() {
	c.RegisterTypeConverter(BoolToBoolConverter)
	c.RegisterTypeConverter(BoolToUintConverter)
	c.RegisterTypeConverter(BoolToUint8Converter)
	c.RegisterTypeConverter(BoolToUint16Converter)
	c.RegisterTypeConverter(BoolToUint32Converter)
	c.RegisterTypeConverter(BoolToUint64Converter)
	c.RegisterTypeConverter(BoolToIntConverter)
	c.RegisterTypeConverter(BoolToInt8Converter)
	c.RegisterTypeConverter(BoolToInt16Converter)
	c.RegisterTypeConverter(BoolToInt32Converter)
	c.RegisterTypeConverter(BoolToInt64Converter)
	c.RegisterTypeConverter(BoolToFloat32Converter)
	c.RegisterTypeConverter(BoolToFloat64Converter)
	c.RegisterTypeConverter(BoolToStringConverter)

	c.RegisterTypeConverter(IntToBoolConverter)
	c.RegisterTypeConverter(IntToUintConverter)
	c.RegisterTypeConverter(IntToUint8Converter)
	c.RegisterTypeConverter(IntToUint16Converter)
	c.RegisterTypeConverter(IntToUint32Converter)
	c.RegisterTypeConverter(IntToUint64Converter)
	c.RegisterTypeConverter(IntToIntConverter)
	c.RegisterTypeConverter(IntToInt8Converter)
	c.RegisterTypeConverter(IntToInt16Converter)
	c.RegisterTypeConverter(IntToInt32Converter)
	c.RegisterTypeConverter(IntToInt64Converter)
	c.RegisterTypeConverter(IntToFloat32Converter)
	c.RegisterTypeConverter(IntToFloat64Converter)
	c.RegisterTypeConverter(IntToStringConverter)

	c.RegisterTypeConverter(Int8ToBoolConverter)
	c.RegisterTypeConverter(Int8ToUintConverter)
	c.RegisterTypeConverter(Int8ToUint8Converter)
	c.RegisterTypeConverter(Int8ToUint16Converter)
	c.RegisterTypeConverter(Int8ToUint32Converter)
	c.RegisterTypeConverter(Int8ToUint64Converter)
	c.RegisterTypeConverter(Int8ToIntConverter)
	c.RegisterTypeConverter(Int8ToInt8Converter)
	c.RegisterTypeConverter(Int8ToInt16Converter)
	c.RegisterTypeConverter(Int8ToInt32Converter)
	c.RegisterTypeConverter(Int8ToInt64Converter)
	c.RegisterTypeConverter(Int8ToFloat32Converter)
	c.RegisterTypeConverter(Int8ToFloat64Converter)
	c.RegisterTypeConverter(Int8ToStringConverter)

	c.RegisterTypeConverter(Int16ToBoolConverter)
	c.RegisterTypeConverter(Int16ToUintConverter)
	c.RegisterTypeConverter(Int16ToUint8Converter)
	c.RegisterTypeConverter(Int16ToUint16Converter)
	c.RegisterTypeConverter(Int16ToUint32Converter)
	c.RegisterTypeConverter(Int16ToUint64Converter)
	c.RegisterTypeConverter(Int16ToIntConverter)
	c.RegisterTypeConverter(Int16ToInt8Converter)
	c.RegisterTypeConverter(Int16ToInt16Converter)
	c.RegisterTypeConverter(Int16ToInt32Converter)
	c.RegisterTypeConverter(Int16ToInt64Converter)
	c.RegisterTypeConverter(Int16ToFloat32Converter)
	c.RegisterTypeConverter(Int16ToFloat64Converter)
	c.RegisterTypeConverter(Int16ToStringConverter)

	c.RegisterTypeConverter(Int32ToBoolConverter)
	c.RegisterTypeConverter(Int32ToUintConverter)
	c.RegisterTypeConverter(Int32ToUint8Converter)
	c.RegisterTypeConverter(Int32ToUint16Converter)
	c.RegisterTypeConverter(Int32ToUint32Converter)
	c.RegisterTypeConverter(Int32ToUint64Converter)
	c.RegisterTypeConverter(Int32ToIntConverter)
	c.RegisterTypeConverter(Int32ToInt8Converter)
	c.RegisterTypeConverter(Int32ToInt16Converter)
	c.RegisterTypeConverter(Int32ToInt32Converter)
	c.RegisterTypeConverter(Int32ToInt64Converter)
	c.RegisterTypeConverter(Int32ToFloat32Converter)
	c.RegisterTypeConverter(Int32ToFloat64Converter)
	c.RegisterTypeConverter(Int32ToStringConverter)

	c.RegisterTypeConverter(Int64ToBoolConverter)
	c.RegisterTypeConverter(Int64ToUintConverter)
	c.RegisterTypeConverter(Int64ToUint8Converter)
	c.RegisterTypeConverter(Int64ToUint16Converter)
	c.RegisterTypeConverter(Int64ToUint32Converter)
	c.RegisterTypeConverter(Int64ToUint64Converter)
	c.RegisterTypeConverter(Int64ToIntConverter)
	c.RegisterTypeConverter(Int64ToInt8Converter)
	c.RegisterTypeConverter(Int64ToInt16Converter)
	c.RegisterTypeConverter(Int64ToInt32Converter)
	c.RegisterTypeConverter(Int64ToInt64Converter)
	c.RegisterTypeConverter(Int64ToFloat32Converter)
	c.RegisterTypeConverter(Int64ToFloat64Converter)
	c.RegisterTypeConverter(Int64ToStringConverter)

	c.RegisterTypeConverter(UintToBoolConverter)
	c.RegisterTypeConverter(UintToUintConverter)
	c.RegisterTypeConverter(UintToUint8Converter)
	c.RegisterTypeConverter(UintToUint16Converter)
	c.RegisterTypeConverter(UintToUint32Converter)
	c.RegisterTypeConverter(UintToUint64Converter)
	c.RegisterTypeConverter(UintToIntConverter)
	c.RegisterTypeConverter(UintToInt8Converter)
	c.RegisterTypeConverter(UintToInt16Converter)
	c.RegisterTypeConverter(UintToInt32Converter)
	c.RegisterTypeConverter(UintToInt64Converter)
	c.RegisterTypeConverter(UintToFloat32Converter)
	c.RegisterTypeConverter(UintToFloat64Converter)
	c.RegisterTypeConverter(UintToStringConverter)

	c.RegisterTypeConverter(Uint8ToBoolConverter)
	c.RegisterTypeConverter(Uint8ToUintConverter)
	c.RegisterTypeConverter(Uint8ToUint8Converter)
	c.RegisterTypeConverter(Uint8ToUint16Converter)
	c.RegisterTypeConverter(Uint8ToUint32Converter)
	c.RegisterTypeConverter(Uint8ToUint64Converter)
	c.RegisterTypeConverter(Uint8ToIntConverter)
	c.RegisterTypeConverter(Uint8ToInt8Converter)
	c.RegisterTypeConverter(Uint8ToInt16Converter)
	c.RegisterTypeConverter(Uint8ToInt32Converter)
	c.RegisterTypeConverter(Uint8ToInt64Converter)
	c.RegisterTypeConverter(Uint8ToFloat32Converter)
	c.RegisterTypeConverter(Uint8ToFloat64Converter)
	c.RegisterTypeConverter(Uint8ToStringConverter)

	c.RegisterTypeConverter(Uint16ToBoolConverter)
	c.RegisterTypeConverter(Uint16ToUintConverter)
	c.RegisterTypeConverter(Uint16ToUint8Converter)
	c.RegisterTypeConverter(Uint16ToUint16Converter)
	c.RegisterTypeConverter(Uint16ToUint32Converter)
	c.RegisterTypeConverter(Uint16ToUint64Converter)
	c.RegisterTypeConverter(Uint16ToIntConverter)
	c.RegisterTypeConverter(Uint16ToInt8Converter)
	c.RegisterTypeConverter(Uint16ToInt16Converter)
	c.RegisterTypeConverter(Uint16ToInt32Converter)
	c.RegisterTypeConverter(Uint16ToInt64Converter)
	c.RegisterTypeConverter(Uint16ToFloat32Converter)
	c.RegisterTypeConverter(Uint16ToFloat64Converter)
	c.RegisterTypeConverter(Uint16ToStringConverter)

	c.RegisterTypeConverter(Uint32ToBoolConverter)
	c.RegisterTypeConverter(Uint32ToUintConverter)
	c.RegisterTypeConverter(Uint32ToUint8Converter)
	c.RegisterTypeConverter(Uint32ToUint16Converter)
	c.RegisterTypeConverter(Uint32ToUint32Converter)
	c.RegisterTypeConverter(Uint32ToUint64Converter)
	c.RegisterTypeConverter(Uint32ToIntConverter)
	c.RegisterTypeConverter(Uint32ToInt8Converter)
	c.RegisterTypeConverter(Uint32ToInt16Converter)
	c.RegisterTypeConverter(Uint32ToInt32Converter)
	c.RegisterTypeConverter(Uint32ToInt64Converter)
	c.RegisterTypeConverter(Uint32ToFloat32Converter)
	c.RegisterTypeConverter(Uint32ToFloat64Converter)
	c.RegisterTypeConverter(Uint32ToStringConverter)

	c.RegisterTypeConverter(Uint64ToBoolConverter)
	c.RegisterTypeConverter(Uint64ToUintConverter)
	c.RegisterTypeConverter(Uint64ToUint8Converter)
	c.RegisterTypeConverter(Uint64ToUint16Converter)
	c.RegisterTypeConverter(Uint64ToUint32Converter)
	c.RegisterTypeConverter(Uint64ToUint64Converter)
	c.RegisterTypeConverter(Uint64ToIntConverter)
	c.RegisterTypeConverter(Uint64ToInt8Converter)
	c.RegisterTypeConverter(Uint64ToInt16Converter)
	c.RegisterTypeConverter(Uint64ToInt32Converter)
	c.RegisterTypeConverter(Uint64ToInt64Converter)
	c.RegisterTypeConverter(Uint64ToFloat32Converter)
	c.RegisterTypeConverter(Uint64ToFloat64Converter)
	c.RegisterTypeConverter(Uint64ToStringConverter)

	c.RegisterTypeConverter(Float32ToBoolConverter)
	c.RegisterTypeConverter(Float32ToUintConverter)
	c.RegisterTypeConverter(Float32ToUint8Converter)
	c.RegisterTypeConverter(Float32ToUint16Converter)
	c.RegisterTypeConverter(Float32ToUint32Converter)
	c.RegisterTypeConverter(Float32ToUint64Converter)
	c.RegisterTypeConverter(Float32ToIntConverter)
	c.RegisterTypeConverter(Float32ToInt8Converter)
	c.RegisterTypeConverter(Float32ToInt16Converter)
	c.RegisterTypeConverter(Float32ToInt32Converter)
	c.RegisterTypeConverter(Float32ToInt64Converter)
	c.RegisterTypeConverter(Float32ToFloat32Converter)
	c.RegisterTypeConverter(Float32ToFloat64Converter)
	c.RegisterTypeConverter(Float32ToStringConverter)

	c.RegisterTypeConverter(Float64ToBoolConverter)
	c.RegisterTypeConverter(Float64ToUintConverter)
	c.RegisterTypeConverter(Float64ToUint8Converter)
	c.RegisterTypeConverter(Float64ToUint16Converter)
	c.RegisterTypeConverter(Float64ToUint32Converter)
	c.RegisterTypeConverter(Float64ToUint64Converter)
	c.RegisterTypeConverter(Float64ToIntConverter)
	c.RegisterTypeConverter(Float64ToInt8Converter)
	c.RegisterTypeConverter(Float64ToInt16Converter)
	c.RegisterTypeConverter(Float64ToInt32Converter)
	c.RegisterTypeConverter(Float64ToInt64Converter)
	c.RegisterTypeConverter(Float64ToFloat32Converter)
	c.RegisterTypeConverter(Float64ToFloat32Converter)
	c.RegisterTypeConverter(Float64ToStringConverter)

	c.RegisterTypeConverter(StringToBoolConverter)
	c.RegisterTypeConverter(StringToUintConverter)
	c.RegisterTypeConverter(StringToUint8Converter)
	c.RegisterTypeConverter(StringToUint16Converter)
	c.RegisterTypeConverter(StringToUint32Converter)
	c.RegisterTypeConverter(StringToUint64Converter)
	c.RegisterTypeConverter(StringToIntConverter)
	c.RegisterTypeConverter(StringToInt8Converter)
	c.RegisterTypeConverter(StringToInt16Converter)
	c.RegisterTypeConverter(StringToInt32Converter)
	c.RegisterTypeConverter(StringToInt64Converter)
	c.RegisterTypeConverter(StringToFloat32Converter)
	c.RegisterTypeConverter(StringToFloat64Converter)
	c.RegisterTypeConverter(StringToStringConverter)

	c.RegisterTypeConverter(StructTypeConverter)
}

func (c *Configuration) BuildMapper() Mapper {
	// Type converters were added to the configuration with the most recent added converter last. However,
	// we want the converters the user of the library added to be the first ones we try. So, we reverse
	// the order of the converters.

	// make a copy of the type converters
	typeConverters := make([]TypeConverter, len(c.typeConverters))
	copy(typeConverters, c.typeConverters)

	// reverse the order of the type converters
	for i, j := 0, len(typeConverters)-1; i < j; i, j = i+1, j-1 {
		typeConverters[i], typeConverters[j] = typeConverters[j], typeConverters[i]
	}

	return &mapper{
		typeConverters,
	}
}
