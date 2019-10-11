package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-emotion/src/utils"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//基本参数结构体
type Basic struct {
	keyword         string   //搜索名称
	StartPage       int      //开始页
	EndPage         int      //结束页
	Path            string   //存放图片路径的地址
	doWorkChan      chan int //爬取页面的管道
	downloadImgChan chan int //下载图片的管道
}

func main() {
	var this Basic
	fmt.Println("请输入关键字:")
	fmt.Scan(&this.keyword)
	fmt.Println("请输入起始页:")
	fmt.Scan(&this.StartPage)
	fmt.Println("请输入结束页:")
	fmt.Scan(&this.EndPage)

	//创建文件夹
	path, err := utils.MkdirFolder(this.keyword)
	if err != nil {
		fmt.Println(err)
	}
	this.Path = path
	//初始化doWorkChan管道
	this.doWorkChan = make(chan int)
	//初始化downloadImgChan管道
	this.downloadImgChan = make(chan int)

	for i := this.StartPage; i <= this.EndPage; i++ {
		go this.doWork(i)
	}

	for i := this.StartPage; i <= this.EndPage; i++ {
		p := <-this.doWorkChan
		fmt.Printf("第 %d 个页面爬取完成\n", p)
	}
	fmt.Println("结束")
}

//爬取图片路径
func (this *Basic) doWork(Page int) {
	//爬取路径
	url := "https://fabiaoqing.com/search/search/keyword/" + this.keyword + "/type/bq/page/" + strconv.Itoa(Page) + ".html"
	//存放图片路径
	var imgSrc []string
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("goquery.NewDocument(url) err = ", err)
	}
	doc.Find(".searchbqppdiv>a>img").Each(func(i int, s *goquery.Selection) {
		result, err := s.Attr("data-original")
		if !err {
			fmt.Println("s.Attr err")
		}
		fmt.Println(result)
		imgSrc = append(imgSrc, result)
	})

	//每页图片数量
	imgLen := len(imgSrc)

	for i := 0; i < imgLen; i++ {
		go this.downloadImg(imgSrc[i], i)
	}
	for i := 0; i < imgLen; i++ {
		<-this.downloadImgChan
	}
	this.doWorkChan <- Page
	return
}

//下载图片
func (this *Basic) downloadImg(src string, i int) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", src, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("Host", "https://fabiaoqing.com")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Cache-Control", "max-age=0")
	request.Header.Add("Upgrade-Insecure-Requests", "1")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	request.Header.Add("Referer", src)
	response, err := client.Do(request)
	defer response.Body.Close()

	//获取文件内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	//获取文件文件名
	imgName := strings.Split(src, "/")[4]
	//图片存放路径
	imgPath := this.Path + "/" + imgName
	//创建空图片文件
	out, err := os.Create(imgPath)
	if err != nil {
		fmt.Println("图片文件创建失败!")
	}
	io.Copy(out, bytes.NewReader(body))
	fmt.Println("下载完成：", src)
	this.downloadImgChan <- i
	return
}
