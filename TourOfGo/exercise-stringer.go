package main 

import "fmt"

type IPAddr [4]byte

//TODO: Add a "String()string" method to IP address


func main() {
	hosts:=map[string]IPAddr{
		"lookback":{17,0,0,1},
		"googleDNS":{8,8,8,8},
	}

	for name,ip:=range hosts{
		fmt.Printf("%v:%v\n",name,ip)
	}
}