package main

import (
	"amber/libs/crypt"
	"fmt"
)

func Hello(name string) string {
	result := "Hello " + name
	return result
}

func main() {
	fmt.Println(Hello("as"))
	fmt.Println(crypt.Crypt("as"))
}
