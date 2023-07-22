package conversion

import (
	"math"
	"strconv"
)

// A table of all the primitive types in Go, and the primitive types they can be mapped to
//
// | Source Type  | Destination Type |
// | ------------ | ---------------- |
// | bool         | bool             |
// | bool         | int              |
// | bool         | int8             |
// | bool         | int16            |
// | bool         | int32            |
// | bool         | int64            |
// | bool         | uint             |
// | bool         | uint8            |
// | bool         | uint16           |
// | bool         | uint32           |
// | bool         | uint64           |
// | bool         | float32          |
// | bool         | float64          |
// | bool         | string           |
// | int          | bool             |
// | int          | int              |
// | int          | int8             |
// | int          | int16            |
// | int          | int32            |
// | int          | int64            |
// | int          | uint             |
// | int          | uint8            |
// | int          | uint16           |
// | int          | uint32           |
// | int          | uint64           |
// | int          | float32          |
// | int          | float64          |
// | int          | string           |
// | int8         | bool             |
// | int8         | int              |
// | int8         | int8             |
// | int8         | int16            |
// | int8         | int32            |
// | int8         | int64            |
// | int8         | uint             |
// | int8         | uint8            |
// | int8         | uint16           |
// | int8         | uint32           |
// | int8         | uint64           |
// | int8         | float32          |
// | int8         | float64          |
// | int8         | string           |
// | int16        | bool             |
// | int16        | int              |
// | int16        | int8             |
// | int16        | int16            |
// | int16        | int32            |
// | int16        | int64            |
// | int16        | uint             |
// | int16        | uint8            |
// | int16        | uint16           |
// | int16        | uint32           |
// | int16        | uint64           |
// | int16        | float32          |
// | int16        | float64          |
// | int16        | string           |
// | int32        | bool             |
// | int32        | int              |
// | int32        | int8             |
// | int32        | int16            |
// | int32        | int32            |
// | int32        | int64            |
// | int32        | uint             |
// | int32        | uint8            |
// | int32        | uint16           |
// | int32        | uint32           |
// | int32        | uint64           |
// | int32        | float32          |
// | int32        | float64          |
// | int32        | string           |
// | int64        | bool             |
// | int64        | int              |
// | int64        | int8             |
// | int64        | int16            |
// | int64        | int32            |
// | int64        | int64            |
// | int64        | uint             |
// | int64        | uint8            |
// | int64        | uint16           |
// | int64        | uint32           |
// | int64        | uint64           |
// | int64        | float32          |
// | int64        | float64          |
// | int64        | string           |
// | uint         | bool             |
// | uint         | int              |
// | uint         | int8             |
// | uint         | int16            |
// | uint         | int32            |
// | uint         | int64            |
// | uint         | uint             |
// | uint         | uint8            |
// | uint         | uint16           |
// | uint         | uint32           |
// | uint         | uint64           |
// | uint         | float32          |
// | uint         | float64          |
// | uint         | string           |
// | uint8        | bool             |
// | uint8        | int              |
// | uint8        | int8             |
// | uint8        | int16            |
// | uint8        | int32            |
// | uint8        | int64            |
// | uint8        | uint             |
// | uint8        | uint8            |
// | uint8        | uint16           |
// | uint8        | uint32           |
// | uint8        | uint64           |
// | uint8        | float32          |
// | uint8        | float64          |
// | uint8        | string           |
// | uint16       | bool             |
// | uint16       | int              |
// | uint16       | int8             |
// | uint16       | int16            |
// | uint16       | int32            |
// | uint16       | int64            |
// | uint16       | uint             |
// | uint16       | uint8            |
// | uint16       | uint16           |
// | uint16       | uint32           |
// | uint16       | uint64           |
// | uint16       | float32          |
// | uint16       | float64          |
// | uint16       | string           |
// | uint32       | bool             |
// | uint32       | int              |
// | uint32       | int8             |
// | uint32       | int16            |
// | uint32       | int32            |
// | uint32       | int64            |
// | uint32       | uint             |
// | uint32       | uint8            |
// | uint32       | uint16           |
// | uint32       | uint32           |
// | uint32       | uint64           |
// | uint32       | float32          |
// | uint32       | float64          |
// | uint32       | string           |
// | uint64       | bool             |
// | uint64       | int              |
// | uint64       | int8             |
// | uint64       | int16            |
// | uint64       | int32            |
// | uint64       | int64            |
// | uint64       | uint             |
// | uint64       | uint8            |
// | uint64       | uint16           |
// | uint64       | uint32           |
// | uint64       | uint64           |
// | uint64       | float32          |
// | uint64       | float64          |
// | uint64       | string           |
// | float32      | bool             |
// | float32      | int              |
// | float32      | int8             |
// | float32      | int16            |
// | float32      | int32            |
// | float32      | int64            |
// | float32      | uint             |
// | float32      | uint8            |
// | float32      | uint16           |
// | float32      | uint32           |
// | float32      | uint64           |
// | float32      | float32          |
// | float32      | float64          |
// | float32      | string           |
// | float64      | bool             |
// | float64      | int              |
// | float64      | int8             |
// | float64      | int16            |
// | float64      | int32            |
// | float64      | int64            |
// | float64      | uint             |
// | float64      | uint8            |
// | float64      | uint16           |
// | float64      | uint32           |
// | float64      | uint64           |
// | float64      | float32          |
// | float64      | float64          |
// | float64      | string           |
// | string       | bool             |
// | string       | int              |
// | string       | int8             |
// | string       | int16            |
// | string       | int32            |
// | string       | int64            |
// | string       | uint             |
// | string       | uint8            |
// | string       | uint16           |
// | string       | uint32           |
// | string       | uint64           |
// | string       | float32          |
// | string       | float64          |
// | string       | string           |

// ----------------------------------------------------------------------------
// Bool
// ----------------------------------------------------------------------------

func ConvertBoolToBool(v bool) (bool, error) {
	return v, nil
}

func ConvertBoolToNumber[T Number](v bool) (T, error) {
	if v {
		return 1, nil
	}
	return 0, nil
}

func ConvertBoolToString(v bool) (string, error) {
	if v {
		return "true", nil
	}
	return "false", nil
}

func ConvertNumberToBool[T Number](v T) (bool, error) {
	return v != 0, nil
}

func ConvertStringToBool(v string) (bool, error) {
	switch v {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, NewFormatError(v, "true or false")
	}
}

// ----------------------------------------------------------------------------
// Int
// ----------------------------------------------------------------------------

func ConvertIntToInt(v int) (int, error) {
	return v, nil
}

