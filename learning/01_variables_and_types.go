package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "Gopher"
	age := 3
	isLearning := true
	pi := 3.14159

	fmt.Println("Name:", name, "Type:", reflect.TypeOf(name))
	fmt.Println("Age:", age, "Type:", reflect.TypeOf(age))
	fmt.Println("Is Learning:", isLearning, "Type:", reflect.TypeOf(isLearning))
	fmt.Println("Pi:", pi, "Type:", reflect.TypeOf(pi))

	const version = "1.0"
	fmt.Println("Version:", version)
}
