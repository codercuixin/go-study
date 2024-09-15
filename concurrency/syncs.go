package main

import "sync"

func main(){
	cond()
}
func cond() {
	var wg sync.WaitGroup
	cond := sync.NewCond(&sync.Mutex{})
	data := make([]int, 0)

	wg.Add(1)
	go func(){
		defer wg.Done()
		for i:= 0; i<5; i++{
			cond.L.Lock();
			data = append(data, i+100)
			cond.L.Unlock();

			//唤醒一个在等待 lock 的
			cond.Signal()
		}
	}()

	for i:= 0; i< 5; i++{
		wg.Add(1)
		go func(id int){
			defer wg.Done()
			cond.L.Lock()
			// 如果没有数据，继续等待，可能被其他 goroutine 拿去数据了
			for len(data) == 0{
				cond.Wait()
			}

			x := data[0]
			data = data[1:]
			cond.L.Unlock()
			println(id, ":", x)
		}(i)
	}
	wg.Wait()
}