func ConvertIntToInt8(v int) (int8, error) {
	if v < math.MinInt8 || v > math.MaxInt8 {
		return 0, NewOverflowError(v, math.MinInt8, math.MaxInt8)
	}
	return int8(v), nil
}

func ConvertIntToInt16(v int) (int16, error) {
	if v < math.MinInt16 || v > math.MaxInt16 {
		return 0, NewOverflowError(v, math.MinInt16, math.MaxInt16)
	}
	return int16(v), nil
}

func ConvertIntToInt32(v int) (int32, error) {
	if v < math.MinInt32 || v > math.MaxInt32 {
		return 0, NewOverflowError(v, math.MinInt32, math.MaxInt32)
	}
	return int32(v), nil
}

func ConvertIntToInt64(v int) (int64, error) {
	return int64(v), nil
}

func ConvertIntToUint(v int) (uint, error) {
	if v < 0 {
		return 0, NewOverflowError(v, 0, uint(math.MaxUint))
	}
	return uint(v), nil
}

func ConvertIntToUint8(v int) (uint8, error) {
	if v < 0 || v > math.MaxUint8 {
		return 0, NewOverflowError(v, 0, uint8(math.MaxUint8))
	}
	return uint8(v), nil
}

func ConvertIntToUint16(v int) (uint16, error) {
	if v < 0 || v > math.MaxUint16 {
		return 0, NewOverflowError(v, 0, uint16(math.MaxUint16))
	}
	return uint16(v), nil
}

func ConvertIntToUint32(v int) (uint32, error) {
	if v < 0 {
		return 0, NewOverflowError(v, 0, uint32(math.MaxUint32))
	}
	return uint32(v), nil
}

func ConvertIntToUint64(v int) (uint64, error) {
	if v < 0 {
		return 0, NewOverflowError(v, 0, uint64(math.MaxUint64))
	}
	return uint64(v), nil
}

func ConvertIntToFloat32(v int) (float32, error) {
	return float32(v), nil
}

func ConvertIntToFloat64(v int) (float64, error) {
	return float64(v), nil
}

func ConvertIntToString(v int) (string, error) {
	return strconv.Itoa(v), nil
}

// ----------------------------------------------------------------------------
// Int8
// ----------------------------------------------------------------------------

func ConvertInt8ToInt(v int8) (int, error) {
	return int(v), nil
}

func ConvertInt8ToInt8(v int8) (int8, error) {
	return v, nil
}

func ConvertInt8ToInt16(v int8) (int16, error) {
	return int16(v), nil
}

func ConvertInt8ToInt32(v int8) (int32, error) {
	return int32(v), nil
}

func ConvertInt8ToInt64(v int8) (int64, error) {
	return int64(v), nil
}

func ConvertInt8ToUint(v int8) (uint, error) {
	if v < 0 {
		return 0, NewOverflowError(v, 0, uint(math.MaxUint))
	}
	return uint(v), nil
}

func ConvertInt8ToUint8(v int8) (uint8, error) {
	if v < 0 {
		return 0, NewOverflowError(v, 0, uint8(math.MaxUint8))
	}
	return uint8(v), nil
}

func ConvertInt8ToUint16(v int8) (uint16, error) {
	if v < 0 {
		return 0, NewOverflowError(v, 0, uint16(math.MaxUint16))
	}
	return uint16(v), nil
}

func ConvertInt8ToUint32(v int8) (uint32, error) {
	if v < 0 {
		return 0, NewOverflowError(v, 0, uint32(math.MaxUint32))
	}
	return uint32(v), nil
}

func ConvertInt8ToUint64(v int8) (uint64, error) {
	if v < 0 {
		return 0, NewOverflowError(v, 0, uint64(math.MaxUint64))
	}
	return uint64(v), nil
}

func ConvertInt8ToFloat32(v int8) (float32, error) {
	return float32(v), nil
}

func ConvertInt8ToFloat64(v int8) (float64, error) {
	return float64(v), nil
}

func ConvertInt8ToString(v int8) (string, error) {
	return strconv.FormatInt(int64(v), 10), nil
}

// ----------------------------------------------------------------------------
// Int16
// ----------------------------------------------------------------------------

func ConvertInt16ToInt(v int16) (int, error) {
	return int(v), nil
}

func ConvertInt16ToInt8(v int16) (int8, error) {
	if v < math.MinInt8 || v > math.MaxInt8 {
		return 0, NewOverflowError(v, math.MinInt8, math.MaxInt8)
	}
	return int8(v), nil
}

func ConvertInt16ToInt16(v int16) (int16, error) {
	return v, nil
}

func ConvertInt16ToInt32(v int16) (int32, error) {
	return int32(v), nil
}

func ConvertInt16ToInt64(v int16) (int64, error) {
	return int64(v), nil
}

func ConvertInt16ToUint(v int16) (uint, error) {
	if v < 0 {
		return 0, NewOverflowError(v, 0, uint(math.MaxUint))
	}
	return uint(v), nil
}

func ConvertInt16ToUint8(v int16) (uint8, error) {
	if v < 0 || v > math.MaxUint8 {
		return 0, NewOverflowError(v, 0, uint8(math.MaxUint8))
	}
	return uint8(v), nil
}

func ConvertInt16ToUint16(v int16) (uint16, error) {
	if v < 0 {
		return 0, NewOverflowError(v, 0, uint16(math.MaxUint16))
	}
	return uint16(v), nil
}

func ConvertInt16ToUint32(v int16) (uint32, error) {
	if v < 0 {
		return 0, NewOverflowError(v, 0, uint32(math.MaxUint32))
	}
	return uint32(v), nil
}

func ConvertInt16ToUint64(v int16) (uint64, error) {
	if v < 0 {
		return 0, NewOverflowError(v, 0, uint64(math.MaxUint64))
	}
	return uint64(v), nil
}

func ConvertInt16ToFloat32(v int16) (float32, error) {
	return float32(v), nil
}

func ConvertInt16ToFloat64(v int16) (float64, error) {
	return float64(v), nil
}

func ConvertInt16ToString(v int16) (string, error) {
	return strconv.FormatInt(int64(v), 10), nil
}

// ----------------------------------------------------------------------------
// Int32
// ----------------------------------------------------------------------------

func ConvertInt32ToInt(v int32) (int, error) {
	return int(v), nil
}

func ConvertInt32ToInt8(v int32) (int8, error) {
	if v < math.MinInt8 || v > math.MaxInt8 {
		return 0, NewOverflowError(v, math.MinInt8, math.MaxInt8)
	}
	return int8(v), nil
}

