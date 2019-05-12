package main

import (
	"fmt"
	"sync"
)

func main() {
	//testOnce1()
	//testOnce2()
	testOnce3()
}

// 测试sync.Once基本功能
func testOnce1() {
	var count int

	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d \n", count)

}

// 注意：sync.Once与调用的第一个函数关联在一起
func testOnce2() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count: %d\n", count)

}

// 注意下面的例子会死锁
func testOnce3() {
	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) }
	onceA.Do(initA)
}
