package main 
func main(){
	// var a,s = 100 , "abc"
	// println(a, s)

	// x := 100
	// println(x )

	// x, y := 1, 2
	// x, y = y+3, x+2
	// println(x, y)

	// const(
	// 	x = 100
	// 	y float32 = x 
	// )
	// println(x, y)

	// const(
	// 	x uint16 = 120
	// 	y
	// )
	// println(x, y)
	// testEnum()
	// testIota2()
	// testIotaBreak()
	// testIotaType()
	testCustomTypeEnum()
}


func testEnum(){
	const(
		x = iota 
		y
		z 
	)
	const(
		_ = iota 
		KB = 1<< (10 * iota)
		MB
		GB
	)
	println(y,z, KB, GB)
}

func testIota2(){
	const(
		_, _ = iota, iota * 10
		a, b 
		c, d
	)
	println(a,b,c, d)
}

func testIotaBreak(){
	const(
		a = iota 
		b  
		c = 100
		d  //100 与上一行常量右值表达式相同
		e = iota  //4， 恢复 iota 自增，包括 c,d 两行。。
		f
	)
	println(d, f)
}

func testIotaType(){
	const(
		a = iota 
		b float32 = iota 
		c  //和上面的一致类型
		d = iota //没有指定类型，默认是 int
	)
	println(a,b,c, d)
}


func testCustomTypeEnum(){
	type color byte 
	const(
		black color = iota 
		red 
		blue
	)
	var test = func (c color){
		println(c)
	}
	test(red)
	const m  =100
	test(m) // 常量和字面量不能超出 color/byte 取值范围即可
	// test(256)
}