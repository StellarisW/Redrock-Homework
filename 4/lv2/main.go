package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	file,err:=os.Create("./plan.txt")
	if err!=nil{
		fmt.Println(err)
	}
	file.Write([]byte("I’m not afraid of difficulties and insist on learning programming"))
	file,err=os.Open("./plan.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	var tmp = make([]byte, 128)
	n,err:=file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(tmp[:n]))
}