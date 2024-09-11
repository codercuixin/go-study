package main 
import "fmt"
func Stringify[T fmt.Stringer](s []T) (ret []string){
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret 
}

type MyString string 
func (s MyString) String() string{
	return string(s)
}

func main(){
	sl := Stringify[MyString]([]MyString{"I", "Love", "golang"})
	fmt.Println(sl)
}