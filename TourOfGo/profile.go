package main 


import "fmt"

type Person struct{
	Name string

	Age int

	Hobbies []string

}


func main() {
	p1:={"sam",20,[]string{"cricket","football"}}

	fmt.Println("==P1==",p1)
}