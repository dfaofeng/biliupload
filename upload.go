package biliupload

import (
	"bytes"
	"github.com/cheggaaa/pb/v3"
	"io"
	"io/ioutil"
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
)
// UploadInfo 返回上传服务器结构体
type uploadInfo struct {
	Complete string `json:"complete"`
	Filename string `json:"filename"`
	OK int `json:"OK"`
	URL string `json:"url"`
}
//主上传方法
func (b *BiliBiliVideo) uploadMain(videoList []map[string]string) *[]videos{
	//遍历传递进来的视频列表
	for i, m := range videoList {
		log.Println(i, m["path"])
		//请求上传地址
		b.GetUploadUrl()
		file, err := os.Open(m["path"])
		if err != nil {
			log.Printf("文件打开失败:%v", err)
			os.Exit(1)
			return nil
		}
		// 最后关闭文件
		defer file.Close()
		fi, _ := file.Stat()
		//读文件大小
		file_size := fi.Size()
		//转为float64获取上传块数量
		file_flo := float64(file_size)
		chunk_totalnum := int(math.Ceil(float64(file_flo)/Size))
		log.Printf("文件大小:%d",file_size)
		log.Printf("切片数量:%d", chunk_totalnum)
		var tmp = make([]byte, int(Size))
		//开始从第一片切
		bar :=pb.StartNew(chunk_totalnum)
		for i := 1; i <= chunk_totalnum; i++ {
			bar.Increment()
			n, err := file.Read(tmp)
			//切完跳出循环
			if err == io.EOF {
				log.Println("切片读取错误,请重试")
				return nil
			}
			//传入切片上传
			b.uploadPart(i, chunk_totalnum, tmp[:n], fi.Name())
			//log.Printf("分片%d上传成功!:总共%d块", i,chunk_totalnum)
		}
		bar.Finish()
		// 通知服务端已全部上传完毕
		data := make(url.Values)
		data["chunks"] = []string{strconv.Itoa(chunk_totalnum)}
		data["filesize"] = []string{strconv.Itoa(int(file_size))}
		data["md5"] = []string{getFileMd5(file)}
		data["name"] = []string{file.Name()}
		data["version"] = []string{"2.0.0.1054"}
		resp, err := http.PostForm(b.Complete, data)
		content, _ := ioutil.ReadAll(resp.Body)
		log.Println("保存视频返回数据:",string(content))
		resp.Body.Close()
		a :=videos{
			Desc:     m["desc"],
			Filename: b.uploadInfo.Filename,
			Title:    m["title"],
		}
		v = append(v, a)
	}
	return &v
}
// uploadPart 上传分p方法
func (b *BiliBiliVideo) uploadPart(i,chunk int,tmp []byte,name string)  {
	// 头部预处理
	body :=bytes.Buffer{}
	writer :=multipart.NewWriter(&body)
	_ = writer.WriteField("version","2.0.0.1054")
	_ = writer.WriteField("filesize", strconv.Itoa(int(Size)))
	_ = writer.WriteField("chunk", strconv.Itoa(i))
	_ = writer.WriteField("chunks", strconv.Itoa(chunk))
	_ = writer.WriteField("md5",GetSign(tmp))
	upload1Writer,_:=writer.CreateFormFile("file",name)
	upload1Writer.Write(tmp)
	_ = writer.Close()
	//post发给服务器
	ck := "PHPSESSID="+b.uploadInfo.Filename
	cook :=make(map[string]string)
	cook["Cookie"] = ck
	cook["Content-Type"] = writer.FormDataContentType()
	_ =PostData(&body,b.uploadInfo.URL,cook)
}