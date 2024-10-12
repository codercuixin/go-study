package main

import (
	"sync"
	// "time"
)

func main() {
	// var m sync.Mutex
	// m.Lock()
	// //mock op
	// time.Sleep(2*time.Millisecond)
	// m.Unlock()

	// //mock slower op, but don't need lock
	// time.Sleep(10 *time.Second)

	testRecursiveLock()
}

func testRecursiveLock(){
	var m sync.Mutex
	m.Lock()
	{
		//fatal error: all goroutines are asleep - deadlock!
		m.Lock()
		m.Unlock()
	}
	m.Unlock()
}