package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	//var a int = 1
	var a int32 = 1
	var b int64
	b = int64(a) // 不支持隐式类型转换
	var c MyInt
	c = MyInt(b) // 别名，也不支持隐式类型转换
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1 // 不支持指针运算
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
