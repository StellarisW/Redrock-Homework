package utils

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"os"
)

//func Upload(r *ghttp.Request) {
//	r.up
//}

func UploadString(fileName string, fileString string) {
	ctx := gctx.New()
	paths := "C:/Users/Stellaris_W/OneDrive - stu.cqupt.edu.cn/桌面/Coding/Jetbrains/Goland/src/github.com/Redrock/Redrock-Homework/13/app/resource/public/resource/upload/"
	paths += fileName
	dstFile, err := os.Create(paths)
	if err != nil {
		g.Log().Errorf(ctx, "Create File of string failed, err: %v\n", err)
		return
	}
	defer dstFile.Close()
	dstFile.WriteString(fileString)
}

//func GetPath() string {
//	path.Join()
//}
