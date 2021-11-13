package main

import "fmt"

func main() {
	over := make(chan bool)
	x:=make(chan int,1)
	for i := 0; i < 10; i++ {
		x<-i
		go func() {
			p:=<-x
			fmt.Println(p)
			if p == 9 {
				over <- true
			}
		}()
	}
	<-over
	fmt.Println("over!!!")
}
