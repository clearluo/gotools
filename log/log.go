package log

import (
	// "path"

	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	// "time"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&DFormatter{})
	logrus.SetLevel(DebugLevel)
	logrus.SetOutput(os.Stdout)
}

//type PanicLevel = logrus.PanicLevel
const (
	PanicLevel logrus.Level = logrus.PanicLevel
	FatalLevel logrus.Level = logrus.FatalLevel
	ErrorLevel logrus.Level = logrus.ErrorLevel
	WarnLevel  logrus.Level = logrus.WarnLevel
	InfoLevel  logrus.Level = logrus.InfoLevel
	DebugLevel logrus.Level = logrus.DebugLevel
)

var (
	fileConn      *rotatelogs.RotateLogs
	isPrintStdout bool
)

type LogConfig struct {
	Filename        string        // 日志文件
	RetainFileCount uint          // 保留日志文件数量
	IsDevelop       bool          // 开发模式会输出日志到标准输出
	FileSplitTime   time.Duration // 日志文件拆分时间间隔 time.Hour*24
}

func WithField(key string, value string) *logrus.Entry {
	return logrus.WithField(key, value)
}

// config need to be correct JSON as string: {"filename":"default.log","maxsize":100}
func SetLogger(config LogConfig) *logrus.Logger {
	/* 日志轮转相关函数
	   `WithLinkName` 为最新的日志建立软连接
	   `WithRotationTime` 设置日志分割的时间，隔多久分割一次
	   WithMaxAge 和 WithRotationCount二者只能设置一个
	     `WithMaxAge` 设置文件清理前的最长保存时间
	     `WithRotationCount` 设置文件清理前最多保存的个数
	*/
	//fileName := config.Filename + ".%Y%m%d"
	fileName := ""
	if config.FileSplitTime > time.Hour {
		fileName = fmt.Sprintf(config.Filename, "%Y%m%d")
	} else if config.FileSplitTime > time.Minute*30 {
		fileName = fmt.Sprintf(config.Filename, "%Y%m%d_%H")
	} else {
		fileName = fmt.Sprintf(config.Filename, "%Y%m%d_%H_%M")
	}
	fileConn, _ := rotatelogs.New(
		fileName,
		//rotatelogs.WithLinkName(config.Filename),
		//rotatelogs.WithRotationCount(config.RetainFileCount),
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationTime(config.FileSplitTime),
	)
	if config.IsDevelop {
		isPrintStdout = true
	}
	logrus.SetOutput(fileConn)
	return logrus.StandardLogger()
}
func GetFileConn() *rotatelogs.RotateLogs {
	return fileConn
}
func SetLevel(logLevel logrus.Level) {
	logrus.SetLevel(logLevel)
}

type DFormatter struct {
	TimestampFormat string
}

func (f *DFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = "[2006/01/02 15:04:05.000]"
	}

	_, file, line, ok := runtime.Caller(9)
	if !ok {
		file = "???"
		line = 0
	}
	file = filepath.Base(file) // 只取文件路径的文件名
	// _, filename := path.Split(file)
	msg := entry.Time.Format(timestampFormat) +
		" " + strings.ToUpper(entry.Level.String()) +
		" [" + file + ":" + strconv.Itoa(line) + "] " +
		entry.Message + "\n"
	if isPrintStdout {
		fmt.Printf(msg)
	}
	return []byte(msg), nil
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

// 以下函数不可以写成这种形式
// func Debug(args ...interface{}) {
// 	logrus.Debug(args...)
// }

func Debug(args ...interface{}) {
	debug(args...)
}

func debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Println(args ...interface{}) {
	println(args...)
}

func println(args ...interface{}) {
	logrus.Println(args...)
}

func Info(args ...interface{}) {
	info(args...)
}

func info(args ...interface{}) {
	logrus.Info(args...)
}

func Warn(args ...interface{}) {
	warn(args...)
}

func warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Error(args ...interface{}) {
	ferror(args...)
}

func ferror(args ...interface{}) {
	logrus.Error(args...)
}

func Fatal(args ...interface{}) {
	fatal(args...)
}

func fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func Panic(args ...interface{}) {
	panic(args...)
}

func panic(args ...interface{}) {
	logrus.Panic(args...)
}