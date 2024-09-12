package main 
// import "os"
func main(){
	// testFILO()
	defer println("defer!")
	panic("panic!")
	// os.Exit(1)
}

func testFILO(){
	defer println(1)
	defer println(2)
	defer println(3)

	println("main")
}