package rrgo

import "testing"

func TestTypes(t *testing.T) {
	const boolConst bool = true

	const intConst int = 0
	const uintConst uint = 0
	const uintptrConst uintptr = 0

	const int8Const int8 = 0
	const int16Const int16 = 0
	const int32Const int32 = 0
	const int64Const int64 = 0

	const runeConst rune = 0

	const uint8Const uint8 = 0
	const uint16Const uint16 = 0
	const uint32Const uint32 = 0
	const uint64Const uint64 = 0

	const byteConst byte = 0

	const float32Const float32 = 0  // 小数点后 7 位
	const float64Const float64 = 0  // 小数点后 15 位

	const complex64Const complex64 = 0 + 0i
	const complex128Const complex128 = 0 + 0i

	const stringConst string = "zhengrr"

	var intVar = 0
	var pointerVar *int = &intVar
	iv := *pointerVar

	iv += 0
}
