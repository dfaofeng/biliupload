package biliupload

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var client = resty.New()
// HttpGet 通用get请求
func HttpGet(a string) (c []byte,d *http.Cookie) {
	req,_:=http.NewRequest("GET",a,nil)
	resp,err :=(&http.Client{Timeout: 5*time.Second}).Do(req)
	cookie_sid :=resp.Cookies()
	if len(cookie_sid) != 0 {
		sid :=cookie_sid[0]
		defer resp.Body.Close()
		body,err :=ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		return body,sid
	}

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body,err :=ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body,nil
}

// GetSign md5加密方法
func GetSign(b[]byte) string {
	return fmt.Sprintf("%x",md5.Sum(b))
}
//文件md5
func getFileMd5(a *os.File) (b string) {
	md5Handle:=md5.New()
	_,err := io.Copy(md5Handle,a)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(md5Handle.Sum(nil))
}
func PostData(data io.Reader,apiUrl string,ck map[string]string) []byte {
	client.SetRetryCount(5)
	client.SetRetryWaitTime(time.Second*5)
	resp,err :=client.R().
		SetBody(data).
		SetHeaders(ck).
		Post(apiUrl)
	if err != nil {
		log.Printf("post错误:%v",err)
		return nil
	}
	return resp.Body()
}
