package main 


import "fmt"

func main() {
	q :=[]int{2,3,4,5,6,7}

	fmt.Println(q)


	r:=[]bool{true,false,true,true,false,true}

	fmt.Println(r)

	s:=[]struct{
		i int
		b bool
	}{
		{2,true},
		{3,false},
		{4,true},
		{5,true},
		{6,false},
		{7,true},
	}
	fmt.Println(s)
}