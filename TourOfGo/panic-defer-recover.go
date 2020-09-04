package main 


import "fmt"


// panic is a way to force your program to stop

// after executing the panic statement we can stop forward execution

//defer executed whenever the funtion ends

//we can recover function from panic

func main() {

	fmt.Println("Hey")
	random()

	fmt.Println("This wont run")
	
}


func random() {

	//panic("We have paniced!!!!")

	defer func() {
		if r:=recover();r!=nil{
			fmt.Println("The function recoverd from panic with reason: %s\n",r)
		}
	}()

	// defer fmt.Println("0")
	// defer fmt.Println("3")
	fmt.Println("1")
    panic("some reason")
	fmt.Println("2")

	
}