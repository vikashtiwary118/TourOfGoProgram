package main 

import "fmt"


func main() {
	names:=[4]string{
		"Ippo",
		"Ippi",
		"Avi",
		"Dampu",
	}

	fmt.Println(names)

	a:=names[0:2]
	b:=names[1:3]

	fmt.Println(a,b)

	b[0]="XXXX"

	fmt.Println(a,b)
	fmt.Println(names)
	
}