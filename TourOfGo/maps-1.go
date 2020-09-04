package main 

import "fmt"

func main() {
	//Maps are similor to python dictionary(i.e hashtable etc)

	//specify key/value pair

	m:=make(map[string]int)

	m["a"]=0
	m["b"]=1
	fmt.Println(m)

	//Print the value of a specific key

	fmt.Println(m["a"])

	//len()  function is overloaded to work with map
	fmt.Println(len(m))

	//remove key/pair from map.(which can be done with delete keyword)
	delete(m , "a")
	fmt.Println(m)

	//another way to initialize a map (if you know already know the value/key)

	//ahead of time
	m2:=map[string]int{"key1":1,"key2":2}

	fmt.Println(m2)

	//the value and weather value is present.
	val,is_val_present:=m["b"]

	fmt.Println(val)
	fmt.Println(is_val_present)

	_, is_val_present2:=m["a"]
	fmt.Println(is_val_present2)




}