package main

import (
	"fmt"
	"reflect"
)
type U struct{
	name string
	age int 
}

type M struct{
	U 
	title string 
}

func pp(t reflect.Type, prefix string){
	for i :=0; i<t.NumField(); i++{
		f := t.Field(i)
		fmt.Println(prefix, f.Name)
		//is an embedded field
		// fmt.Println(f.Name, f.Anonymous)
		if f.Anonymous{
			pp(f.Type, prefix + "	")
		}
	}
}

func main(){
	// testPrintField()
	// findStructField()
	extractStructTag()
}

func testPrintField(){
	var m M 
	t := reflect.TypeOf(&m)

	if t.Kind() == reflect.Pointer{
		// 如果是 pointer 就获取基类型
		t = t.Elem()
		fmt.Println("changed type")
	}
	pp(t, "")
}

func findStructField(){
	var m M 
	t := reflect.TypeOf(m)
	// 按名称查找
	name, _ := t.FieldByName("name")
	fmt.Println(name.Name, name.Offset)

	// 按多级索引查找
	age := t.FieldByIndex([]int{0, 1})
	fmt.Println(age.Name, age.Offset)
}

type User struct{
	id int `field:"uid" type:"int"`
	name string `field:"username" type:"varchar(50)"`
}

func extractStructTag(){
	var u User 
	t := reflect.TypeOf(u)
	for i:= 0; i<t.NumField(); i++{
		f := t.Field(i)
		fmt.Println(f.Name, f.Tag.Get("field"), f.Tag.Get("type"))
	}
}
