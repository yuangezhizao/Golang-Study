package condition_test

import (
	"runtime"
	"testing"
)

func TestIfMultiSec(t *testing.T) {
	//if v, err := someFuc(); err == nil {
	if a := 1 == 1; a {
		t.Log("1==1")
	} else {
		t.Log("2==2")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		t.Log("OS X.")
		//break // 默认相当于已添加 break 来退出一个 case
	case "linux":
		t.Log("Linux.")
	default:
		// freebad, openbsd, plan9, windows...
		t.Log(os)
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2: // 单个 case 中，可以出现多个结果选项，使用逗号分隔
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch { // 可以不设定 switch 之后的条件表达式，在此种情况下，整个 switch 结构与多个 if...else... 的逻辑作用等同
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknown")
		}
	}
}
