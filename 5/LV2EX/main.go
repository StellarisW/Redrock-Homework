package main

import (
	"fmt"
	"main/clock"
	"time"
)

func main() {
	fmt.Println(time.Now())
	for {
		clock.StartInterface()
		var opr int
		fmt.Scanf("%d", &opr)
		switch opr {
		case 1:
			{
				fmt.Println("请选择计时器种类(输入数字):")
				fmt.Println("1.一次性的定时器")
				fmt.Println("2.重复提醒的定时器")
				fmt.Scanf("%d", &opr)
				clock.SetTimer(opr)
			}
		case 2:
			{
				clock.TimerOut()
				fmt.Scanf("%d", &opr)
				item, _ := clock.L.Find(uint(opr))
				i := item.GetValue()
				timer := i.(clock.Timer)
				timer.D<-1
				timer.Timer.Stop()
				clock.L.Delete(uint(opr))
			}
		case 3:
			{
				clock.TimerOut()
				fmt.Scanf("%d", &opr)
				item, _ := clock.L.Find(uint(opr))
				i := item.GetValue()
				timer := i.(clock.Timer)
				timer.CancelNext()
			}
		}
	}
}