func ConvertInt32ToInt16(v int32) (int16, error) {
	if v < math.MinInt16 || v > math.MaxInt16 {
		return 0, NewOverflowError(v, math.MinInt16, math.MaxInt16)
	}
	return int16(v), nil
}

func ConvertInt32ToInt32(v int32) (int32, error) {
	return v, nil
}

func ConvertInt32ToInt64(v int32) (int64, error) {
	return int64(v), nil
}

func ConvertInt32ToUint(v int32) (uint, error) {
	if v < 0 {
		return 0, NewOverflowError(v, 0, uint(math.MaxUint))
	}
	return uint(v), nil
}

func ConvertInt32ToUint8(v int32) (uint8, error) {
	if v < 0 || v > math.MaxUint8 {
		return 0, NewOverflowError(v, 0, uint8(math.MaxUint8))
	}
	return uint8(v), nil
}

func ConvertInt32ToUint16(v int32) (uint16, error) {
	if v < 0 || v > math.MaxUint16 {
		return 0, NewOverflowError(v, 0, uint16(math.MaxUint16))
	}
	return uint16(v), nil
}

func ConvertInt32ToUint32(v int32) (uint32, error) {
	if v < 0 {
		return 0, NewOverflowError(v, 0, uint32(math.MaxUint32))
	}
	return uint32(v), nil
}

func ConvertInt32ToUint64(v int32) (uint64, error) {
	if v < 0 {
		return 0, NewOverflowError(v, 0, uint64(math.MaxUint64))
	}
	return uint64(v), nil
}

func ConvertInt32ToFloat32(v int32) (float32, error) {
	return float32(v), nil
}

func ConvertInt32ToFloat64(v int32) (float64, error) {
	return float64(v), nil
}

func ConvertInt32ToString(v int32) (string, error) {
	return strconv.FormatInt(int64(v), 10), nil
}

// ----------------------------------------------------------------------------
// Int64
// ----------------------------------------------------------------------------

func ConvertInt64ToInt(v int64) (int, error) {
	if v < math.MinInt32 || v > math.MaxInt32 {
		return 0, NewOverflowError(v, math.MinInt32, math.MaxInt32)
	}
	return int(v), nil
}

func ConvertInt64ToInt8(v int64) (int8, error) {
	if v < math.MinInt8 || v > math.MaxInt8 {
		return 0, NewOverflowError(v, math.MinInt8, math.MaxInt8)
	}
	return int8(v), nil
}

func ConvertInt64ToInt16(v int64) (int16, error) {
	if v < math.MinInt16 || v > math.MaxInt16 {
		return 0, NewOverflowError(v, math.MinInt16, math.MaxInt16)
	}
	return int16(v), nil
}

func ConvertInt64ToInt32(v int64) (int32, error) {
	if v < math.MinInt32 || v > math.MaxInt32 {
		return 0, NewOverflowError(v, math.MinInt32, math.MaxInt32)
	}
	return int32(v), nil
}

func ConvertInt64ToInt64(v int64) (int64, error) {
	return v, nil
}

func ConvertInt64ToUint(v int64) (uint, error) {
	if v < 0 {
		return 0, NewOverflowError(v, 0, uint(math.MaxUint))
	}
	return uint(v), nil
}

func ConvertInt64ToUint8(v int64) (uint8, error) {
	if v < 0 || v > math.MaxUint8 {
		return 0, NewOverflowError(v, 0, uint8(math.MaxUint8))
	}
	return uint8(v), nil
}

func ConvertInt64ToUint16(v int64) (uint16, error) {
	if v < 0 || v > math.MaxUint16 {
		return 0, NewOverflowError(v, 0, uint16(math.MaxUint16))
	}
	return uint16(v), nil
}

func ConvertInt64ToUint32(v int64) (uint32, error) {
	if v < 0 || v > math.MaxUint32 {
		return 0, NewOverflowError(v, 0, uint32(math.MaxUint32))
	}
	return uint32(v), nil
}

func ConvertInt64ToUint64(v int64) (uint64, error) {
	if v < 0 {
		return 0, NewOverflowError(v, 0, uint64(math.MaxUint64))
	}
	return uint64(v), nil
}

func ConvertInt64ToFloat32(v int64) (float32, error) {
	return float32(v), nil
}

func ConvertInt64ToFloat64(v int64) (float64, error) {
	return float64(v), nil
}

func ConvertInt64ToString(v int64) (string, error) {
	return strconv.FormatInt(v, 10), nil
}

// ----------------------------------------------------------------------------
// Uint
// ----------------------------------------------------------------------------

func ConvertUintToInt(v uint) (int, error) {
	if v > math.MaxInt32 {
		return 0, NewOverflowError(v, 0, math.MaxInt32)
	}
	return int(v), nil
}

func ConvertUintToInt8(v uint) (int8, error) {
	if v > math.MaxInt8 {
		return 0, NewOverflowError(v, 0, math.MaxInt8)
	}
	return int8(v), nil
}

func ConvertUintToInt16(v uint) (int16, error) {
	if v > math.MaxInt16 {
		return 0, NewOverflowError(v, 0, math.MaxInt16)
	}
	return int16(v), nil
}

func ConvertUintToInt32(v uint) (int32, error) {
	if v > math.MaxInt32 {
		return 0, NewOverflowError(v, 0, math.MaxInt32)
	}
	return int32(v), nil
}

func ConvertUintToInt64(v uint) (int64, error) {
	if v > math.MaxInt64 {
		return 0, NewOverflowError(v, 0, math.MaxInt64)
	}
	return int64(v), nil
}

func ConvertUintToUint(v uint) (uint, error) {
	return v, nil
}

func ConvertUintToUint8(v uint) (uint8, error) {
	if v > math.MaxUint8 {
		return 0, NewOverflowError(v, 0, uint8(math.MaxUint8))
	}
	return uint8(v), nil
}

func ConvertUintToUint16(v uint) (uint16, error) {
	if v > math.MaxUint16 {
		return 0, NewOverflowError(v, 0, uint16(math.MaxUint16))
	}
	return uint16(v), nil
}

func ConvertUintToUint32(v uint) (uint32, error) {
	if v > math.MaxUint32 {
		return 0, NewOverflowError(v, 0, uint32(math.MaxUint32))
	}
	return uint32(v), nil
}

func ConvertUintToUint64(v uint) (uint64, error) {
	return uint64(v), nil
}

func ConvertUintToFloat32(v uint) (float32, error) {
	return float32(v), nil
}

func ConvertUintToFloat64(v uint) (float64, error) {
	return float64(v), nil
}

func ConvertUintToString(v uint) (string, error) {
	return strconv.FormatUint(uint64(v), 10), nil
}

