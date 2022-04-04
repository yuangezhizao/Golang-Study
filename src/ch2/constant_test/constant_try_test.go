package constant_test

import (
	"testing"
)

// 常量定义，快速设置连续值
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	//Open = 1 << iota
	//Close
	//Pending
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTryBit(t *testing.T) {
	//a := 7 // 0111
	a := 1 // 0001
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
