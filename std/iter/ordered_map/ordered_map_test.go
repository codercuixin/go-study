package ordered_map

import (
	"fmt"
	"iter"
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
		om := New[string, MyStruct](10)
		om.Add("key1", MyStruct{id: 1, name: "test1"})
		om.Add("key3", MyStruct{id: 3, name: "test3"})
		om.Add("key2", MyStruct{id: 2, name: "test2"})

		for _, k := range om.Keys() {
			fmt.Printf("Value: %s\n", om.Get(k))
		}
	})
}

func TestOrderedMap2(t *testing.T) {
	t.Run("print all elements in order.", func(t *testing.T) {
		om := New[string, MyStruct](10)
		om.Add("key1", MyStruct{id: 1, name: "test1"})
		om.Add("key3", MyStruct{id: 3, name: "test3"})
		om.Add("key2", MyStruct{id: 2, name: "test2"})

		for v := range om.Values() {
			fmt.Printf("Values: %s\n", v)
		}
	})
}

func TestOrderedMap3(t *testing.T) {
	t.Run("print all elements in order.", func(t *testing.T) {
		om := New[string, MyStruct](10)
		om.Add("key1", MyStruct{id: 1, name: "test1"})
		om.Add("key3", MyStruct{id: 3, name: "test3"})
		om.Add("key2", MyStruct{id: 2, name: "test2"})

		for k, v := range om.KeyValues() {
			fmt.Printf("Key: %s, Values: %s\n", k, v)
		}
	})
}


type MyStructs []MyStruct

func (ms MyStructs) Names() iter.Seq[string] {
 return func(yield func(string) bool) {
  for _, s := range ms {
   if !yield(s.name) {
    return
   }
  }
 }
}

func TestOrderedMap4(t *testing.T) {

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


