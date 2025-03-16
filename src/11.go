package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Generating a random number between 1 and 10
	var n = rand.Intn(10) + 1
	fmt.Println("The random number is:", n)
}
