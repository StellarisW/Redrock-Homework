package main

import (
	"fmt"
	"reflect"
)

func Receiver(i interface{}) {
	v := reflect.ValueOf(i)
	ty := v.Kind()
	switch ty {
	case reflect.Int:
		println("这是一个int类型")
	case reflect.String:
		println("这是一个String类型")
	case reflect.Bool:
		println("这是一个bool类型")
	case reflect.Int64:
		println("这是一个Int64类型")

	}
}

func main() {
	var (
		a int
		b string
		c bool
		d int64
	)
	fmt.Scan(a,b,c,d)
	Receiver(a)
	Receiver(b)
	Receiver(c)
	Receiver(d)
}
