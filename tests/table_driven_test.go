package main 
import "testing"

func addN(x, y int)int{
	return x + y
}

func TestAddN(t *testing.T){
	var tests = []struct{
		x int 
		y int 
		want int 
	}{
		{1, 1, 2},
		{2, 2, 6},
		{3,2, 5},
	}

	for _, tt := range tests{
		//规避闭包延迟
		o := tt 
		//并发子测试
		t.Run("", func(t *testing.T){
			t.Parallel()
			got := add(o.x, o.y)
			if got!= o.want{
				t.Errorf("add(%d, %d):want  %d, got %d",
				o.x, o.y, o.want, got);
			}
		})
	}
}