package main

import "fmt"

func main() {
    sum := add(5, 3)
    fmt.Println("5 + 3 =", sum)
    
    a, b := swap("hello", "world")
    fmt.Println(a, b)
    
    total := sumAll(1, 2, 3, 4, 5)
    fmt.Println("Sum:", total)
}

func add(x, y int) int {
    return x + y
}

func swap(x, y string) (string, string) {
    return y, x
}

func sumAll(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
