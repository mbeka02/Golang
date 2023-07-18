package main

import (
	"fmt"

	"example.com/greetings"
)


func main(){
	fmt.Println("Hey there")
	message :=greetings.Hello("John")
	fmt.Println(message)
}