// ----------------------------------------------------------------------------
// Uint8
// ----------------------------------------------------------------------------

func ConvertUint8ToInt(v uint8) (int, error) {
	return int(v), nil
}

func ConvertUint8ToInt8(v uint8) (int8, error) {
	if v > math.MaxInt8 {
		return 0, NewOverflowError(v, 0, math.MaxInt8)
	}
	return int8(v), nil
}

func ConvertUint8ToInt16(v uint8) (int16, error) {
	return int16(v), nil
}

func ConvertUint8ToInt32(v uint8) (int32, error) {
	return int32(v), nil
}

func ConvertUint8ToInt64(v uint8) (int64, error) {
	return int64(v), nil
}

func ConvertUint8ToUint(v uint8) (uint, error) {
	return uint(v), nil
}

func ConvertUint8ToUint8(v uint8) (uint8, error) {
	return v, nil
}

func ConvertUint8ToUint16(v uint8) (uint16, error) {
	return uint16(v), nil
}

func ConvertUint8ToUint32(v uint8) (uint32, error) {
	return uint32(v), nil
}

func ConvertUint8ToUint64(v uint8) (uint64, error) {
	return uint64(v), nil
}

func ConvertUint8ToFloat32(v uint8) (float32, error) {
	return float32(v), nil
}

func ConvertUint8ToFloat64(v uint8) (float64, error) {
	return float64(v), nil
}

func ConvertUint8ToString(v uint8) (string, error) {
	return strconv.FormatUint(uint64(v), 10), nil
}

// ----------------------------------------------------------------------------
// Uint16
// ----------------------------------------------------------------------------

func ConvertUint16ToInt(v uint16) (int, error) {
	return int(v), nil
}

func ConvertUint16ToInt8(v uint16) (int8, error) {
	if v > math.MaxInt8 {
		return 0, NewOverflowError(v, 0, math.MaxInt8)
	}
	return int8(v), nil
}

func ConvertUint16ToInt16(v uint16) (int16, error) {
	if v > math.MaxInt16 {
		return 0, NewOverflowError(v, 0, math.MaxInt16)
	}
	return int16(v), nil
}

func ConvertUint16ToInt32(v uint16) (int32, error) {
	return int32(v), nil
}

func ConvertUint16ToInt64(v uint16) (int64, error) {
	return int64(v), nil
}

func ConvertUint16ToUint(v uint16) (uint, error) {
	return uint(v), nil
}

func ConvertUint16ToUint8(v uint16) (uint8, error) {
	return uint8(v), nil
}

func ConvertUint16ToUint16(v uint16) (uint16, error) {
	return v, nil
}

func ConvertUint16ToUint32(v uint16) (uint32, error) {
	return uint32(v), nil
}

func ConvertUint16ToUint64(v uint16) (uint64, error) {
	return uint64(v), nil
}

func ConvertUint16ToFloat32(v uint16) (float32, error) {
	return float32(v), nil
}

func ConvertUint16ToFloat64(v uint16) (float64, error) {
	return float64(v), nil
}

func ConvertUint16ToString(v uint16) (string, error) {
	return strconv.FormatUint(uint64(v), 10), nil
}

// ----------------------------------------------------------------------------
// Uint32
// ----------------------------------------------------------------------------

func ConvertUint32ToInt(v uint32) (int, error) {
	if v > math.MaxInt32 {
		return 0, NewOverflowError(v, 0, math.MaxInt32)
	}
	return int(v), nil
}

func ConvertUint32ToInt8(v uint32) (int8, error) {
	if v > math.MaxInt8 {
		return 0, NewOverflowError(v, 0, math.MaxInt8)
	}
	return int8(v), nil
}

func ConvertUint32ToInt16(v uint32) (int16, error) {
	if v > math.MaxInt16 {
		return 0, NewOverflowError(v, 0, math.MaxInt16)
	}
	return int16(v), nil
}

func ConvertUint32ToInt32(v uint32) (int32, error) {
	if v > math.MaxInt32 {
		return 0, NewOverflowError(v, 0, math.MaxInt32)
	}
	return int32(v), nil
}

func ConvertUint32ToInt64(v uint32) (int64, error) {
	return int64(v), nil
}

func ConvertUint32ToUint(v uint32) (uint, error) {
	return uint(v), nil
}

func ConvertUint32ToUint8(v uint32) (uint8, error) {
	if v > math.MaxUint8 {
		return 0, NewOverflowError(v, 0, uint8(math.MaxUint8))
	}
	return uint8(v), nil
}

func ConvertUint32ToUint16(v uint32) (uint16, error) {
	if v > math.MaxUint16 {
		return 0, NewOverflowError(v, 0, uint16(math.MaxUint16))
	}
	return uint16(v), nil
}

func ConvertUint32ToUint32(v uint32) (uint32, error) {
	return v, nil
}

func ConvertUint32ToUint64(v uint32) (uint64, error) {
	return uint64(v), nil
}

func ConvertUint32ToFloat32(v uint32) (float32, error) {
	return float32(v), nil
}

func ConvertUint32ToFloat64(v uint32) (float64, error) {
	return float64(v), nil
}

func ConvertUint32ToString(v uint32) (string, error) {
	return strconv.FormatUint(uint64(v), 10), nil
}

// ----------------------------------------------------------------------------
// Uint64
// ----------------------------------------------------------------------------

func ConvertUint64ToInt(v uint64) (int, error) {
	if v > math.MaxInt64 {
		return 0, NewOverflowError(v, 0, math.MaxInt64)
	}
	return int(v), nil
}

func ConvertUint64ToInt8(v uint64) (int8, error) {
	if v > math.MaxInt8 {
		return 0, NewOverflowError(v, 0, math.MaxInt8)
	}
	return int8(v), nil
}

func ConvertUint64ToInt16(v uint64) (int16, error) {
	if v > math.MaxInt16 {
		return 0, NewOverflowError(v, 0, math.MaxInt16)
	}
	return int16(v), nil
}

func ConvertUint64ToInt32(v uint64) (int32, error) {
	if v > math.MaxInt32 {
		return 0, NewOverflowError(v, 0, math.MaxInt32)
	}
	return int32(v), nil
}

func ConvertUint64ToInt64(v uint64) (int64, error) {
	if v > math.MaxInt64 {
		return 0, NewOverflowError(v, 0, math.MaxInt64)
	}
	return int64(v), nil
}

func ConvertUint64ToUint(v uint64) (uint, error) {
	if v > math.MaxUint {
		return 0, NewOverflowError(v, 0, uint64(math.MaxUint))
	}
	return uint(v), nil
}

