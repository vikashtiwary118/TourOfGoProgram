package main 

import "fmt"

func main() {

	defer func () {
		if r:=recover();r!=nil{
			fmt.Println("Ran Into A Error")
			main()
		}
	}()

	functions:=map[string]func(int,int) int{
		"add":add,
		"subtract":subtract,
		"multiply":multiply,
		"divide":divide,

	}

	var currentNumber int
	fmt.Println("Enter the Start Number")
	fmt.Scan(&currentNumber)

	for true{
		var funtionName string
		var number int
		fmt.Println("What function do you want to use:")
		fmt.Scan(&funtionName)
		if funtionName=="done"{
			break
		}

		fmt.Println("Enter the number")
		fmt.Scan(&number)

		currentNumber=functions[funtionName](currentNumber,number)


	}

	fmt.Println("Your Number Is ")

	fmt.Println(currentNumber)
	
}

func add(x,y int) int{

	return x+y
	
}

func subtract(x,y int) int{

	return x-y
	
}

func multiply(x,y int) int{

	return x*y
	
}

func divide(x,y int) int{

	return x/y
	
}