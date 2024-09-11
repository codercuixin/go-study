package main
import "fmt"

//cannot define new methods on non-local type
// func (int) test(){

// }

//invalid receiver type N (pointer or interface type)compilerInvalidRecv
// type N *int 
// func(N) test(){

// }

// type M int 
// func (m M) test(){
// 	println("test")
// }
// // method M.test already declared at go-study/methods/method_overview.go:16:12compilerDuplicateMethod

// func (m *M) test(){
// 	println("test")
// }

type N int 
func(n N) toString() string{
	return fmt.Sprintf("%v",n)
}
func (N) test(){
	println("test")
}

func( n N) copy(){
	n++
	fmt.Printf("copy %p, %v\n",&n,n)
}

func (n *N) ref(){
	*n++
	fmt.Printf("ref %p, %v\n",n,*n)
}

func main(){
	var n N = 100
	// println(n.toString())
	// n.test()
	n.copy()
	fmt.Printf("main %p, %v\n",&n,n)
	n.ref()

	var p *N = &n

	p.copy()
	fmt.Printf("main %p, %v\n",p,*p)
	p.ref()
	fmt.Printf("main %p, %v\n",p,*p)

}