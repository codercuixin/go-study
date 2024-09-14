
package main 

type N struct{}
func (N)test(){}
type node struct{
	value interface{
		test()
	}
}

func main(){
	var t interface{
		test()
	} = N{}

	n := node{value: t}
	n.value.test()
}
