package main 
import "reflect"
func main(){
	type X int 
	var  x X = 1 
	v := reflect.ValueOf(x)
	t := v.Type()

	println(t.PkgPath()+"."+t.Name())

	switch v.Kind(){
	case reflect.Int:
			println(v.Int())
	case reflect.String:
		println(v.String())
	}

	var y int = 2
	ty := reflect.TypeOf(y)
	println(ty)
	println(t == ty)
	println(t.Kind() == ty.Kind())

}