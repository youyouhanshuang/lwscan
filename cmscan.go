package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/bitly/go-simplejson"
)

//CMS查询
type cmsHandler struct{}

func (m *cmsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	cookie := randomString(26) //生成26位随机数
	sendGetCMS(`http://finger.tidesec.com/uc/center/api/?time=1641109499&code=c60dnh56fklV3wWeRC%2FL3z8cxaqg00J4sROixIP99El5pL3u%2BY76PmG2m6xk9UaGuaIIL4I16Nx3rhBmxAlM%2FE7WaK8CECwvOc1pt%2BnjjQFqKus13wufF1%2BV2MM5i0%2F4ogz0w%2Frg1idBK4UJVLj9XNTyxnFn5cFgSdwCqFftsKdfkk8`, cookie)
	durl := "http://finger.tidesec.com/home/index/reget"
	postData := "target=http://www.lishouhong.com"
	res := sendPostDataCMS(durl, postData, cookie)
	js, err := simplejson.NewJson(res) //使用simplejson处理json
	if err != nil {
		panic(err.Error())
	}
	Msg, _ := js.Get("msg").String()
	fmt.Println(string(Msg))
	Finger, err := js.Get("res").Get("finger").String() //取出finger的json
	if err != nil {
		panic(err.Error())
	}
	Banner, err := js.Get("res").Get("banner").String() //取出banner的json
	//fmt.Println(Banner)
	if err != nil {
		panic(err.Error())
	}
	FingerJson := []byte(Finger) //转为切片
	BannerJson := []byte(Banner)
	fingerJson, err := simplejson.NewJson(FingerJson) //此为处理后的finger的json
	if err != nil {
		panic(err.Error())
	}
	bannerJson, err := simplejson.NewJson(BannerJson) //此为处理后的banner的json
	if err != nil {
		panic(err.Error())
	}
	//开始输出cms数据
	w.Write([]byte(html1))
	w.Write([]byte("CMS查询"))
	w.Write([]byte(html2))
	w.Write([]byte("<table>"))
	httpFlusher(w)
	//url
	Url, _ := js.Get("res").Get("url").String()
	w.Write([]byte(htmlTag("<tr>", "</tr>", htmlTag("<td>", "</td>", "Url地址")+htmlTag("<td>", "</td>", Url))))
	//网站标题
	Title, _ := fingerJson.Get("title").String()
	w.Write([]byte(htmlTag("<tr>", "</tr>", htmlTag("<td>", "</td>", "网站标题")+htmlTag("<td>", "</td>", Title))))
	//访问状态码
	State, _ := fingerJson.Get("state").String()
	w.Write([]byte(htmlTag("<tr>", "</tr>", htmlTag("<td>", "</td>", "状态码")+htmlTag("<td>", "</td>", State))))
	//中间件
	HttpServer, _ := fingerJson.Get("httpserver").String()
	w.Write([]byte(htmlTag("<tr>", "</tr>", htmlTag("<td>", "</td>", "中间件")+htmlTag("<td>", "</td>", HttpServer))))
	//CMS信息
	Cms, _ := js.Get("res").Get("cms").String()
	cmslist := deljq(delAHerfCms(delString(Cms, `"`)), "[", "]")
	w.Write([]byte(htmlTag("<tr>", "</tr>", htmlTag("<td>", "</td>", "CMS信息")+htmlTag("<td>", "</td>", cmslist))))
	//banner块
	w.Write([]byte("<tr>"))
	w.Write([]byte(htmlTag("<td>", "</td>", "banner")))
	w.Write([]byte("<td>"))
	for k := range bannerJson.MustMap() {
		i := bannerJson.Get(k).MustStringArray()
		ii := reStringArrayToString(i, ",")
		w.Write([]byte(ii + "<br />"))
	}
	w.Write([]byte("</td>"))
	w.Write([]byte("</tr>"))
	//OS
	Os, _ := js.Get("res").Get("os").String()
	w.Write([]byte(htmlTag("<tr>", "</tr>", htmlTag("<td>", "</td>", "OS")+htmlTag("<td>", "</td>", Os))))
	//WAF
	Waf, _ := js.Get("res").Get("waf").String()
	w.Write([]byte(htmlTag("<tr>", "</tr>", htmlTag("<td>", "</td>", "WAF")+htmlTag("<td>", "</td>", Waf))))
	//IP地址块
	httpFlusher(w)
}
