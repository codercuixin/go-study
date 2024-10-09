package main 

import(
	"embed"
	"fmt"
)
//go:embed main.go
var main_go string 

//go:embed main.go
var main_go_bytes []byte


//go:embed mylib/* mylib2/*
//go:embed *.go
var dir embed.FS
func main(){
	fmt.Println(main_go)
	fmt.Println("------------")
	fmt.Println(string(main_go_bytes))
	fmt.Println("------------")
	bytes, _ := dir.ReadFile("main.go");
	fmt.Println(string (bytes))
	fmt.Println("------------")
	bytes, _ = dir.ReadFile("mylib2/add.go")
	fmt.Println(string (bytes))
}

