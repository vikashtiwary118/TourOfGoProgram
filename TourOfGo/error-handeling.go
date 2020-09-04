package main 

import ("fmt"
         "errors")

//there is no any cocept of try and catch like othr languages.Here Errors are treated as normal variables 

func main() {
	total,err:=sumOf(8,10)

	if err!=nil{
		fmt.Println("This was an error",err)
	}else{
		fmt.Println(total)
	}
	
}


func sumOf(start,end int) (int,error) {
	if start>end{
		return 0,errors.New("Start is greater than end")
	}
	total:=0
	for i := start; i <= end; i++ {
		total+=i
	}

	return total,nil
	
}