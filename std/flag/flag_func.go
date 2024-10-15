package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	fs := flag.NewFlagSet("ExampleFunc", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	var ip net.IP 
	fs.Func("ip", "`IP address` to parse", func(s string)error{
		ip = net.ParseIP(s)
		if ip == nil{
			return errors.New("count not parse IP")
		}
		return nil 
	})
	fs.Parse([]string{"-ip", "127.0.0.1"})
	fmt.Printf("{ip: %v, loopback: %t}\n\n", ip, ip.IsLoopback())

	fs.Parse([]string{"-ip", "256.0.0.1"})
	fmt.Printf("{ip: %v, loopback: %t}\n\n", ip, ip.IsLoopback())
}