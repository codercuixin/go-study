package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	//basic()
	// jsonStructTag()
	advance()
}

func basic() {
	x := 100
	d := struct {
		X *int
		S string
		Y []int
	}{
		X: &x,
		S: "hello world!",
		Y: []int{1, 2, 3, 4},
	}

	b, err := json.Marshal(d)
	if err != nil {
		log.Fatalln(err)
	}
	println(string(b))

	buf := bytes.NewBuffer(nil)
	json.Indent(buf, b, "", "	")
	println(buf.String())
}

func jsonStructTag() {
	x := 100
	d := struct {
		X *int   `json:"id"`
		S string `json:"name"`
		Y []int  `json:"data,omitempty"`
	}{
		X: &x,
		S: "hello world!",
	}

	b, err := json.Marshal(d)
	if err != nil {
		log.Fatalln(err)
	}
	println(string(b))

	buf := bytes.NewBuffer(nil)
	json.Indent(buf, b, "", "	")
	println(buf.String())
}

type U struct {
	Id   int
	Nane string
}
type M struct {
	U
	Title string
}

func advance() {
	d := M{U{1, "user1"}, "cxo"}

	b, err := json.Marshal(d)
	if err != nil {
		log.Fatalln(err)
	}
	println(string(b))

	func() {
		var x M
		if err := json.Unmarshal(b, &x); err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%+v\n", x)
	}()
	func() {
		//ID 可以识别到。后续字母忽略大小写
		var x struct {
			ID    int
			Title string
		}
		if err := json.Unmarshal(b, &x); err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%+v\n", x)
	}()
}
