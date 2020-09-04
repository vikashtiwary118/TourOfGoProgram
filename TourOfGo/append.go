package main 


import "fmt"

func main() {
	var s[]int

	printSlice(s)
    //append works on nil series
	s=append(s,0)
	printSlice(s)
    //slice grows as needed
	s=append(s,1)
	printSlice(s)

	//we can add more than one lements in array together

	s=append(s,2,3,4)
	printSlice(s)


}

func printSlice(s []int) {

	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)
	
}