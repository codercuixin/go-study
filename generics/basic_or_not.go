package main 
type BasicInterface interface{
	M1()
}

type NonBasicInterface interface{
	BasicInterface
	~int |

~ string  // 包含类型元素
}

type MyString3 string 

func (MyString3) M1(){

}

func foo3[T NonBasicInterface](a T){
	// 非基本接口作为约束
}

func bar3[T BasicInterface](a T){
	// 基本接口作为约束
}

type MyStruct [T *int]struct{}

func main(){
	var s = MyString3("hello")
	var bi  BasicInterface = s // 基本接口类型支持常规用法
	// var nbi NonBasicInterface = s // 非基本类型不支持常规用法
	bi.M1()
	// nbi.M1()
	foo3(s)
	bar3(s)
}
