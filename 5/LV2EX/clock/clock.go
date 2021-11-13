package clock

import (
	"fmt"
	"main/list"
	"runtime"
	"time"
)

type Timer struct {
	Timer *time.Timer
	t     time.Time
	D     chan int
	kind  int //0:一次性 1:每天 2:每周
	info  string
}

var (
	L = list.NewList()
)

func read() int {
	var x int
	fmt.Scanf("%d", &x)
	return x
}
func readWeekday(now time.Time) int {
	var x int
	fmt.Scanf("%d", &x)
	n := int(now.Weekday())
	if n == 0 {
		n = 7
	}
	if n > x {
		return now.AddDate(0, 0, x+7-n).Day()
	} else if n < x {
		return now.AddDate(0, 0, x-n).Day()
	} else {
		return now.Day()
	}
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("出现错误:")
	}
}

func StartInterface() {
	fmt.Println("请执行操作(输入数字):")
	fmt.Println("1.设置一个新的计时器")
	fmt.Println("2.删除一个计时器")
	fmt.Println("3.取消指定计时器的下一次提醒")
}

func SetTimer(kind int) {
	defer recovery()
	var timer Timer
	var info string
	now := time.Now()
	timer.D=make(chan int)
	if kind == 1 {
		fmt.Println("请输入:\n年/月/日/时/分/秒(用数字表示,空格隔开)")
		t := time.Date(read(), time.Month(read()), read(), read(), read(), read(), 0, time.Local)
		if t.Before(now) {
			panic("输入的时间已过期！！！\n")
		}
		timer.Timer = time.NewTimer(t.Sub(now))
		fmt.Println("请输入需要提醒的信息")
		fmt.Scanf("%s", &info)
		timer.t = t
		timer.kind = 0
		timer.info = info
		L.Append(list.NewNode(timer))
		go func(){
			for {
				select {
				case <-timer.Timer.C:
					fmt.Println(timer.info)
				case <-timer.D:
					runtime.Goexit()
				}
			}
		}()
	} else {
		fmt.Println("请选择重复时间:")
		fmt.Println("1:每天 2:每周")
		fmt.Scanf("%d", &kind)
		switch kind {
		case 1:
			{

				fmt.Println("请输入:\n时/分/秒(用数字表示,空格隔开)")
				t := time.Date(now.Year(), now.Month(), now.Day(), read(), read(), read(), 0, time.Local)
				if t.Before(now) {
					t = t.Add(24 * time.Hour)
				}
				timer.Timer = time.NewTimer(t.Sub(now))
				fmt.Println("请输入需要提醒的信息")
				fmt.Scanf("%s", &info)
				timer.t = t
				timer.kind = 1
				timer.info = info
				L.Append(list.NewNode(timer))
				go func() {
					for {
						select {
						case <-timer.Timer.C:
							fmt.Println(timer.info)
							timer.Timer.Reset(24 * time.Hour)
						case <-timer.D:
							runtime.Goexit()
						}
					}
				}()
			}
		case 2:
			{
				fmt.Println("请输入:\n周几/时/分/秒(用数字表示,空格隔开)")
				t := time.Date(now.Year(), now.Month(), readWeekday(now), read(), read(), read(), 0, time.Local)
				if t.Before(now) {
					timer.Timer = time.NewTimer(168*time.Hour - now.Sub(t))
				} else {
					timer.Timer = time.NewTimer(t.Sub(now))
				}
				fmt.Println("请输入需要提醒的信息")
				fmt.Scanf("%s", &info)
				timer.t = t
				timer.kind = 2
				timer.info = info
				L.Append(list.NewNode(timer))
				go func(){
					for {
						select {
						case <-timer.Timer.C:
							fmt.Println(timer.info)
							timer.Timer.Reset(168 * time.Hour)
						case <-timer.D:
							runtime.Goexit()
						}
					}
				}()
			}
		default:
			panic("输入错误\n")
		}
	}
}

func TimerOut() {
	defer recovery()
	pre := L.Head
	for i := 1; nil != pre; i++ {
		fmt.Printf("%d.", i)
		i := pre.GetValue()
		timer, _ := i.(Timer)
		switch timer.kind {
		case 0:
			timeFormat := "2006-01-02 15:04:05"
			fmt.Println("一次性闹钟")
			fmt.Println("  提醒时间: ", timer.t.Format(timeFormat))
		case 1:
			timeFormat := "15:04:05"
			fmt.Println("每天提醒的闹钟")
			fmt.Println("  提醒时间: 每天", timer.t.Format(timeFormat))
		case 2:
			timeFormat := "15:04:05"
			fmt.Println("这是一个每周提醒的闹钟")
			fmt.Println("  提醒时间: 每周", timer.t.Weekday(), timer.t.Format(timeFormat))
		}
		pre = pre.Next
	}
}

func (timer Timer) CancelNext() {
	recovery()
	switch timer.kind {
	case 0:
		panic("该闹钟为一次性闹钟！\n")
	case 1:
		now := time.Now()
		t := time.Date(now.Year(), now.Month(), now.Day(), timer.t.Hour(), timer.t.Minute(), timer.t.Second(), 0, timer.t.Location())
		if t.Before(now) {
			t = t.Add(24 * time.Hour)
		}
		timer.Timer = time.NewTimer(t.Sub(now))
	case 2:
		now := time.Now()
		n := int(now.Weekday())
		x := int(timer.t.Weekday())
		if n == 0 {
			n = 7
		}
		if n > x {
			x = now.AddDate(0, 0, x+7-n).Day()
		} else if n < x {
			x = now.AddDate(0, 0, x-n).Day()
		} else {
			x = now.Day()
		}
		t := time.Date(now.Year(), now.Month(), x, timer.t.Hour(), timer.t.Minute(), timer.t.Second(), 0, timer.t.Location())
		if t.Before(now) {
			timer.Timer = time.NewTimer(168*time.Hour - now.Sub(t))
		} else {
			timer.Timer = time.NewTimer(t.Sub(now))
		}
		L.Append(list.NewNode(timer))
	}
}
