package main

type Config struct {
	AccessToken string `yaml:"token"`
	Path string `yaml:"path"`
	DirPath string `yaml:"dirpath"`
	Video *ConfigVideo `yaml:"video"`
}
type ConfigVideo struct {
	Build int `yaml:"build"`
	Copyright int8 `yaml:"copyright"`
	Cover string `yaml:"cover"`
	Desc string `yaml:"desc"`
	NoReprint int8 `yaml:"noreprint"`
	OpenElec int8 `yaml:"openelec"`
	Source string `yaml:"source"`
	Tag string `yaml:"tag"`
	Tid int `yaml:"tid"`
	Title string `yaml:"title"`
}
//func main() {
//	file,err :=ioutil.ReadFile("config.yaml")
//	if err != nil {
//		log.Printf("配置文件读取失败:%v",err)
//		return
//	}
//	log.Println("配置文件读取成功!")
//	var a Config
//	_ =yaml.Unmarshal(file,&a)
//	client :=NewBiliBiliVideo(a.AccessToken)
//	client.getUserInfo()
//	if a.Path =="" {
//		files,err :=ioutil.ReadDir(a.DirPath)
//		if err != nil {
//			log.Fatal(err)
//			return
//		}
//		log.Println("目录读取成功")
//		for i, file := range files {
//			Video = append(Video, InitMap(a.DirPath+file.Name(),"p"+strconv.Itoa(i+1),""))
//		}
//	}else {
//		files,err :=os.Open(a.Path)
//		if err != nil {
//			log.Fatal(err)
//			return
//		}
//		log.Println("文件读取成功")
//		Video = append(Video, InitMap(files.Name(),"p1",""))
//		defer files.Close()
//	}
//	log.Println(Video)
//	vi :=client.uploadMain(Video)
//	d :=&VideosStruct{
//		Build:     a.Video.Build,
//		Copyright: a.Video.Copyright,
//		Cover:     a.Video.Cover,
//		Desc:      a.Video.Desc,
//		NoReprint: a.Video.NoReprint,
//		OpenElec:  a.Video.OpenElec,
//		Source:    a.Video.Source,
//		Tag:       a.Video.Tag,
//		Tid:       a.Video.Tid,
//		Title:     a.Video.Title,
//		Videos:    vi,
//	}
//	e :=client.addVideo(d)
//	log.Println(string(e))
//}

