package main 

import ("fmt"
        "database/sql"
        "github.com/go-sql-driver/mysql")



func main() {

	fmt.Println("Go Mysql Tutorial")


	db,err:=sql.Open("mysql")
	
}