package tool

import (
	"io"
	"os"
	"time"
)

const (
	LOGPATH = "log/"
	FORMAT = "20060102"
	FORMATS="2006-01-02 15:04:05"
	LineFeed = "\r\n"
)


func Log(level string, msg string) error {
	if !IsExist(LOGPATH) {
		return CreateDir(LOGPATH)
	}
	var (
		err error
		f   *os.File
	)
	fileName := LOGPATH + time.Now().Format(FORMAT) + ".log"
	f, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	msg="["+level+"]-["+time.Now().Format(FORMATS)+"]-[" +msg+"]"
	_, err = io.WriteString(f, msg+LineFeed)
	defer f.Close()
	return err
}

//CreateDir  文件夹创建
func CreateDir(LOGPATH string) error {
	err := os.MkdirAll(LOGPATH, os.ModePerm)
	if err != nil {
		return err
	}
	os.Chmod(LOGPATH, os.ModePerm)
	return nil
}

//IsExist  判断文件夹/文件是否存在  存在返回 true
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
