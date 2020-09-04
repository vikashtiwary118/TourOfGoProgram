package main 


import ("fmt"
        "math")


func pow(x,n float64) float64{
	var v =math.Pow(x,n)

	return v

	
}

func main() {
	fmt.Println(pow(3,2),pow(3,3))
}