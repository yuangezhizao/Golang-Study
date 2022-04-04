package operator_test

import "testing"

const (
	//Open = 1 << iota
	//Close
	//Pending
	Readable = 1 << iota
	Writeable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 3}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c) // 长度不同，无法编译通过
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	a := 7              // 0111
	a = a &^ Readable   // 按位清零操作符，清除可读权限
	a = a &^ Executable // 按位清零操作符，清除可执行权限
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
