package main 
import "runtime"

func test(x []byte){
	println(len(x))

	runtime.SetFinalizer(&x, func(*[]byte){
		println("drop")
	})
	runtime.GC()

	println("exit.")

}

func main(){
	test(make([]byte, 10<<20))
}