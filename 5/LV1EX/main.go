package main

import (
	"fmt"
	"strings"
)

type skill string

var (
	skillBase map[string]string
	tempBase  map[string]string
)

func Check(s string) bool {
	wordList := []string{
		"卷",
	}
	for _, v := range wordList {
		if strings.Index(s, v) != -1 {
			fmt.Printf("存在敏感词: %s\n", v)
			return false
		}
	}
	return true
}

func TCreate(skillNames string, skillKinds string) {
	if skillKinds == "" {
		fmt.Println("请输入你想创建的技能类别:")
		fmt.Scanf("%s", &skillKinds)
	}
	fmt.Println("请自定义该技能类别的文字模板\n%表示技能名\n如:\n输入: 你死定了！%我不知道你该怎么从%活下来\n" +
		"表示: 你死定了！+技能名+我不知道你该怎么从+技能名+活下来")
	var in string
	fmt.Scanf("%s", &in)
	in = strings.Replace(in, "%", skillNames, -1)
	tempBase[skillKinds] = in
	skillBase[skillNames] = skillKinds
}

func (skillNames skill) Release(temp string) {
	out := tempBase[temp]
	strings.Replace(out, "%", string(skillNames), -1)
	fmt.Println(out)
}

func (skillNames skill) Create() {
	if !Check(string(skillNames)) {
		fmt.Println("请重新输入")
		return
	}
	fmt.Println("请输入该技能的类别:")
	var skillKind string
	fmt.Scanf("%s", &skillKind)
	if Check(skillKind) {
		_, ok := tempBase[skillKind]
		if ok {
			tempBase[string(skillNames)] = skillKind
		} else {
			fmt.Println("没有该技能类别,是否创建？\nYes/No")
			var opr string
			fmt.Scanf("%s", &opr)
			if opr == "Yes" {
				TCreate(string(skillNames), skillKind)
			} else {
				fmt.Println("创建技能失败！")
			}
		}
	}
}

func main() {
	skillBase = make(map[string]string)
	tempBase = make(map[string]string)
	for {
	flag:
		func() {
			fmt.Println("请选择操作:\n1.释放技能\n2.添加技能\n3.退出游戏")
		}()
		var opr int
		fmt.Scanf("%d", &opr)
		switch opr {
		case 1:
			{
				fmt.Println("请输入技能名:")
				var skillName string
				var opr string
				fmt.Scanf("%s", &skillName)
				skillKind, ok := skillBase[skillName]
				if ok {
					skill(opr).Release(skillKind)
				} else {
					fmt.Println("你没有拥有该技能,是否创建？\nYes/No")
					fmt.Scanf("%s", &opr)
					if opr=="Yes"&&Check(opr){
						skill(skillName).Create()
					} else {
						fmt.Println("释放技能失败！")
					}
				}
			}
		case 2:
			{
				fmt.Println("请输入创建的技能名:")
				for{
					var opr string
					fmt.Scanf("%s", &opr)
					_, ok := skillBase[opr]
					if ok {
						fmt.Println("已存在该技能,请重新输入")
					} else if !Check(opr) {
						fmt.Println("请重新输入")
					} else {
						skill(opr).Create()
						break
					}
				}

			}
		case 3:
			break
		default:
			{
				fmt.Println("输入错误！\n请重新输入!")
				goto flag
			}
		}
	}
}