func ConvertUint64ToUint8(v uint64) (uint8, error) {
	if v > math.MaxUint8 {
		return 0, NewOverflowError(v, 0, uint8(math.MaxUint8))
	}
	return uint8(v), nil
}

func ConvertUint64ToUint16(v uint64) (uint16, error) {
	if v > math.MaxUint16 {
		return 0, NewOverflowError(v, 0, uint16(math.MaxUint16))
	}
	return uint16(v), nil
}

func ConvertUint64ToUint32(v uint64) (uint32, error) {
	if v > math.MaxUint32 {
		return 0, NewOverflowError(v, 0, uint32(math.MaxUint32))
	}
	return uint32(v), nil
}

func ConvertUint64ToUint64(v uint64) (uint64, error) {
	return v, nil
}

func ConvertUint64ToFloat32(v uint64) (float32, error) {
	return float32(v), nil
}

func ConvertUint64ToFloat64(v uint64) (float64, error) {
	return float64(v), nil
}

func ConvertUint64ToString(v uint64) (string, error) {
	return strconv.FormatUint(uint64(v), 10), nil
}

// ----------------------------------------------------------------------------
// Float32
// ----------------------------------------------------------------------------

func ConvertFloat32ToInt(v float32) (int, error) {
	if v > math.MaxInt32 || v < math.MinInt32 {
		return 0, NewOverflowError(v, math.MinInt32, math.MaxInt32)
	}
	return int(v), nil
}

func ConvertFloat32ToInt8(v float32) (int8, error) {
	if v > math.MaxInt8 || v < math.MinInt8 {
		return 0, NewOverflowError(v, math.MinInt8, math.MaxInt8)
	}
	return int8(v), nil
}

func ConvertFloat32ToInt16(v float32) (int16, error) {
	if v > math.MaxInt16 || v < math.MinInt16 {
		return 0, NewOverflowError(v, math.MinInt16, math.MaxInt16)
	}
	return int16(v), nil
}

func ConvertFloat32ToInt32(v float32) (int32, error) {
	if v > math.MaxInt32 || v < math.MinInt32 {
		return 0, NewOverflowError(v, math.MinInt32, math.MaxInt32)
	}
	return int32(v), nil
}

func ConvertFloat32ToInt64(v float32) (int64, error) {
	if v > math.MaxInt64 || v < math.MinInt64 {
		return 0, NewOverflowError(v, math.MinInt64, math.MaxInt64)
	}
	return int64(v), nil
}

func ConvertFloat32ToUint(v float32) (uint, error) {
	if v > math.MaxUint32 || v < 0 {
		return 0, NewOverflowError(v, 0, uint32(math.MaxUint32))
	}
	return uint(v), nil
}

func ConvertFloat32ToUint8(v float32) (uint8, error) {
	if v > math.MaxUint8 || v < 0 {
		return 0, NewOverflowError(v, 0, uint8(math.MaxUint8))
	}
	return uint8(v), nil
}

func ConvertFloat32ToUint16(v float32) (uint16, error) {
	if v > math.MaxUint16 || v < 0 {
		return 0, NewOverflowError(v, 0, uint16(math.MaxUint16))
	}
	return uint16(v), nil
}

func ConvertFloat32ToUint32(v float32) (uint32, error) {
	if v > math.MaxUint32 || v < 0 {
		return 0, NewOverflowError(v, 0, uint32(math.MaxUint32))
	}
	return uint32(v), nil
}

func ConvertFloat32ToUint64(v float32) (uint64, error) {
	if v > math.MaxUint64 || v < 0 {
		return 0, NewOverflowError(v, 0, uint64(math.MaxUint64))
	}
	return uint64(v), nil
}

func ConvertFloat32ToFloat32(v float32) (float32, error) {
	return v, nil
}

func ConvertFloat32ToFloat64(v float32) (float64, error) {
	return float64(v), nil
}

func ConvertFloat32ToString(v float32) (string, error) {
	return strconv.FormatFloat(float64(v), 'f', -1, 32), nil
}

// ----------------------------------------------------------------------------
// Float64
// ----------------------------------------------------------------------------

func ConvertFloat64ToInt(v float64) (int, error) {
	if v > math.MaxInt32 || v < math.MinInt32 {
		return 0, NewOverflowError(v, math.MinInt32, math.MaxInt32)
	}
	return int(v), nil
}

func ConvertFloat64ToInt8(v float64) (int8, error) {
	if v > math.MaxInt8 || v < math.MinInt8 {
		return 0, NewOverflowError(v, math.MinInt8, math.MaxInt8)
	}
	return int8(v), nil
}

func ConvertFloat64ToInt16(v float64) (int16, error) {
	if v > math.MaxInt16 || v < math.MinInt16 {
		return 0, NewOverflowError(v, math.MinInt16, math.MaxInt16)
	}
	return int16(v), nil
}

func ConvertFloat64ToInt32(v float64) (int32, error) {
	if v > math.MaxInt32 || v < math.MinInt32 {
		return 0, NewOverflowError(v, math.MinInt32, math.MaxInt32)
	}
	return int32(v), nil
}

func ConvertFloat64ToInt64(v float64) (int64, error) {
	if v > math.MaxInt64 || v < math.MinInt64 {
		return 0, NewOverflowError(v, math.MinInt64, math.MaxInt64)
	}
	return int64(v), nil
}

func ConvertFloat64ToUint(v float64) (uint, error) {
	if v > math.MaxUint32 || v < 0 {
		return 0, NewOverflowError(v, 0, uint32(math.MaxUint32))
	}
	return uint(v), nil
}

func ConvertFloat64ToUint8(v float64) (uint8, error) {
	if v > math.MaxUint8 || v < 0 {
		return 0, NewOverflowError(v, 0, uint8(math.MaxUint8))
	}
	return uint8(v), nil
}

func ConvertFloat64ToUint16(v float64) (uint16, error) {
	if v > math.MaxUint16 || v < 0 {
		return 0, NewOverflowError(v, 0, uint16(math.MaxUint16))
	}
	return uint16(v), nil
}

func ConvertFloat64ToUint32(v float64) (uint32, error) {
	if v > math.MaxUint32 || v < 0 {
		return 0, NewOverflowError(v, 0, uint32(math.MaxUint32))
	}
	return uint32(v), nil
}

func ConvertFloat64ToUint64(v float64) (uint64, error) {
	if v > math.MaxUint64 || v < 0 {
		return 0, NewOverflowError(v, 0, uint64(math.MaxUint64))
	}
	return uint64(v), nil
}

func ConvertFloat64ToFloat32(v float64) (float32, error) {
	return float32(v), nil
}

