package handlers

import (
	"fmt"
	"os"
)

func RunCode() {
Menu:
	fmt.Println("******* Welcome To WorkSpace *******")
	fmt.Println("*** 1. 体脂计算器(BMI) ***")
	fmt.Println("*** 2. 计算直线斜率 ***")
	fmt.Println("*** exit ***")
	code := ""
	fmt.Scanf("Please Enter Code: %s\n\n", &code)

	switch code {
	case "1":
		BodyFatCalculator()
	case "2":
		SlopesEqual()
	case "exit":
		fmt.Println("Welcome to next time.")
		os.Exit(0)
	default:
		defer fmt.Printf("Please enter the correct code. \n\n")
		goto Menu
	}
}
