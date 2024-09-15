package  main 
import "time"
var c int 

func inc() int{
	c ++ 
	return c 
}

func main(){
	a := 100
	go func (x, y int){
		time.Sleep(time.Second)
		println("go:", x, y)
	}(a, inc())

	a += 100
	println("main:", a, inc())
	time.Sleep(time.Second * 3)
}
