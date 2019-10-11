# go-emoticon

#### 介绍
golang高性能爬虫,爬取表情包网站方便和小伙伴斗图

#### 软件架构
* golang版本1.13.1
* 使用`github.com/PuerkitoBio/goquery`爬虫库 文档地址 https://www.itlipeng.cn/2017/04/25/goquery-%E6%96%87%E6%A1%A3/
* 使用多协程优化爬取速度
* 爬取地址 https://fabiaoqing.com
* 爬取速度取决于网速和io
#### 安装教程

* 安装`goquery`包

#### 使用说明

* go run main.go 或者 编译后运行
* 根据提示输入要爬取的内容和页数
* 等待爬取完成
* 然后就可以愉快的去和小伙伴斗图啦～
