package biliupload

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Videos 稿件信息结构体
type videos struct {
	Desc string `json:"desc"`
	Filename string `json:"filename"`
	Title string `json:"title"`
}
// VideosStruct 投稿信息结构体
type VideosStruct struct {
	//build参数,默认为1054,代表投稿设备类型
	Build int `json:"build"`
	//投稿类型,1为自制,2为转载
	Copyright int8 `json:"copyright"`
	//投稿封面,为空时,默认是视频第一帧
	Cover string `json:"cover"`
	//稿件简介
	Desc string `json:"desc"`
	//视频是否禁止转载标志0无1禁止
	NoReprint int8 `json:"no_reprint"`
	//是否开启充电面板，0为关闭1为开启
	OpenElec int8 `json:"open_elec"`
	//转载投稿需要,转载来源
	Source string `json:"source"`
	//tag标签
	Tag string `json:"tag"`
	//分区id,详细请参考
	//https://github.com/FortuneDayssss/BilibiliUploader/wiki/Bilibili%E5%88%86%E5%8C%BA%E5%88%97%E8%A1%A8
	Tid int `json:"tid"`
	//稿件主标题
	Title string `json:"title"`
	//上传结构体
	Videos *[]videos `json:"videos"`
}
// AddVideo 投稿方法
func (b *BiliBiliVideo) AddVideo(s *VideosStruct) []byte {
	//json反序列化
	post_video,_:=json.Marshal(s)
	addvideoApi := "http://member.bilibili.com/x/vu/client/add"
	sign := GetSign([]byte("access_key="+b.AccessToken+"af125a0d5279fd576c1b4418a3e8276d"))
	videoUrl :=addvideoApi+"?access_key="+b.AccessToken+"&sign="+sign
	log.Println(string(post_video))
	req,err :=http.NewRequest("POST",videoUrl,bytes.NewBuffer(post_video))
	if err != nil {
		log.Printf("%v",err)
		return nil
	}
	req.Header.Set("Connection","keep-alive")
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("User-Agent","")
	client :=&http.Client{}
	resp,err :=client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body,_:=ioutil.ReadAll(resp.Body)
	return body
}
