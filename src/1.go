
package main

import "fmt"

func main() {
	num1 := 4
	num2 := 7
	result := calculateSum(num1, num2)
	fmt.Println("The sum of", num1, "and", num2, "is", result)
}

func calculateSum(a int, b int) int {
	return a + b
}