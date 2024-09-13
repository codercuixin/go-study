package main

import (
	"log"
	"math/rand"
	"sync"
	"time"

)
type CircleQueue struct{
	sync.Mutex 
	data []int 
	head int 
	tail int 
}
// use slice as stack

func NewCircleQueue(cap int) *CircleQueue{
	return &CircleQueue{data: make([]int, cap)}
}

func (q *CircleQueue) Put(v int) bool{
	q.Lock()
	defer q.Unlock()
	
	if q.tail -q.head == len(q.data){
		return false
	}

	q.data[q.tail% len(q.data)] = v 
	q.tail++
	return true 
}

func (q *CircleQueue) Get() (int,  bool){
	q.Lock()
	defer q.Unlock()

	if q.tail - q.head == 0{
		return 0, false 
	}

	v := q.data[q.head % len(q.data)]
	q.head++
	return v,true 
}


func main(){
	rand.Seed(time.Now().UnixNano())
	const max = 100000
	src := rand.Perm(max) // 随机测试数据
	dst := make([]int,0 ,max)

	q := NewCircleQueue(6)
	var wg sync.WaitGroup
	wg.Add(2)
	
	//put
	go func(){
		defer wg.Done()
		for _,v := range src{
			for{
				if ok := q.Put(v); !ok{
					continue
				}
				break
			}
		}
	}()

	//get
	go func(){
		defer wg.Done()
		for len(dst) < max{
			if v, ok:= q.Get(); ok{
				dst =append(dst, v)
				continue
			}
		}
	}()
	wg.Wait()

	if *(*[max]int)(src) != *(*[max]int)(dst){
		log.Fatalln("fail")
	}
	log.Printf("%+v\n", *q)
}