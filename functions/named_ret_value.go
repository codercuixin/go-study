package main 

func add(x, y int)(z int){
	z = x + y 
	return
}

func addOverride(x, y int)(z int){
	{
		z := x+y
		//result parameter z not in scope at return (see details)compilerOutOfScopeResult
		// return
		return z 
	}
}