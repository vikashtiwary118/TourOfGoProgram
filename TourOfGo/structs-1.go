package main 

import "fmt"

type user struct{
	username string

	age int
}

func main() {

	jack:=user{
		username:"jack",
		age:18,
	}

	fmt.Println(jack)
}