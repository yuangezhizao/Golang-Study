package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	//defer func() {
	//	fmt.Println("Finally!")
	//}()

	defer func() {
		if err := recover(); err != nil {
			// 恢复错误
			fmt.Println("revovered from", err)
		}
	}()

	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	//os.Exit(-1)
	//fmt.Println("End")
}
