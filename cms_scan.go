package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
)

//CMS查询
type cms1Handler struct{}

func (m *cms1Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	durl := "http://whatweb.bugscaner.com/what.go"
	setdm := r.PostFormValue("domain")
	postData := "url=" + setdm
	res := sendPostData(durl, postData)
	js, err := simplejson.NewJson(res) //使用simplejson处理json
	if err != nil {
		panic(err.Error())
	}
	log.Println("正在检测\"" + setdm + "\"的cms信息,请稍等...")
	httpFlusher(w)
	w.Write([]byte(html1))
	w.Write([]byte("CMS查询"))
	w.Write([]byte(html2))
	w.Write([]byte("<table>"))
	// status, _ := js.Get("status_code").Int() //状态码
	// fmt.Println(status)
	// url, _ := js.Get("url").String() //请求url
	//fmt.Println(url)
	for key := range js.MustMap() {
		sta, _ := js.Get("status").Int()
		if sta == 10086 {
			w.Write([]byte("查询频率过高,请稍后再试"))
			break
		}
		if key == "status" {
			continue
		}
		i := js.Get(key).MustStringArray()
		if i == nil {
			j, _ := js.Get(key).String()
			if j == "" {
				l, _ := js.Get(key).Int()
				ll := fmt.Sprint(l)
				w.Write([]byte("<tr><th>" + key + "</th><th>" + ll + "</th></tr><br />"))
				httpFlusher(w)
				continue
			}
			w.Write([]byte("<tr><th>" + key + "</th><th>" + j + "</th></tr><br />"))
			httpFlusher(w)
			continue
		}
		ii := reStringArrayToString(i, ",")
		w.Write([]byte("<tr><th>" + key + "</th><th>" + ii + "</th></tr><br />"))
		httpFlusher(w)
	}
	w.Write([]byte("</table></body></html>"))
}
