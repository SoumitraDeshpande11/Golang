package main

import "fmt"

func main() {
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

    fmt.Print("Counting: ")
    for i := 1; i <= 3; i++ {
        fmt.Print(i, " ")
    }
    fmt.Println()

    
    n := 5
    for n > 0 {
        fmt.Print(n, " ")
        n--
    }
    fmt.Println()

    switch time := 11; {
    case time < 12:
        fmt.Println("Good morning!")
    case time < 17:
        fmt.Println("Good afternoon!")
    default:
        fmt.Println("Good evening!")
    }
}
