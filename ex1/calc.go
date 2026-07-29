package main

import (
	"fmt"
	"strconv"
)

func calcOper(val float64, result *float64, oper string) {
	switch oper {
	case "+":
		*result = *result + val
	case "/":
		*result = *result / val
	case "*":
		*result = *result * val
	case "-":
		*result = *result - val
	}
}

func validOper(oper string) bool {
	switch oper {
	case "+", "/", "*", "-":
		return true
	default:
		return false
	}
}

func main() {
	var stringVal string
	var result float64
	oper := "+"
	for i := 1; i < 4; i++ {
		if i == 1 || i == 3 {
			fmt.Printf("input value %d: ", i)
			fmt.Scan(&stringVal)
			val, err := strconv.ParseFloat(stringVal, 64)
			if err != nil {
				fmt.Println("Invalid input")
				i--
				continue
			}
			calcOper(val, &result, oper)
		} else {
			fmt.Printf("Operand: ")
			fmt.Scan(&oper)
			if !validOper(oper) {
				fmt.Println("Invalid input")
				i--
				continue
			}
		}
	}
	fmt.Printf("%.3f", result)
}
