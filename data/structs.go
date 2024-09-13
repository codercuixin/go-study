package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type node struct{
	id int 
	value string 
	next *node 
}

func main(){
	// basic()
	// anom_struct()
	// anom_struct2()
	// cannot_cmp_struct_with_map()
	// cmp_named_struct_with_unnamed()
	// null_struct_size()
	// struct_tag()
	reflect_use_tag()
}

func reflect_use_tag(){
	type User struct{
		id int 	`field:"uid" type:"integer"`
		name string `field:"name" type:"text"`
	}
	t := reflect.TypeOf(User{})
	for i:=0; i<t.NumField(); i++{
		f := t.Field(i)
		fmt.Println(f.Name, f.Tag.Get("field"), f.Tag.Get("type"))
	}
}

func struct_tag(){
	type  user struct{
		id int `id`
	}

	type user2 struct{
		id int 	`uid`
	}

	u1 := user{1}
	u2 := user2{2}
 	//invalid operation: u1 == u2 (mismatched types user and user2)
	// _ = u1 ==  u2 

	//内存布局相同，支持转换
	u1 = user(u2)
	fmt.Println(u1)
}



func null_struct_size(){
	var a struct{}
	var b [10000]struct{}
	s := b[:]
	println(unsafe.Sizeof(a), unsafe.Sizeof(b))
	println(len(s), cap(s))
}

func cmp_named_struct_with_unnamed(){
	type data struct{
		x int 
	}

	d1 := data{x:100}
	d2 := struct{
		x int 
	}{x:100}

	fmt.Println(d1 == d2)
}

func cannot_cmp_struct_with_map(){
	// type data struct{
	// 	x int 
	// 	y map[string]int 
	// }

	// d1 := data{x:100}
	// d2 := data{x:100}

	//    invalid operation: d1 == d2 (struct containing map[string]int cannot be compared)
	// _ = d1==d2 

}

func anom_struct2(){
	type file struct{
		name string 
		attr struct{
			owner int 
			perm int 
		}
	}
	f := file{
		name: "test.dat",
		//missing type in composite literal
		// attr:{
		// 	owner:1,
		// 	perm: 0755,
		// },
	}
	f.attr.owner =1
	f.attr.perm = 0755
	fmt.Println(f)
}

func anom_struct(){
	//匿名结构
	user := struct{
		id int 
		name string 
	}{
		id: 1,
		name : "user1",
	}
	fmt.Printf("%+v\n", user)

	var color struct{
		r int 
		g int 
		b int 
	}
	color.r = 1 
	color.g = 2
	color.b = 3
	fmt.Printf("%+v\n", color)
}

func basic(){
//too few values in struct literal of type node
	// _ = node{1, "a"}
	n := node{
		id :2 ,
		value : "abc", // 注意结尾逗号
	}
	fmt.Println(n)
}