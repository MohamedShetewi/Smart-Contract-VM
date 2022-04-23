package main

import (
	"math/bits"
)

const size = 8

type DataWord [size]uint32

func newDataWord() *DataWord {
	return &DataWord{}
}

func (x *DataWord) setDataWord(byteArr []byte) {

	for i, j := 0, 0; i < size; i++ {
		for lShift := 0; lShift < 4; lShift, j = lShift+1, j+1 {
			(*x)[i] = (*x)[i] | (byteArr[j] << (lShift * 8))
		}
	}
}

func (x *DataWord) toInt() uint32 {
	return (*x)[0]
}

func (x *DataWord) setUint32(a uint32, i uint) {
	x[i] = a
}

func (x *DataWord) Add(y *DataWord) (result *DataWord) {
	var carry uint32 = 0
	for i := 0; i < len(result); i++ {
		result[0], carry = bits.Add32(x[0], y[0], carry)
	}
	return
}

func (x *DataWord) Sub(y *DataWord) (result *DataWord) {
	var borrow uint32 = 0
	for i := 0; i < len(result); i++ {
		result[i], borrow = bits.Sub32(x[i], y[i], borrow)
	}
	return
}

func (x *DataWord) Multiply(y *DataWord) (*DataWord, bool) {
	result := mul(x, y)
	ans := (*DataWord)(result[:size])

	var isOverFlow = false
	for i := size; i < size*2; i++ {
		isOverFlow = isOverFlow || result[i] != 0
	}

	return ans, isOverFlow
}

func mul(x, y *DataWord) (result [size * 2]uint32) {

	for Yi := 0; Yi < len(y); Yi++ {
		var carry uint32 = 0
		Ri := Yi
		Xi := 0
		for ; Xi < len(x); Xi = Xi + 1 {
			var lastRes = result[Xi+Ri]

			carry, result[Xi+Ri] = multiplyHelper(lastRes, x[Xi], y[Yi], carry)
		}
		result[Ri+Xi] = carry
	}
	return
}

func multiplyHelper(z, x, y, carry uint32) (hi, lo uint32) {
	hi, lo = bits.Mul32(x, y)
	lo, carry = bits.Add32(lo, carry, 0)
	hi, _ = bits.Add32(hi, 0, carry)
	lo, carry = bits.Add32(lo, z, 0)
	hi, _ = bits.Add32(hi, 0, carry)
	return hi, lo
}

func (x *DataWord) Div(y *DataWord) (result *DataWord) {

	return
}

func (x *DataWord) Mod(y *DataWord) (result *DataWord) {

	return
}

func (x *DataWord) GT(y *DataWord) bool {
	_, borrow := bits.Sub32(x[0], y[0], 0)
	for i := 1; i < len(x); i++ {
		_, borrow = bits.Sub32(x[i], y[i], borrow)
	}
	return borrow == 0
}

func (x *DataWord) LT(y *DataWord) bool {
	return !x.GT(y)
}

func (x *DataWord) SLT(y *DataWord) bool {
	dataWordSign := x.sign()
	xSign := y.sign()

	if xSign > dataWordSign {
		return true
	} else if xSign < dataWordSign {
		return false
	} else {
		return x.LT(y)
	}
}

func (x *DataWord) SGT(y *DataWord) bool {
	dataWordSign := x.sign()
	xSign := y.sign()

	if xSign < dataWordSign {
		return true
	} else if xSign > dataWordSign {
		return false
	} else {
		return x.GT(y)
	}
}

/*
	Returns the sign of the dataWord
	if dataWord > 0 return 1
	if dataWord < 0 return -1
	if dataWord == 0 return 0
*/
func (x *DataWord) sign() int {
	if x.IsZero() {
		return 0
	}
	if x[len(x)-1]&1<<31 != 0 {
		return -1
	}
	return 1
}

func (x *DataWord) Eq(y *DataWord) bool {
	isEqual := true
	for i := 0; i < len(x); i++ {
		isEqual = isEqual && x[i] == y[i]
	}
	return isEqual
}

func (x *DataWord) IsZero() bool {
	for i := 0; i < len(x); i++ {
		if x[i] != 0 {
			return false
		}
	}
	return true
}

func (x *DataWord) And(y *DataWord) (result *DataWord) {

	for i := 0; i < len(x); i++ {
		result[i] = x[i] & y[i]
	}
	return
}

func (x *DataWord) Or(y *DataWord) (result *DataWord) {

	for i := 0; i < len(x); i++ {
		x[i] = x[i] | y[i]
	}
	return
}

func (x *DataWord) Not() (result *DataWord) {

	for i := 0; i < len(x); i++ {
		result[i] = ^x[i]
	}
	return
}

func (x *DataWord) Xor(y *DataWord) (result *DataWord) {

	for i := 0; i < len(x); i++ {
		result[i] = x[i] ^ y[i]
	}
	return
}
