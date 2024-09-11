package main 
import "fmt"
type ordered2 interface{
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
	~string
}
type Stringer2 interface{
	comparable
	ordered2 
	String() string 
}

func StringifyWithoutZero[T Stringer2](s []T) (ret []string){
	var zero T 
	for _, v := range s {
		if v == zero{
			continue
		}
		ret = append(ret, v.String())
	}
	return ret 
}

func StringifyLessThan[T Stringer2](s []T, max T)(ret []string){
	var zero T 
	for _, v := range s{
		if v == zero || v >= max{
			continue
		}
		ret = append(ret, v.String())
	}
	return ret 
}


type MyString2 string 

func (s MyString2) String() string{
	return string(s)
}

func main(){
	sl := StringifyWithoutZero([]MyString2{"I", "Love", " ", "GoLang"})
	fmt.Println(sl)

	s2 := StringifyLessThan([]MyString2{"I", "", "love", "", "golang"}, MyString2("cpp")) // 输出：[I] fmt.Println(sl)
	fmt.Println(s2)
}