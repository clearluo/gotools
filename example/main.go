package main

import (
	"os"
	"path/filepath"
	"time"

	"gitee.com/clearluo/gotools/log"

	_ "gitee.com/clearluo/gotools/db"
	_ "gitee.com/clearluo/gotools/log"
	_ "gitee.com/clearluo/gotools/util"
)

func initDir() {
	os.MkdirAll("./logs", os.ModePerm)
}
func initLog() {
	fileName := filepath.Join("./logs", "console.%v.log")
	logConfig := log.LogConfig{
		Filename:        fileName,
		RetainFileCount: 2048,
		IsDevelop:       true,
		FileSplitTime:   time.Hour,
	}
	logConn := log.SetLogger(logConfig)
	_ = logConn
	//log.SetLevel(log.DebugLevel)
}
func init() {
	initDir()
	initLog()
}
func main() {
	newLog := log.WithField("abc", "def")
	for {
		newLog.Warn("now:", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}
}
