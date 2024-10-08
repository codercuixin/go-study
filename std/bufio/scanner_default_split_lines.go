package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil{
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}