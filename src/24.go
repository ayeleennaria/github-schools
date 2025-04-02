package main

import "fmt"

func main() {
    var score int = 100
    if score < 50 {
        fmt.Println("low")
    } else if score >= 50 && score <= 69 {
        fmt.Println("average")
    } else if score >= 70 && score <= 89 {
        fmt.Println("high")
    } else if score > 90 {
        fmt.Println("great")
    } else {
        fmt.Println("unknown")
    }
}
