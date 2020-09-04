package main 

import ("fmt"
          "math")


var pow=[]int{1,2,4,8,16,32,64,128}


func main() {
	fmt.Println("2^3=",math.Pow(2,3))
	for i,v:=range pow{
		fmt.Printf("2**%d=%d\n",i,v)
	}
}

