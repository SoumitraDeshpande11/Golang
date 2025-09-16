package main

import "fmt"

func main() {
    // Arrays
    var a [5]int
    a[4] = 100
    fmt.Println("Array:", a)

    // Slices
    s := make([]string, 3)
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    s = append(s, "d", "e")
    fmt.Println("Slice:", s)
    fmt.Println("Slice[2:4]:", s[2:4])

    // Maps
    m := make(map[string]int)
    m["one"] = 1
    m["two"] = 2
    fmt.Println("Map:", m)
    
    delete(m, "two")
    _, exists := m["two"]
    fmt.Println("Key 'two' exists:", exists)
}
