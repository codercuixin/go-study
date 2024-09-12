package main 

func main(){
	// _ = func(s string) string{
	// 	return s 
	// }("abc")

	// add := func(x, y int) int{
	// 	return x + y 
	// }
	// _ = add(1, 2)

	// wrap(func(s string)string{
	// 	return "hello "+s 
	// })()

	// calc := struct{
	// 	add func(int, int) int 
	// 	mul func(int, int) int 
	// }{
	// 	add: func(x, y int) int{ return x+y},
	// 	mul: func(x, y int )int{return x *y},
	// }
	// _ = calc.add(1, 2)

	// passChannel()
}

func wrap(f func(string)string) func(){
	return func() {
		println(f("abc"))
	}
}

func passChannel(){
	funcChannel := make(chan func())
	go func(){
		defer close(funcChannel)
		f := <- funcChannel 
		f()
	}()
	funcChannel <- func(){
		println("hello world")
	}
	<- funcChannel
}

