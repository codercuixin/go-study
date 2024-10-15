package main

import "flag"

func main() {
	demo1()

}


func demo1(){
	// go run demo.go -flagname=234
	var flagvar int

	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
	flag.Parse()
	println(flagvar)
}