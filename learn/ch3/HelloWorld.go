package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go sayHello(&wg)
	//wg.Wait()
	//t1()
	memoryTest()
}

// 测试goroutine占用内存情况
func memoryTest() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c }

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}

func t1() {
	var strs []*string
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		strs = append(strs, &salutation)
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
	fmt.Println(strs)
}

func sayHello(wg *sync.WaitGroup) {
	fmt.Println("hello")
	wg.Done()
}
