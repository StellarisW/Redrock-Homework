package main

import (
	"fmt"
	"math/rand"
	"time"
)

type IntSlice []int

func (s *IntSlice) BubbleSort() {
	l:= len(*s)
	for i:=0;i<l-1;i++{
		for j:=0;j<l-i-1;j++{
			if (*s)[j]<(*s)[j+1]{
				(*s)[j],(*s)[j+1]=(*s)[j+1],(*s)[j]
			}
		}
	}
}

func main() {
	s:=make(IntSlice,0,101)
	rand.Seed(time.Now().Unix())
	for i:=1;i<=100;i++{
		data:=rand.Intn(100)
		s = append(s, data)
	}
	fmt.Print("排序前:","\n",s,"\n")
	//sort.Sort(sort.Reverse(sort.IntSlice(s)))
	s.BubbleSort()
	fmt.Print("排序后:","\n",s)
}
