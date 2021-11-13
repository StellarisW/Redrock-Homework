package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var x int64

func add(ch chan int64) {
	for i := 0; i < 50000; i++ {
		p:=<-ch
		p++
		ch<-p
	}
	wg.Done()
}
func main() {
	ch:=make(chan int64,1)
	ch <- x
	wg.Add(2)
	go add(ch)
	go add(ch)
	wg.Wait()
	x=<-ch
	fmt.Println(x)
}
