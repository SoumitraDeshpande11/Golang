package main

import "fmt"

type Person struct {
    name string
    age  int
}

func (p Person) greet() string {
    return "Hello, my name is " + p.name
}

func (p *Person) birthday() {
    p.age++
}

func main() {
    p := Person{name: "Alice", age: 25}
    fmt.Println(p.greet())
    
    p.birthday()
    fmt.Printf("%s is now %d years old\n", p.name, p.age)
}
