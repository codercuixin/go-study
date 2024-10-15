package main

import (
	"bytes"
	"log"
	"fmt"
)

func main() {
	// basic()
	logCalldepth()
}

func logCalldepth(){
	var buf bytes.Buffer
	// var logger = log.New(&buf, "logger-prefix: ", log.Ldate|log.Ltime|log.Lshortfile)
	var logger = log.New(&buf, "logger-prefix: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)

	var infof = func(info string){
		// 1 is demo.go
		// 2 is demo.go
		// 3 is demo.go
		// 4 is proc.go
		// 5 is asm_amd64.s
		logger.Output(1, info)
	}
	infof("hello world")
	fmt.Print(&buf)
}


func basic(){
	var buf bytes.Buffer
	// var logger = log.New(&buf, "logger-prefix: ", log.Ldate|log.Ltime|log.Lshortfile)
	var logger = log.New(&buf, "logger-prefix: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)

	logger.Print("Hello, log file!")

	fmt.Print(&buf)
}