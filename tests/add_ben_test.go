package main 
import "testing"

//go test -v -bench . -run None . 
// run None 表示不运行单元测试
func BenchmarkAdd(b *testing.B){
	println("b.N = ", b.N)
	for i := 0; i<b.N; i++{
		addB(1, 2)
	}
}
func addB(x, y int)int{
	return x +y
}