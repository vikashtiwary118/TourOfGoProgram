package main 

import "fmt"

func main() {
	c:=make(chan int,10)
	x,y:=0,1

	x,y=y,x+y

	c<-x

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
		
	}

	



}