package main 

type X int 
func main(){
	var a int = 100
	//cannot use a (variable of type int) as X value in variable 
	// var b X = a 
	b := (X)(a) 
	println(b)
}

func testUnmapped(){
	type A []int 
	var a [2]int
	var b [3]int = [3]int{1,2,3}
	//cannot use b (variable of type [3]int) as [2]int value in assignme
	a = b 


}