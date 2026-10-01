package main

import "fmt"

// func main() {
// 	res := Calculator(10, 5, "minus")
// 	fmt.Println(res)
// }

func main() {
	// minRes := Calculator(10, 5, "minus")
	// fmt.Println("Minus:", minRes)

	// plusRes := Calculator(10, 5, "plus")
	// fmt.Println("Plus:", plusRes)

	// divRes := Calculator(10, 5, "divide")
	// fmt.Println("Divide:", divRes)

	// mulRes := Calculator(10, 5, "multiply")
	// fmt.Println("Multiply:", mulRes)

	fmt.Println("Multiply:", Calculator(10, 5, "plus"))
    fmt.Println("Minus:", Calculator(10, 5, "minus"))
    fmt.Println("Multiply:", Calculator(10, 5, "multiply"))
    fmt.Println("Divide:", Calculator(10, 5, "divide"))

}	

func Calculator(a int, b int, operation string) int {
	if operation == "plus" {
		result := a + b
		return result

	}
	if operation == "minus" {
		result := a - b
		return result
	}

	if operation == "divide" {
		if b == 0 {
			fmt.Println("Ошибка: на 0 делить нельзя!")
			return 0
		}	 
		result := a / b
		return result
	}

	if operation == "multiply" {
		result := a*b
		return result
	}
	return -1
}