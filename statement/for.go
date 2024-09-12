package main 

func count()int{
	print("count.")
	return 3
}

func main(){
	//testFor()
	// testReuseLovalVar()
	// testNewLocalVar()
	// forRange()
	// forSingeValue()
	// rangeCopyData()
	rangeN()
}

func rangeN(){
	for i:= range 5 {
		println(i)
	}
}

func rangeCopyData(){
	data := [...]int{1, 2, 3}
	for i, v := range data{
		if i== 0{
			data[0] +=100
			data[1] +=200
			data[2] += 300
		}
		println(v, data[i])
	}
}
func forSingeValue(){
	data := []int{10, 11, 12}
	for index := range data{
		println(index )
	}
	for _, value := range data{
		println(value )
	}
	for range data{
		println("a")
	}
}

func forRange(){
	data := []int{10, 11, 12}
	for i, s := range data{
		println(i, s )
	}
}

func testReuseLovalVar(){
	for i:=0; i<3; i++{
		println(&i)
	}

}

func testNewLocalVar(){
	m := make(map[int]*int)
	for i:=0; i<3; i++{
		m[i] = &i;
		println(&i)
	}
}

func testFor(){
	for i, c:=0, count(); i <c ; i++{
		println("a", i)
	}
	c := 0
	for c <count(){
		println("b", c)
		c++
	}
}