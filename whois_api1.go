package main

import (
	"log"
	"net/http"
	"strings"
)

//Whois_api1
type wis1Handler struct{}

func (m *wis1Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	domain := r.PostFormValue("domain")
	log.Println("正在使用备用whois查询,目标域名:" + domain + "....")
	//domain := `baidu.com`
	durl := "https://site.ip138.com/" + domain + "/whois.htm"
	data := sendGetData(durl)
	htmlData := (string(data))
	whoisStart := strings.Index(htmlData, `<div class="whois" id="whois">`)
	First := htmlData[whoisStart:]
	httpFlusher(w)
	if strings.Contains(First, "whois服务器远程获取超时") {
		whoisEnd := strings.Index(First, `<p><a class="btn"`)
		Last := First[:whoisEnd]
		w.Write([]byte(Last))
		w.Write([]byte("</div>"))
	} else {
		whoisEnd := strings.Index(First, `<a href="javascript:;" rel="nofollow" class="update" id="update"`)
		Last := First[:whoisEnd]
		w.Write([]byte(Last))
	}
}