func ConvertFloat64ToFloat64(v float64) (float64, error) {
	return v, nil
}

func ConvertFloat64ToString(v float64) (string, error) {
	return strconv.FormatFloat(v, 'f', -1, 64), nil
}

// ----------------------------------------------------------------------------
// String
// ----------------------------------------------------------------------------

func ConvertStringToInt(v string) (int, error) {
	return strconv.Atoi(v)
}

func ConvertStringToInt8(v string) (int8, error) {
	i, err := strconv.ParseInt(v, 10, 8)
	return int8(i), err
}

func ConvertStringToInt16(v string) (int16, error) {
	i, err := strconv.ParseInt(v, 10, 16)
	return int16(i), err
}

func ConvertStringToInt32(v string) (int32, error) {
	i, err := strconv.ParseInt(v, 10, 32)
	return int32(i), err
}

func ConvertStringToInt64(v string) (int64, error) {
	return strconv.ParseInt(v, 10, 64)
}

func ConvertStringToUint(v string) (uint, error) {
	i, err := strconv.ParseUint(v, 10, 32)
	return uint(i), err
}

func ConvertStringToUint8(v string) (uint8, error) {
	i, err := strconv.ParseUint(v, 10, 8)
	return uint8(i), err
}

func ConvertStringToUint16(v string) (uint16, error) {
	i, err := strconv.ParseUint(v, 10, 16)
	return uint16(i), err
}

func ConvertStringToUint32(v string) (uint32, error) {
	i, err := strconv.ParseUint(v, 10, 32)
	return uint32(i), err
}

func ConvertStringToUint64(v string) (uint64, error) {
	return strconv.ParseUint(v, 10, 64)
}

func ConvertStringToFloat32(v string) (float32, error) {
	f, err := strconv.ParseFloat(v, 32)
	return float32(f), err
}

func ConvertStringToFloat64(v string) (float64, error) {
	return strconv.ParseFloat(v, 64)
}

func ConvertStringToString(v string) (string, error) {
	return v, nil
}

