package main

import (
	"sync"
	"time"
	"fmt"
	"reflect"
)

func main() {
	// rec_send1_useFor()
	// rec_send_useForRange()
	// limit_direction()
	// randomSelect()
	// randomSelectEvenSameChan()
	// randomSelectWithDefault()
	reflectSelect()
}

func rec_send1_useFor(){
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan int)

	//recv
	go func(){
		defer wg.Done()
		for{
			x, ok := <-c 
			if !ok{
				return
			}
			println(x)
		}
	}()

	//send
	go func (){
		defer wg.Done()
		defer close(c)
		c <- 1
		c <- 2
		c <- 3
	}()

	wg.Wait()
}

func rec_send_useForRange(){
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan int)

	//recv
	go func(){
		defer wg.Done()
		for x:= range c{
			println(x)
		}
		
	}()

	//send
	go func (){
		defer wg.Done()
		defer close(c)
		c <- 1
		c <- 2
		c <- 3
	}()

	wg.Wait()
}

func limit_direction(){
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)
	var send chan <- int = c
	var recv <- chan int = c 
	
	// recv
	go func(){
		defer wg.Done()
		for x := range recv{
			println(x)
		}
	}()

	//send
	go func(){
		defer wg.Done()
		defer close(send)
		send <- 1 
		send <- 2
		send <- 3
	}()
	
	wg.Wait()
}

func single_dir_cannot_convert_two(){
	// var a, b chan int 
	// a = make(chan int, 2)
	// var recv <- chan int = a 
	// var send chan <- int = a 
	// //cannot convert recv (variable of type <-chan int) to type chan intcompilerInvalidConversion

	// b = (chan int)(recv)
	// //annot convert send (variable of type chan<- int) to type chan 
	// b = (chan int)(send)

}

func makeSingleDir(){
	// //没有意义
	// q := make(<- chan struct{})
	// go func(){
	// 	//invalid operation: cannot close receive-only channel q (variable of type <-chan struct{})
	// 	close(q)
	// }()
	// <-q 
}


func randomSelect(){
	var wg sync.WaitGroup
	wg.Add(2)

	cha, chb := make(chan int), make(chan int)

	//recv 
	go func(){
		defer wg.Done()
		for{
			x := 0
			ok := false
			//random
			select{
			case x, ok = <- cha :
				if!ok{
					cha = nil 
				}
			case x, ok = <- chb :
				if !ok{
					chb = nil
				}
			}
			if cha == nil && chb == nil{
				return
			}
			println(x)
		}
	}()

	//send
	go func(){
		defer wg.Done()
		defer close(cha)
		defer close(chb)

		for i:=0; i<10; i++{
			//random send
			select{
			case cha <- i:
			case chb <-i+10:
			}
		}
	}()
	//wait in main
	wg.Wait()
}

func randomSelectEvenSameChan(){
	var wg sync.WaitGroup
	wg.Add(2)

	cha := make(chan int)

	//recv 
	go func(){
		defer wg.Done()
		for{
			x := 0
			ok := false
			//random
			select{
			case x, ok = <- cha :
				println("branch1 read", x)
			case x, ok = <- cha :
				println("branch2 read", x)
			}
			if !ok{
				return
			}
		}
	}()

	//send
	go func(){
		defer wg.Done()
		defer close(cha)

		for i:=0; i<10; i++{
			//random send
			select{
			case cha <- i:
			case cha <-i+10:
			}
		}
	}()
	//wait in main
	wg.Wait()
}

func randomSelectWithDefault(){
	exit := make(chan struct{})
	c := make(chan int)
	go func(){
		defer close(exit)
		for{
			select{
			case x, ok := <-c :
				if(!ok){
					return
				}
				println(x)
			default:
			}

			fmt.Println("wait...")
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 3)
	c <- 100
	close(c)
	<-exit 
}

func reflectSelect(){
	exit := make(chan struct{})
	chans := make([]chan int, 0)
	chans = append(chans, make(chan int))
	chans = append(chans, make(chan int))
	go func(){
		defer close(exit)
		// read cases 
		cases := make([]reflect.SelectCase, len(chans))
		for i, c := range chans{
			cases[i] = reflect.SelectCase{
				Dir: reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			}
		}
		for{
			index, value, ok := reflect.Select(cases)
			// 检查，如果都退出了，就 return
			if !ok{
				chans[index] = nil 
				n := 0 
				for _, c := range chans{
					if c== nil{
						n++
					}
					if n==len(chans){
						return
					}
				}
				continue
			}

			println(index, value.Int(), ok)
			
		}

	}()

	chans[1] <- 101
	chans[0] <- 100
	for _, c := range chans{
		close(c)
	}
	//wait exit
	<-exit 
}