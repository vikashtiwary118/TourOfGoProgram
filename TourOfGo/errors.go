package main 

import ("fmt"
        "time")


type MyError struct{
	When time.Time 
	What string 
}

func (e MyError) Error() string{
	return fmt.Sprint("at %v,%s",e.When,e.What)
	
}

func run()error {
	return &MyError{
		time.Now(),
		"it dint's work",
	}

	
}


func main() {
	if err :=run(); err!=nil{
		fmt.Println(err)
	}
}