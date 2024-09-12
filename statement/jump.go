package main 


func main(){
	// normal_usage()
	// nestJumpToTag()
	nestJumpToTag2()
}

func nestJumpToTag2(){
start:
	switch 1{
	case 1:
		println(1)
		for{
			break start 
		}
		println(2)
	}
	println(3)
}

func nestJumpToTag(){
outer:
	for x:=0; x <10; x++{
inner:
		for y:=0; y<10; y++{
			if x % 2 == 0 {
				continue outer 
			}
			if y > 3{
				println()
				break inner 
			}
			print(x, ":", y, " ")
		}
	}
}

func test(){
	goto test 
test: 	
}

// func wrong_usage(){
// 	for{
// 	loop:
// 	}
// 	// goto test // label test not declared
// 	// goto loop //goto loop jumps into blockcompilerJumpIntoBlock
// }

// func wrong_usage2(){
// 	 goto done //goto done jumps over variable declaration at line 23compilerJumpOverDecl

// 	 v := 0
// done: 
// 	println(v)
// }



func normal_usage(){
	for i:=0; i<3; i++{
		if i> 1{
			goto exit
		}
		println(i)
	}
exit:
	println("exit.")
}