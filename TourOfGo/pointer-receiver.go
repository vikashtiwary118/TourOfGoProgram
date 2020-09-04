package main 


import "fmt"

//Pointer is basically address to a variable
func main() {

	x:=5
	setTo(&x,10)

	pointer:=&x


	fmt.Println(pointer)

	// *pointer=8

	fmt.Println(x)
	
}

func setTo(addr *int, newValue int){
	*addr=newValue

}