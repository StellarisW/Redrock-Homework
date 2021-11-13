package main

import "fmt"

func pri(ch chan int){
	for i:=1;i<=50;i++{
		p:=<-ch
		fmt.Printf("%d ",p)
		p++
		ch<-p
	}
}

func main(){
	ch:=make(chan int,1)
	ch<-1
	pri(ch)
	pri(ch)
}
