package main

import (
	"errors"
	"fmt"
)

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed,the divider is zero.
	dividee:%d
	divider:0
	`
	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math:square root of negative number")
	}
	return float64(f * f), errors.New("")
}

func main() {
	fmt.Println(Sqrt(2))

	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10=", result)
	}
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}
