package main
import "fmt"


func main(){
	// bitop()
	// bitEnum()
	scope()
}

func scope(){
	 x := 1
	 {
		x := "abc"
		y := 2
		println(x, y)
	 }
	 println(x)
}

func bitop(){
	a := 0b01100101
	b := 0b11010100

	x := a ^ b 
	c := a &^ b 
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", c)
}

func bitEnum(){
	const(
		read byte = 1 << iota 
		write // 2^1
		exec // 2^2
		freeze // 2^3
	)
	a := read | write | freeze
	b := read | exec 
	c := a &^b 
	println(c == write | freeze)
	println(c & write ==write)
	println(c & read == read)
}