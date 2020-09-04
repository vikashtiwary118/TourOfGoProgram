package main 

import "fmt"

//create a struct 

type Myerror struct{}

//struct implements Error Method

func (myErr *Myerror) Error() string{
	return "Some thing Expected happend"
}


func main() {
	//create Error

	myErr:=&Myerror{}

	//print error message

	fmt.Println(myErr)


}