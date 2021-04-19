package pool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gitee.com/clearluo/gotools/log"
	"gitee.com/clearluo/gotools/util"
)

func SendHttpMyClient(bodyStr string, remoteIP string) error {
	//bufLog := bytes.Buffer{}
	bufLog := NewMyBytesBuff()
	defer bufLog.Free()
	httpClient := NewHttpClient()
	defer httpClient.Free()
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("SendHttpMyClient panic:[err:%v req:%v ip:%v]", err, bodyStr, remoteIP)
		} else {
			log.Debugf("SendHttpMyClient[%v]bufLog:\n%v", remoteIP, bufLog.String())
		}
		bufLog.Reset()
	}()
	eventBody := map[string]interface{}{}
	reqBodyByte, _ := json.Marshal(eventBody)
	bufLog.WriteString("httpReqBody[client]:" + string(reqBodyByte) + util.GetPosInfoStr())
	// 3.发送
	//client := &http.Client{Timeout: time.Second * 3} // 改连接池
	bodyReq := strings.NewReader(string(reqBodyByte))
	req, err := http.NewRequest("POST", "url", bodyReq)
	if err != nil {
		bufLog.WriteString(err.Error() + util.GetPosInfoStr())
		return err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	resp, err := httpClient.Do(req)
	if err != nil {
		bufLog.WriteString(err.Error() + util.GetPosInfoStr())
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		bufLog.WriteString(err.Error() + util.GetPosInfoStr())
		return err
	}
	bufLog.WriteString(fmt.Sprintf("httpResp: %v", string(body)) + util.GetPosInfoStr())
	return nil
}
