package main 

import ("log"
        "os"
        "fmt")



func main() {
	f, err := os.Open("xyz.ext")
	if err != nil {
	    log.Fatal(err)   //when error occurs it call log
	                     //Fatal is use to print error message and stop  it
	}
	
	fmt.Println("hahahah",f)
	
// do something with the open *File f
}