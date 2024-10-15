package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

var species = flag.String("species", "gohper", "the species we are studying")
var goherType string 
func init(){
	const(
		defaultGopher = "pocket"
		usage = "the variety of gohper"
	)
	flag.StringVar(&goherType, "gohper_type", defaultGopher, usage)
	flag.StringVar(&goherType, "g", defaultGopher, usage+" (shorthand)")
}

type interval []time.Duration 
func(i *interval) String() string{
	return fmt.Sprint(*i)
}
func(i *interval) Set(value string)error{
	if(len(*i) > 0){
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ","){
		duration, err := time.ParseDuration(dt)
		if err != nil{
			return err 
		}
		*i = append(*i, duration)
	}
	return nil 
}


var intervalFlag interval 
func init(){
	flag.Var(&intervalFlag, "deltaT", `comma-separated list 
	of intervals to use between events`)

}

func main(){
	//  go run demo2.go -species=bee -gohper_type=newbee -deltaT=20m,15s
	flag.Parse()
	fmt.Printf("species is %s\n", *species)
	fmt.Printf("gohper type is %s\n", goherType)
	fmt.Printf("interval is %s\n", intervalFlag)
}