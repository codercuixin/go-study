package main

import "sync"

func main(){
	var wg sync.WaitGroup
	// 使用外部锁作为静态资源保护
	cond := sync.NewCond(&sync.Mutex{})
	data := make([]int, 0)
	wg.Add(1)
	go func(){
		defer wg.Done()
		for i:=0; i<5; i++{
			cond.L.Lock()
			data = append(data, i+100)
			cond.L.Unlock()
			//唤醒一个 
			cond.Signal()
		}
		//唤醒所有（剩余）
		// cond.Broadcast()
	}()

	for i:=0; i<5; i++{
		wg.Add(1)
		go func(id int){
			defer wg.Done()
			cond.L.Lock()
			for(len(data) == 0){
				cond.Wait()
			}
			x:= data[0]
			data = data[1:]
			cond.L.Unlock()
			println("reader-", id, " read", x)
		}(i)
	}
	wg.Wait()
}