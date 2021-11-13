package main

import (
	"fmt"
)

var(
	n,opr int
	a=make([]int,100)
)

func main() {
	fmt.Println("input")
	fmt.Scan(&n)
	for i:=0;i<n;i++{
		fmt.Scan(&opr)
		a[i]=opr
	}
	for i:=0;i<n-1;i++{
		for j:=0;j<n-i-1;j++{
			if a[j]>a[j+1]{
				a[j],a[j+1]=a[j+1],a[j]
			}
		}
	}
	fmt.Println("output")
	for key,val:=range a{
		if key==n{
			break
		}
		fmt.Printf("%d ",val)
	}
}