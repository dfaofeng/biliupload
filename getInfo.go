package BiliBiliUploadv2

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

// 分片数量
var (
	Size float64 = 2 * 1024 * 1024
	Video = make([]map[string]string,0)
	v = make([]videos,0)
)
// BiliBiliVideo AccessToken结构体
type BiliBiliVideo struct {
	// token
	AccessToken string `json:"access_token"`
	// cookie
	Cookie *http.Cookie
	// 上传服务器信息
	uploadInfo
	// 上传文件信息
	videos
	// 投稿相关
	VideosStruct
	UserInfo
}
// UserInfo 用户信息结构体
type UserInfo struct {
	Ts int `json:"ts"`
	Code int `json:"code"`
	Data struct {
		Mid int `json:"mid"`
		Appid int `json:"appid"`
		AccessToken string `json:"access_token"`
		ExpiresIn int `json:"expires_in"`
		Userid string `json:"userid"`
		Uname string `json:"uname"`
	} `json:"data"`
}
// NewBiliBiliVideo 构造函数
func NewBiliBiliVideo(token string) *BiliBiliVideo {
	return &BiliBiliVideo{AccessToken: token}
}

// InitMap 稿件信息构造函数
func InitMap(path,title, desc string) map[string]string {
	v :=make(map[string]string)
	v["path"]=path
	v["title"]=title
	v["desc"]=desc
	return v
}
// getUserInfo 获取用户信息方法
func (b *BiliBiliVideo) getUserInfo()  {
	user :="https://passport.bilibili.com/api/oauth2/info?"
	login_params :="access_token="+b.AccessToken+"&appkey=aae92bc66f3edfab&platform=pc&ts="+strconv.Itoa(int(time.Now().Unix()))
	sign :=GetSign([]byte (login_params+"af125a0d5279fd576c1b4418a3e8276d"))
	login_sign :=login_params+"&sign="+sign
	//获取用户信息
	Info,_ :=HttpGet(user+login_sign)
	log.Println(string(Info))
	var ua UserInfo
	_ = json.Unmarshal(Info,&ua)
	b.UserInfo = ua
}
// GetUploadUrl 获取上传信息的结构体和对应的cookie
func (b *BiliBiliVideo) GetUploadUrl()  {
	apiurl :="http://member.bilibili.com/preupload?access_key="+b.AccessToken+"&mid="+strconv.Itoa(b.UserInfo.Data.Mid)+"&profile=ugcfr%2Fpc3"
	//获取上传服务器地址
	d, sid := HttpGet(apiurl)
	log.Println("获取上传服务器信息成功...")
	var upload uploadInfo
	_ = json.Unmarshal(d,&upload)
	b.Cookie = sid
	b.uploadInfo = upload
}
