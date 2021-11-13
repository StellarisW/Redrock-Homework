package main

import (
	"fmt"
)

var (
	coins = 0
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (cost int) {
	for _, name := range users {
		for _, c := range name {
			switch c {
			case 'e', 'E':
				distribution[name] += 1
				coins += 1
			case 'i', 'I':
				distribution[name] += 2
				coins += 2
			case 'o', 'O':
				distribution[name] += 3
				coins += 3
			case 'u', 'U':
				distribution[name] += 4
				coins += 4
			}
		}
	}
	cost = coins
	return
}

func main() {
	cost := dispatchCoin()
	fmt.Printf("一共花费了%d个金币\n", cost)
	for val:=range distribution{
		fmt.Printf("%s分到了%d个金币\n",val,distribution[val])
	}
}
