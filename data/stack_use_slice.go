package  main 
import "fmt"
type Stack []int 
// use slice as stack

func NewStack() *Stack{
	s := make(Stack, 0, 10)
	return &s 
}

func (s *Stack) Push(v int){
	*s = append(*s, v)
}

func (s *Stack) Pop() (int,  bool){
	if len(*s) == 0{
		return 0, false 
	}

	s2, n := *s, len(*s)
	v := s2[n-1]
	*s = s2[:n-1]
	return v, true
}


func main(){
	s := NewStack()
	for i := 0; i<5; i++{
		s.Push(i+10)
	}
	for i :=0; i<7; i++{
		fmt.Println(s.Pop())
	}
}