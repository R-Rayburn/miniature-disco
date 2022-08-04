package main

import "fmt"

func main() {
	greeting := "Hi there!"
	fmt.Println(greeting)
	fmt.Println([]byte(greeting))
	fruits := []string{"apple", "pear", "orange", "banana"}
	bowl := fruits[1:2]
	fmt.Println(bowl)
}
