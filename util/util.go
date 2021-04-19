package util

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"

	"gitee.com/clearluo/gotools/log"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func AssertMarshal(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// 获取整点以来的秒数
func GetSecondFromHour() int64 {
	return time.Now().Unix() % 3600
}

//获取整天以来的秒数
func GetSecondFromDay() int64 {
	return time.Now().Unix() - GetSecondByDay00()
}

//获取整周以来的秒数
func GetSecondFromWeek() int64 {
	day := int64(time.Now().Weekday())
	return 86400*(day-1) + GetSecondFromDay()
}

// 获取整月以来的秒数
func GetSecondFromMonth() int64 {
	day := int64(time.Now().Day())
	return 86400*(day-1) + GetSecondFromDay()
}

//获取当前日期(20170802)零点对应的Unix时间戳
func GetSecondByDay00() int64 {
	timeStr := time.Now().Format("2006-01-02")
	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return t.Unix()
}

// 获取当前日期前后n天对应的日期证书,0代表获取当前日期整数
func GetDateByN(n int) int64 {
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, n)
	dayStr := yesTime.Format("20060102")
	day, _ := strconv.ParseInt(dayStr, 10, 64)
	return day
}

// 根据时间戳获取对应日期整数
func GetTDayByUnixTime(nowUnix int64) int64 {
	if nowUnix < 1 {
		return 0
	}
	tm := time.Unix(nowUnix, 0)
	nowDay, err := strconv.ParseInt(tm.Format("20060102"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return nowDay
}

// 统计某函数执行时间
// 使用方式
// defer utils.Profiling("test")()
func Profiling(msg string) func() {
	start := time.Now()
	return func() {
		log.Info(fmt.Sprintf("%s[%s]:%s", msg, "use", time.Since(start)))
	}
}

// 判断文件夹是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func GetFieldName(t reflect.Type) (map[string]string, error) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		err := fmt.Errorf("Check type error not Struct")
		fmt.Println(err)
		fmt.Println(err)
		return nil, err
	}
	fieldNum := t.NumField()
	result := make(map[string]string, 0)
	for i := 0; i < fieldNum; i++ {
		result[t.Field(i).Tag.Get("json")] = t.Field(i).Type.Name()
	}
	return result, nil
}

type TplJson struct {
	Length int                      `json:"length"`
	Data   []map[string]interface{} `json:"data"`
}

func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func GetSignal() int {
	if rand.Intn(2) == 0 {
		return 1
	}
	return -1
}
func RandMN(min int, max int) int {
	x := max - min
	n := rand.Intn(x) + min
	return n
}

func GetMysqlParam(n int) string {
	if n < 1 {
		return ""
	}
	str := "?"
	for i := 1; i < n; i++ {
		str += ",?"
	}
	return str
}

func GetIndexByWeigth(data []int) int {
	sum := 0
	index := 0
	for _, v := range data {
		sum += v
	}
	randN := rand.Intn(sum)
	sum = 0
	for i, v := range data {
		sum += v
		if randN < sum {
			index = i
			break
		}
	}
	return index
}

func GetPosInfo() (string, int) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "unknow", 0
	}
	return filepath.Base(file), line
}

func GetPosInfoStr() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "[]"
	}
	return fmt.Sprintf("[%v:%v]\n", filepath.Base(file), line)
}

func FunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
	GBK     = Charset("GBK")
)

func readCommond(ctx context.Context, wg *sync.WaitGroup, std io.ReadCloser, buf *bytes.Buffer) {
	reader := bufio.NewReader(std)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			if runtime.GOOS == "windows" {
				readString = ConvertByte2String([]byte(readString), "GB18030")
			}
			buf.WriteString(readString)
			//fmt.Print(readString)
		}
	}
}

// 执行命令
func ExecCmd(dir string, str string) (string, error) {
	if dir == "" {
		dir = "./"
	}
	var out []byte
	var err error
	cmdSli := strings.Split(str, "|")
	for _, cmdStr := range cmdSli {
		cmdStr = strings.Trim(cmdStr, " ")
		param := strings.Split(cmdStr, " ")
		if len(param) < 1 {
			continue
		}
		cmd := exec.Command(param[0], param[1:]...)
		cmd.Dir = dir
		cmd.Stdin = bytes.NewBuffer(out)
		out, err = cmd.CombinedOutput()
		if err != nil {
			break
		}
	}
	return string(out), err
}

func ExecCmd2(ctx context.Context, cmd string) (string, error) {
	var c *exec.Cmd
	if runtime.GOOS == "windows" {
		c = exec.CommandContext(ctx, "cmd", "/C", cmd) // windows
	} else {
		c = exec.CommandContext(ctx, "bash", "-c", cmd) // mac linux
	}
	stdout, err := c.StdoutPipe()
	if err != nil {
		return "", err
	}
	stderr, err := c.StderrPipe()
	if err != nil {
		return "", err
	}
	var wg sync.WaitGroup
	// 因为有2个任务, 一个需要读取stderr 另一个需要读取stdout
	stdoutBuf := &bytes.Buffer{}
	stderrBuf := &bytes.Buffer{}
	wg.Add(2)
	go readCommond(ctx, &wg, stderr, stderrBuf)
	go readCommond(ctx, &wg, stdout, stdoutBuf)
	// 这里一定要用start,而不是run 详情请看下面的图
	err = c.Start()
	if err != nil {
		return "", err
	}
	// 等待任务结束
	wg.Wait()
	if stderrBuf.Len() > 0 {
		return stdoutBuf.String(), fmt.Errorf("%v", stderrBuf)
	}
	return stdoutBuf.String(), nil
}

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case GBK:
		var decodeBytes, _ = simplifiedchinese.GBK.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
