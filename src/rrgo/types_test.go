package rrgo

import "testing"

// TestTypes 类型。
//
// See also https://golang.org/ref/spec#Types
func TestTypes(t *testing.T) {
	const valBool bool = true

	const valInt int = 0
	const valUint uint = 0
	const valUintptr uintptr = 0

	const valInt8 int8 = 0
	const valInt16 int16 = 0
	const valInt32 int32 = 0
	const valRune rune = 0
	const valInt64 int64 = 0

	const valUint8 uint8 = 0
	const valByte byte = 0
	const valUint16 uint16 = 0
	const valUint32 uint32 = 0
	const valUint64 uint64 = 0

	const valFloat32 float32 = 0 // 小数点后 7 位
	const valFloat64 float64 = 0 // 小数点后 15 位

	const valComplex64 complex64 = 0 + 0i
	const valComplex128 complex128 = 0 + 0i

	const valString string = "zhengrr"

	val := 0
	ptr := &val
	if *ptr != 0 {
		t.Fail()
	}
}
