package main

import (
	"fmt"
)

var (
	user string
	password string
	input string
)

func main(){
	account := make(map[string]string)
	fmt.Println("请创建账号:")
	fmt.Scan(&user)
	fmt.Println("请输入密码:")
	fmt.Scan(&password)
	fmt.Println("请确认密码:")
	for{
		fmt.Scan(&input)
		if input==password{
			account[user]=password
			break
		}
		fmt.Println("密码输入不一致，请重新输入密码")
	}
	fmt.Println("请输入账号:")
	fmt.Scan(&user)
	fmt.Println("请输入密码:")
	fmt.Scan(&password)
	if password!=account[user]{
		fmt.Println("密码错误!")
		return
	}
	fmt.Println("密码正确!")
}
