package main 

type Tester2 interface {
	test()
	string() string
}

type N2 struct {
	name string
}

func (n N2) test(){
	println("test")
}

func (n N2) string() string{
	return n.name
}

func testInterface(x Tester2){
	println(x.string())
}

func main(){
	var t Tester2
	t = N2{name: "test"}
	testInterface(t)
}