package main

import "fmt"

func main() {
	fmt.Println(sayHi("Corey"))
}

func sayHi(person string) string {
	return fmt.Sprintf("Hi %s", person)
}
