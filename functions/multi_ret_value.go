package main

import "errors"
import "log"

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x/ y, nil 
}

func main(){
	z, err := div(6,2)
	if err != nil{
		log.Fatalln(err)
	}
	println(z)
}

func log(x int, err error){
	fmt.Println(x, err)
}
