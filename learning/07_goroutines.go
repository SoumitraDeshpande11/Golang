package main

import (
    "fmt"
    "time"
)

func count(thing string) {
    for i := 1; i <= 3; i++ {
        fmt.Println(i, thing)
        time.Sleep(time.Millisecond * 500)
    }
}

func main() {
    go count("sheep")
    go count("fish")
    
    time.Sleep(2 * time.Second)
    fmt.Println("Done!")
}
