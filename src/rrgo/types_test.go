package rrgo

import "testing"

// TestTypes 类型。
//
// 常量：const identifier [type] = value
// 变量：var identifier [type] [= value]
//      identifier := value
//
// See also https://golang.org/ref/spec#Types
func TestTypes(t *testing.T) {
	var (
		valBool       bool
		valInt        int
		valUInt       uint
		valUIntPtr    uintptr
		valInt8       int8
		valInt16      int16
		valInt32      int32
		valRune       rune // int32
		valInt64      int64
		valUInt8      uint8
		valByte       byte // uint8
		valUInt16     uint16
		valUInt32     uint32
		valUInt64     uint64
		valFloat32    float32 // 小数点后 7 位
		valFloat64    float64 // 小数点后 15 位
		valComplex64  complex64
		valComplex128 complex128
		valString     string
		valPointer    *int
	)

	if valBool != false ||
		valInt != 0 ||
		valUInt != 0 ||
		valUIntPtr != 0 ||
		valInt8 != 0 ||
		valInt16 != 0 ||
		valInt32 != 0 ||
		valRune != '\u0000' ||
		valInt64 != 0 ||
		valUInt8 != 0 ||
		valByte != 0 ||
		valUInt16 != 0 ||
		valUInt32 != 0 ||
		valUInt64 != 0 ||
		valFloat32 != 0.0 ||
		valFloat64 != 0.0 ||
		valComplex64 != 0+0i ||
		valComplex128 != 0+0i ||
		valString != "" ||
		valPointer != nil {
		t.Fail()
	}
}
