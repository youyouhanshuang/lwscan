package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

var pool = "1234567890abcdefghijklmnopqrstuvwxyz"

//生成随机数字符串
func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}

//发送GET请求无返回内容 cms专用
func sendGetCMS(durl string, cookie string) {
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", durl, nil)
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0")
	reqest.Header.Add("Cookie", ` think_var=zh-cn; PHPSESSID=`+cookie+`; __tins__19980795=%7B%22sid%22%3A%201641109481990%2C%20%22vd%22%3A%202%2C%20%22expires%22%3A%201641111291655%7D; __51cke__=; __51laig__=13`)
	response, err := client.Do(reqest)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
}

//将字符切片以指定字符串相连,并做过滤"和<a>标签处理
func reStringArrayToString(sliceStr []string, useChar string) string {
	var reStr string
	for i := 0; i < len(sliceStr); i++ {
		reStr = reStr + delAHerfCms(delString(sliceStr[i], `"`))
		if i != len(sliceStr)-1 {
			reStr = reStr + useChar
		}
	}
	return reStr
}

//Post cms专用
func sendPostDataCMS(durl string, postdata string, cookie string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("POST", durl, strings.NewReader(postdata))
	if err != nil {
		log.Println("发送失败,请检查参数或网络配置")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("User-Agent", " Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0")
	req.Header.Set("Cookie", ` think_var=zh-cn; PHPSESSID=`+cookie+`; __tins__19980795=%7B%22sid%22%3A%201640664898223%2C%20%22vd%22%3A%201%2C%20%22expires%22%3A%201640666698223%7D; __51cke__=; __51laig__=1`)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("发送失败,请检查参数或网络配置")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("发送失败,请检查参数或网络配置")
	}
	return body
}

//删除指定字符串
func delString(str string, delStr string) string {
	var fNum int
	var newStr string
	var newByte []byte
	var reByte []byte
	for {
		if strings.Contains(str, delStr) {
			fNum = strings.Index(str, delStr)
			if str[fNum+1] == 'u' {
				newStr = str[:fNum+1]
				str = str[fNum+1:]
				newByte = []byte(newStr)
				reByte = append(reByte, newByte...)
				newByte = nil
			} else {
				newStr = str[:fNum]
				str = str[fNum+1:]
				newByte = []byte(newStr)
				reByte = append(reByte, newByte...)
				newStr = ""
				newByte = nil
			}
		} else {
			newByte = []byte(str)
			reByte = append(reByte, newByte...)
			break
		}
	}
	return string(reByte)
}

//删除a标签 及<\/a >,</a >,</a>
func delAHerfCms(html string) (newHtml string) {
	var aherfFirst int
	var aherfLast int
	var aherfFirstHtml string
	for { //删<a herf>
		if strings.Contains(string(html), "<a href=") {
			aherfFirst = strings.Index(html, "<a href=")
			aherfFirstHtml = html[aherfFirst:]
			aherfLast = strings.Index(aherfFirstHtml, ">")
			html = html[:aherfFirst] + aherfFirstHtml[aherfLast+1:]
		} else {
			break
		}
	}
	for { //删</a>
		if strings.Contains(string(html), "<\\/a >") {
			aherfFirst = strings.Index(html, "<\\/a >")
			aherfFirstHtml = html[aherfFirst:]
			html = html[:aherfFirst] + aherfFirstHtml[6:]
		} else if strings.Contains(string(html), "</a >") {
			aherfFirst = strings.Index(html, "</a >")
			aherfFirstHtml = html[aherfFirst:]
			html = html[:aherfFirst] + aherfFirstHtml[5:]
		} else if strings.Contains(string(html), "</a>") {
			aherfFirst = strings.Index(html, "</a >")
			aherfFirstHtml = html[aherfFirst:]
			html = html[:aherfFirst] + aherfFirstHtml[4:]
		} else {
			break
		}
	}
	newHtml = html
	return newHtml
}
