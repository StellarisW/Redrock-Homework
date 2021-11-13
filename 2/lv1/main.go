package main

import "fmt"

func main() {
	a:="hello"
	l:=len(a)
	for i:=range a{
		fmt.Printf("%c",a[l-i-1])
	}
}