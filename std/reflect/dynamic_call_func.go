package main 
import "fmt"
import "reflect"

type X int 
func (x X) Add(y int) int{
	return int(x) + y 
}

func ( x X) Print(format string, y ...any){
	fmt.Printf(format, y...)
}

func main(){
	var x X  = 100
	v := reflect.ValueOf(&x)

	func(){
		m := v.MethodByName("Add")
		in := []reflect.Value{reflect.ValueOf(1)}
		out := m.Call(in)
		for _, v := range out{
			fmt.Println(v)
		}
	}()

	func(){
		m := v.MethodByName("Print")
		in := []reflect.Value{
			reflect.ValueOf("%d, %d, %d\n"),
			reflect.ValueOf([]any{1, 2,3}),
		}
		out := m.CallSlice(in)
		for _, v := range out{
			fmt.Println(v)
		}
	}()
}