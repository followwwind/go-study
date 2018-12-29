package main

import (
	"fmt"
)

// 定义一个 DivideError 结构
type DivideError struct {
	divideE int
	divideR int
}

// 实现     `error` 接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
  	divideE: %d
   	divideR: 0
	`
	return fmt.Sprintf(strFormat, de.divideE)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDivideE int, varDivideR int) (result int, errorMsg string) {
	if varDivideR == 0 {
		dData := DivideError{
			divideE: varDivideE,
			divideR: varDivideR,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDivideE / varDivideR, ""
	}

}

func main() {

	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

}
