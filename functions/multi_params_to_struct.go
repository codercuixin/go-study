

package main

import "fmt"
import "time"
import "log"
type Option struct{
	addr string 
	port int 
	path string 
	timeout time.Duration 
	log *log.Logger 
}

func newOption() *Option{
	return &Option{
		addr: "0.0.0.0",
		port: 8080,
		path: "/var/test",
		timeout: time.Second * 5,
		log: nil,
	}
}

func server(option *Option){
	fmt.Println(option)
}

func main(){
	opt := newOption()
	opt.port = 8085 
	server(opt)
}