package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

//发送GET请求并返回内容
func sendGetData(durl string) []byte {
	client := &http.Client{}
	var data []byte
	reqest, _ := http.NewRequest("GET", durl, nil)
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0")
	response, err := client.Do(reqest)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		fmt.Println("查询失败,请检查参数及网络设置")
		return data
	}
	data, _ = ioutil.ReadAll(response.Body)
	return data
}

//发送Post请求并返回内容
func sendPostData(durl string, postdata string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("POST", durl, strings.NewReader(postdata))
	if err != nil {
		log.Println("发送失败,请检查参数或网络配置")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("User-Agent", " Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0")
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

//删除a标签
func delAHerf(html string) (newHtml string) {
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
		if strings.Contains(string(html), "</a>") {
			aherfFirst = strings.Index(html, "</a>")
			aherfFirstHtml = html[aherfFirst:]
			html = html[:aherfFirst] + aherfFirstHtml[4:]
		} else {
			break
		}
	}
	newHtml = html
	return newHtml
}

//删除i标签
func delIClass(html string) (newHtml string) {
	var aherfFirst int
	var aherfLast int
	var aherfFirstHtml string
	for { //删<i class=>
		if strings.Contains(string(html), "<i class=") {
			aherfFirst = strings.Index(html, "<i class=")
			aherfFirstHtml = html[aherfFirst:]
			aherfLast = strings.Index(aherfFirstHtml, "</i>")
			html = html[:aherfFirst] + aherfFirstHtml[aherfLast+4:]
		} else {
			break
		}
	}
	// for { //删</i>
	// 	if strings.Contains(string(html), "</i>") {
	// 		aherfFirst = strings.Index(html, "</i>")
	// 		aherfFirstHtml = html[aherfFirst:]
	// 		html = html[:aherfFirst] + aherfFirstHtml[4:]
	// 	} else {
	// 		break
	// 	}
	// }
	newHtml = html
	return newHtml
}

//删除指定字符串
// func delString(str string, delstr string) (newstr string) {
// 	var strFirst int
// 	var strlen int
// 	for {
// 		if strings.Contains(string(str), delstr) {
// 			strFirst = strings.Index(str, delstr)
// 			strlen = len(delstr)
// 			str = str[:strFirst] + str[strFirst+strlen:]
// 		} else {
// 			break
// 		}
// 	}
// 	newstr = str
// 	return newstr
// }

//按指定次数删除指定字符串
func delStringToNum(str string, delstr string, delnum int) (newstr string) {
	var strFirst int
	var strlen int
	for i := 0; i < delnum; i++ {
		if strings.Contains(string(str), delstr) {
			strFirst = strings.Index(str, delstr)
			strlen = len(delstr)
			str = str[:strFirst] + str[strFirst+strlen:]
		} else {
			break
		}
	}
	newstr = str
	return newstr
}

//截取字符串
func deljq(str string, firstStr string, endStr string) (newStr string) {
	var strOne int
	var strTow int
	var strLen int
	var firstNewStr string
	//截掉前端
	if strings.Contains(str, firstStr) {
		strOne = strings.Index(str, firstStr)
		strLen = len(firstStr)
		firstNewStr = str[strOne+strLen:]
	}
	//截掉后端
	if strings.Contains(firstNewStr, endStr) {
		strTow = strings.Index(firstNewStr, endStr)
		newStr = firstNewStr[:strTow]
	}
	return newStr
}

//处理为pingscan为正确json
func regJsonData(Data []byte) []byte {
	exp := "([a-zA-Z]\\w*):"
	reg := regexp.MustCompile(exp)
	regStr := reg.ReplaceAllString(string(Data), `"$1":`)
	//字符串替换值为http中的内容
	newStr := strings.Replace(regStr, `"http":`, "http:", -1)
	return []byte(newStr)
}

//字符串替换,-1表示全部替换, 0表示不替换, 1表示替换第一个, 2表示替换第二个...
func setJsChar(Data []byte, oldStr string, newStr string, setMode int) []byte {
	newJs := strings.Replace(string(Data), oldStr, newStr, setMode)
	return []byte(newJs)
}

//刷新数据
func httpFlusher(w http.ResponseWriter) {
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	} else {
		log.Println("Damn, no flush")
	}
}

//输出指定html标签表格样式
func htmlTag(firstStr string, endStr string, inStr string) string {
	str := firstStr + inStr + endStr
	return str
}
