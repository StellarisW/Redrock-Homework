package main

import "fmt"

var (
	a, b int
	ans int
	s       string
)

func main() {
	fmt.Println("input:")
	fmt.Scan(&a)
	fmt.Scan(&s)
	fmt.Scan(&b)
	switch s {
	case "+":
		ans = a + b
	case "-":
		ans = a - b
	case "*":
		ans = a * b
	case "/":
		ans = a / b
	default:
		fmt.Println("无效的输入!")
	}
	fmt.Printf("output:\n%d",ans)
}
