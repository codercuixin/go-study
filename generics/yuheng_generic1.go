
package main 

import "cmp"
import "fmt"

func maxValue[T cmp.Ordered](x, y T) T{
	if x > y {
		return x 
	}
	return y
}

func main(){
	fmt.Println(maxValue(1, 2))
	fmt.Println(maxValue(1.1, 2.2))
}