package utils

import (
	"errors"
	"fmt"
	"os"
	"time"
)

//获取日期
func GetDate() (date string) {
	now := time.Now()
	date = fmt.Sprintf("%d%d%d%d%d%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Month(), now.Second())
	return
}

//创建文件夹
func MkdirFolder(keyword string) (path string, err error) {
	path = "../../file/" + keyword + GetDate()
	//path = "./"+ keyword + GetDate()
	err = os.Mkdir(path, 777)
	if err != nil {
		err = errors.New("文件夹创建失败")
		return
	}
	//通过chmod重新赋权限
	_ = os.Chmod(path, 0777)
	return
}