var (
	IntToIntConverter         = NewTypeConverter(ConvertIntToInt)
	IntToInt8Converter        = NewTypeConverter(ConvertIntToInt8)
	IntToInt16Converter       = NewTypeConverter(ConvertIntToInt16)
	IntToInt32Converter       = NewTypeConverter(ConvertIntToInt32)
	IntToInt64Converter       = NewTypeConverter(ConvertIntToInt64)
	IntToUintConverter        = NewTypeConverter(ConvertIntToUint)
	IntToUint8Converter       = NewTypeConverter(ConvertIntToUint8)
	IntToUint16Converter      = NewTypeConverter(ConvertIntToUint16)
	IntToUint32Converter      = NewTypeConverter(ConvertIntToUint32)
	IntToUint64Converter      = NewTypeConverter(ConvertIntToUint64)
	IntToFloat32Converter     = NewTypeConverter(ConvertIntToFloat32)
	IntToFloat64Converter     = NewTypeConverter(ConvertIntToFloat64)
	IntToStringConverter      = NewTypeConverter(ConvertIntToString)
	Int8ToIntConverter        = NewTypeConverter(ConvertInt8ToInt)
	Int8ToInt8Converter       = NewTypeConverter(ConvertInt8ToInt8)
	Int8ToInt16Converter      = NewTypeConverter(ConvertInt8ToInt16)
	Int8ToInt32Converter      = NewTypeConverter(ConvertInt8ToInt32)
	Int8ToInt64Converter      = NewTypeConverter(ConvertInt8ToInt64)
	Int8ToUintConverter       = NewTypeConverter(ConvertInt8ToUint)
	Int8ToUint8Converter      = NewTypeConverter(ConvertInt8ToUint8)
	Int8ToUint16Converter     = NewTypeConverter(ConvertInt8ToUint16)
	Int8ToUint32Converter     = NewTypeConverter(ConvertInt8ToUint32)
	Int8ToUint64Converter     = NewTypeConverter(ConvertInt8ToUint64)
	Int8ToFloat32Converter    = NewTypeConverter(ConvertInt8ToFloat32)
	Int8ToFloat64Converter    = NewTypeConverter(ConvertInt8ToFloat64)
	Int8ToStringConverter     = NewTypeConverter(ConvertInt8ToString)
	Int16ToIntConverter       = NewTypeConverter(ConvertInt16ToInt)
	Int16ToInt8Converter      = NewTypeConverter(ConvertInt16ToInt8)
	Int16ToInt16Converter     = NewTypeConverter(ConvertInt16ToInt16)
	Int16ToInt32Converter     = NewTypeConverter(ConvertInt16ToInt32)
	Int16ToInt64Converter     = NewTypeConverter(ConvertInt16ToInt64)
	Int16ToUintConverter      = NewTypeConverter(ConvertInt16ToUint)
	Int16ToUint8Converter     = NewTypeConverter(ConvertInt16ToUint8)
	Int16ToUint16Converter    = NewTypeConverter(ConvertInt16ToUint16)
	Int16ToUint32Converter    = NewTypeConverter(ConvertInt16ToUint32)
	Int16ToUint64Converter    = NewTypeConverter(ConvertInt16ToUint64)
	Int16ToFloat32Converter   = NewTypeConverter(ConvertInt16ToFloat32)
	Int16ToFloat64Converter   = NewTypeConverter(ConvertInt16ToFloat64)
	Int16ToStringConverter    = NewTypeConverter(ConvertInt16ToString)
	Int32ToIntConverter       = NewTypeConverter(ConvertInt32ToInt)
	Int32ToInt8Converter      = NewTypeConverter(ConvertInt32ToInt8)
	Int32ToInt16Converter     = NewTypeConverter(ConvertInt32ToInt16)
	Int32ToInt32Converter     = NewTypeConverter(ConvertInt32ToInt32)
	Int32ToInt64Converter     = NewTypeConverter(ConvertInt32ToInt64)
	Int32ToUintConverter      = NewTypeConverter(ConvertInt32ToUint)
	Int32ToUint8Converter     = NewTypeConverter(ConvertInt32ToUint8)
	Int32ToUint16Converter    = NewTypeConverter(ConvertInt32ToUint16)
	Int32ToUint32Converter    = NewTypeConverter(ConvertInt32ToUint32)
	Int32ToUint64Converter    = NewTypeConverter(ConvertInt32ToUint64)
	Int32ToFloat32Converter   = NewTypeConverter(ConvertInt32ToFloat32)
	Int32ToFloat64Converter   = NewTypeConverter(ConvertInt32ToFloat64)
	Int32ToStringConverter    = NewTypeConverter(ConvertInt32ToString)
	Int64ToIntConverter       = NewTypeConverter(ConvertInt64ToInt)
	Int64ToInt8Converter      = NewTypeConverter(ConvertInt64ToInt8)
	Int64ToInt16Converter     = NewTypeConverter(ConvertInt64ToInt16)
	Int64ToInt32Converter     = NewTypeConverter(ConvertInt64ToInt32)
	Int64ToInt64Converter     = NewTypeConverter(ConvertInt64ToInt64)
	Int64ToUintConverter      = NewTypeConverter(ConvertInt64ToUint)
	Int64ToUint8Converter     = NewTypeConverter(ConvertInt64ToUint8)
	Int64ToUint16Converter    = NewTypeConverter(ConvertInt64ToUint16)
	Int64ToUint32Converter    = NewTypeConverter(ConvertInt64ToUint32)
	Int64ToUint64Converter    = NewTypeConverter(ConvertInt64ToUint64)
	Int64ToFloat32Converter   = NewTypeConverter(ConvertInt64ToFloat32)
	Int64ToFloat64Converter   = NewTypeConverter(ConvertInt64ToFloat64)
	Int64ToStringConverter    = NewTypeConverter(ConvertInt64ToString)
	UintToIntConverter        = NewTypeConverter(ConvertUintToInt)
	UintToInt8Converter       = NewTypeConverter(ConvertUintToInt8)
	UintToInt16Converter      = NewTypeConverter(ConvertUintToInt16)
	UintToInt32Converter      = NewTypeConverter(ConvertUintToInt32)
	UintToInt64Converter      = NewTypeConverter(ConvertUintToInt64)
	UintToUintConverter       = NewTypeConverter(ConvertUintToUint)
	UintToUint8Converter      = NewTypeConverter(ConvertUintToUint8)
	UintToUint16Converter     = NewTypeConverter(ConvertUintToUint16)
	UintToUint32Converter     = NewTypeConverter(ConvertUintToUint32)
	UintToUint64Converter     = NewTypeConverter(ConvertUintToUint64)
	UintToFloat32Converter    = NewTypeConverter(ConvertUintToFloat32)
	UintToFloat64Converter    = NewTypeConverter(ConvertUintToFloat64)
	UintToStringConverter     = NewTypeConverter(ConvertUintToString)
	Uint8ToIntConverter       = NewTypeConverter(ConvertUint8ToInt)
	Uint8ToInt8Converter      = NewTypeConverter(ConvertUint8ToInt8)
	Uint8ToInt16Converter     = NewTypeConverter(ConvertUint8ToInt16)
	Uint8ToInt32Converter     = NewTypeConverter(ConvertUint8ToInt32)
	Uint8ToInt64Converter     = NewTypeConverter(ConvertUint8ToInt64)
	Uint8ToUintConverter      = NewTypeConverter(ConvertUint8ToUint)
	Uint8ToUint8Converter     = NewTypeConverter(ConvertUint8ToUint8)
	Uint8ToUint16Converter    = NewTypeConverter(ConvertUint8ToUint16)
	Uint8ToUint32Converter    = NewTypeConverter(ConvertUint8ToUint32)
	Uint8ToUint64Converter    = NewTypeConverter(ConvertUint8ToUint64)
	Uint8ToFloat32Converter   = NewTypeConverter(ConvertUint8ToFloat32)
	Uint8ToFloat64Converter   = NewTypeConverter(ConvertUint8ToFloat64)
	Uint8ToStringConverter    = NewTypeConverter(ConvertUint8ToString)
	Uint16ToIntConverter      = NewTypeConverter(ConvertUint16ToInt)
	Uint16ToInt8Converter     = NewTypeConverter(ConvertUint16ToInt8)
	Uint16ToInt16Converter    = NewTypeConverter(ConvertUint16ToInt16)
	Uint16ToInt32Converter    = NewTypeConverter(ConvertUint16ToInt32)
	Uint16ToInt64Converter    = NewTypeConverter(ConvertUint16ToInt64)
	Uint16ToUintConverter     = NewTypeConverter(ConvertUint16ToUint)
	Uint16ToUint8Converter    = NewTypeConverter(ConvertUint16ToUint8)
	Uint16ToUint16Converter   = NewTypeConverter(ConvertUint16ToUint16)
	Uint16ToUint32Converter   = NewTypeConverter(ConvertUint16ToUint32)
	Uint16ToUint64Converter   = NewTypeConverter(ConvertUint16ToUint64)
	Uint16ToFloat32Converter  = NewTypeConverter(ConvertUint16ToFloat32)
	Uint16ToFloat64Converter  = NewTypeConverter(ConvertUint16ToFloat64)
	Uint16ToStringConverter   = NewTypeConverter(ConvertUint16ToString)
	Uint32ToIntConverter      = NewTypeConverter(ConvertUint32ToInt)
	Uint32ToInt8Converter     = NewTypeConverter(ConvertUint32ToInt8)
	Uint32ToInt16Converter    = NewTypeConverter(ConvertUint32ToInt16)
	Uint32ToInt32Converter    = NewTypeConverter(ConvertUint32ToInt32)
	Uint32ToInt64Converter    = NewTypeConverter(ConvertUint32ToInt64)
	Uint32ToUintConverter     = NewTypeConverter(ConvertUint32ToUint)
	Uint32ToUint8Converter    = NewTypeConverter(ConvertUint32ToUint8)
	Uint32ToUint16Converter   = NewTypeConverter(ConvertUint32ToUint16)
	Uint32ToUint32Converter   = NewTypeConverter(ConvertUint32ToUint32)
	Uint32ToUint64Converter   = NewTypeConverter(ConvertUint32ToUint64)
	Uint32ToFloat32Converter  = NewTypeConverter(ConvertUint32ToFloat32)
	Uint32ToFloat64Converter  = NewTypeConverter(ConvertUint32ToFloat64)
	Uint32ToStringConverter   = NewTypeConverter(ConvertUint32ToString)
	Uint64ToIntConverter      = NewTypeConverter(ConvertUint64ToInt)
	Uint64ToInt8Converter     = NewTypeConverter(ConvertUint64ToInt8)
	Uint64ToInt16Converter    = NewTypeConverter(ConvertUint64ToInt16)
	Uint64ToInt32Converter    = NewTypeConverter(ConvertUint64ToInt32)
	Uint64ToInt64Converter    = NewTypeConverter(ConvertUint64ToInt64)
	Uint64ToUintConverter     = NewTypeConverter(ConvertUint64ToUint)
	Uint64ToUint8Converter    = NewTypeConverter(ConvertUint64ToUint8)
	Uint64ToUint16Converter   = NewTypeConverter(ConvertUint64ToUint16)
	Uint64ToUint32Converter   = NewTypeConverter(ConvertUint64ToUint32)
	Uint64ToUint64Converter   = NewTypeConverter(ConvertUint64ToUint64)
	Uint64ToFloat32Converter  = NewTypeConverter(ConvertUint64ToFloat32)
	Uint64ToFloat64Converter  = NewTypeConverter(ConvertUint64ToFloat64)
	Uint64ToStringConverter   = NewTypeConverter(ConvertUint64ToString)
	Float32ToIntConverter     = NewTypeConverter(ConvertFloat32ToInt)
	Float32ToInt8Converter    = NewTypeConverter(ConvertFloat32ToInt8)
	Float32ToInt16Converter   = NewTypeConverter(ConvertFloat32ToInt16)
	Float32ToInt32Converter   = NewTypeConverter(ConvertFloat32ToInt32)
	Float32ToInt64Converter   = NewTypeConverter(ConvertFloat32ToInt64)
	Float32ToUintConverter    = NewTypeConverter(ConvertFloat32ToUint)
	Float32ToUint8Converter   = NewTypeConverter(ConvertFloat32ToUint8)
	Float32ToUint16Converter  = NewTypeConverter(ConvertFloat32ToUint16)
	Float32ToUint32Converter  = NewTypeConverter(ConvertFloat32ToUint32)
	Float32ToUint64Converter  = NewTypeConverter(ConvertFloat32ToUint64)
	Float32ToFloat32Converter = NewTypeConverter(ConvertFloat32ToFloat32)
	Float32ToFloat64Converter = NewTypeConverter(ConvertFloat32ToFloat64)
	Float32ToStringConverter  = NewTypeConverter(ConvertFloat32ToString)
	Float64ToIntConverter     = NewTypeConverter(ConvertFloat64ToInt)
	Float64ToInt8Converter    = NewTypeConverter(ConvertFloat64ToInt8)
	Float64ToInt16Converter   = NewTypeConverter(ConvertFloat64ToInt16)
	Float64ToInt32Converter   = NewTypeConverter(ConvertFloat64ToInt32)
	Float64ToInt64Converter   = NewTypeConverter(ConvertFloat64ToInt64)
	Float64ToUintConverter    = NewTypeConverter(ConvertFloat64ToUint)
	Float64ToUint8Converter   = NewTypeConverter(ConvertFloat64ToUint8)
	Float64ToUint16Converter  = NewTypeConverter(ConvertFloat64ToUint16)
	Float64ToUint32Converter  = NewTypeConverter(ConvertFloat64ToUint32)
	Float64ToUint64Converter  = NewTypeConverter(ConvertFloat64ToUint64)
	Float64ToFloat32Converter = NewTypeConverter(ConvertFloat64ToFloat32)
	Float64ToFloat64Converter = NewTypeConverter(ConvertFloat64ToFloat64)
	Float64ToStringConverter  = NewTypeConverter(ConvertFloat64ToString)
	StringToIntConverter      = NewTypeConverter(ConvertStringToInt)
	StringToInt8Converter     = NewTypeConverter(ConvertStringToInt8)
	StringToInt16Converter    = NewTypeConverter(ConvertStringToInt16)
	StringToInt32Converter    = NewTypeConverter(ConvertStringToInt32)
	StringToInt64Converter    = NewTypeConverter(ConvertStringToInt64)
	StringToUintConverter     = NewTypeConverter(ConvertStringToUint)
	StringToUint8Converter    = NewTypeConverter(ConvertStringToUint8)
	StringToUint16Converter   = NewTypeConverter(ConvertStringToUint16)
	StringToUint32Converter   = NewTypeConverter(ConvertStringToUint32)
	StringToUint64Converter   = NewTypeConverter(ConvertStringToUint64)
	StringToFloat32Converter  = NewTypeConverter(ConvertStringToFloat32)
	StringToFloat64Converter  = NewTypeConverter(ConvertStringToFloat64)
	StringToStringConverter   = NewTypeConverter(ConvertStringToString)
	BoolToBoolConverter       = NewTypeConverter(ConvertBoolToBool)
	BoolToIntConverter        = NewTypeConverter(ConvertBoolToNumber[int])
	BoolToInt8Converter       = NewTypeConverter(ConvertBoolToNumber[int8])
	BoolToInt16Converter      = NewTypeConverter(ConvertBoolToNumber[int16])
	BoolToInt32Converter      = NewTypeConverter(ConvertBoolToNumber[int32])
	BoolToInt64Converter      = NewTypeConverter(ConvertBoolToNumber[int64])
	BoolToUintConverter       = NewTypeConverter(ConvertBoolToNumber[uint])
	BoolToUint8Converter      = NewTypeConverter(ConvertBoolToNumber[uint8])
	BoolToUint16Converter     = NewTypeConverter(ConvertBoolToNumber[uint16])
	BoolToUint32Converter     = NewTypeConverter(ConvertBoolToNumber[uint32])
	BoolToUint64Converter     = NewTypeConverter(ConvertBoolToNumber[uint64])
	BoolToFloat32Converter    = NewTypeConverter(ConvertBoolToNumber[float32])
	BoolToFloat64Converter    = NewTypeConverter(ConvertBoolToNumber[float64])
	BoolToStringConverter     = NewTypeConverter(ConvertBoolToString)
	IntToBoolConverter        = NewTypeConverter(ConvertNumberToBool[int])
	Int8ToBoolConverter       = NewTypeConverter(ConvertNumberToBool[int8])
	Int16ToBoolConverter      = NewTypeConverter(ConvertNumberToBool[int16])
	Int32ToBoolConverter      = NewTypeConverter(ConvertNumberToBool[int32])
	Int64ToBoolConverter      = NewTypeConverter(ConvertNumberToBool[int64])
	UintToBoolConverter       = NewTypeConverter(ConvertNumberToBool[uint])
	Uint8ToBoolConverter      = NewTypeConverter(ConvertNumberToBool[uint8])
	Uint16ToBoolConverter     = NewTypeConverter(ConvertNumberToBool[uint16])
	Uint32ToBoolConverter     = NewTypeConverter(ConvertNumberToBool[uint32])
	Uint64ToBoolConverter     = NewTypeConverter(ConvertNumberToBool[uint64])
	Float32ToBoolConverter    = NewTypeConverter(ConvertNumberToBool[float32])
	Float64ToBoolConverter    = NewTypeConverter(ConvertNumberToBool[float64])
	StringToBoolConverter     = NewTypeConverter(ConvertStringToBool)
)
