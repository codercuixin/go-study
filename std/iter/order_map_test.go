package ordered_map_test

import (
	"fmt"
	ordered_map "go-study/std/iter"
	"testing"
)

type MyStruct struct {
	id   int
	name string
}

func (ms MyStruct) String() string {
	return fmt.Sprintf("ID: %d, Name: %s", ms.id, ms.name)
}

func TestOrderedMap(t *testing.T) {
	t.Run("print all elements in order.", func(t *testing.T) {
		om := ordered_map.New[string, MyStruct](10)
		om.Add("key1", MyStruct{id: 1, name: "test1"})
		om.Add("key3", MyStruct{id: 3, name: "test3"})
		om.Add("key2", MyStruct{id: 2, name: "test2"})

		for _, k := range om.Keys() {
			fmt.Printf("Value: %s\n", om.Get(k))
		}
	})
}







 func TestMyStructs(t *testing.T) {
	t.Run("print only name in struct", func(t *testing.T) {
		var myStructs MyStructs = MyStructs{
		 MyStruct{id: 1, name: "test1"},
		 MyStruct{id: 3, name: "test3"},
		 MyStruct{id: 2, name: "test2"},
		}
		for name := range myStructs.Names() {
		 fmt.Printf("Name: %s\n", name)
		}
	   })
}
