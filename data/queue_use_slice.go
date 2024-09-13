package  main 
import "fmt"
type Queue []int 
// use slice as stack

func NewQueue() *Queue{
	s := make(Queue, 0, 10)
	return &s 
}

func (s *Queue) Put(v int){
	*s = append(*s, v)
}

func (s *Queue) Get() (int,  bool){
	if len(*s) == 0{
		return 0, false 
	}

	s2:= *s
	v := s2[0]
	*s = s2[1:]
	return v, true
}


func main(){
	s := NewQueue()
	for i := 0; i<5; i++{
		s.Put(i+10)
	}
	for i :=0; i<7; i++{
		fmt.Println(s.Get())
	}
}