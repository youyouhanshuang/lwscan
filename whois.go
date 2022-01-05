package main

import (
	"log"
	"net/http"
	"strings"
)

//Whois
type wisHandler struct{}

func (m *wisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	domain := r.PostFormValue("domain")
	log.Println("正在进行whois查询,目标域名:" + domain + "....")
	durl := "http://nwhois.cn/" + domain + "&realtime=1"
	var htmlData string
	var data []byte
	httpFlusher(w)
	for i := 0; i <= 6; i++ {
		data = sendGetData(durl)
		if strings.Contains(string(data), "<span>电子邮件:</span>") {
			htmlData = (string(data))
			break
		}
		if i == 6 {
			w.Write([]byte("查询不到或查询出错,请稍后再试"))
		}
	}
	w.Write([]byte(htmlMode + "<h2>信息列表</h2>"))
	//取列表
	whoisListStart := strings.Index(htmlData, `<div id="list">`)
	ListStart := htmlData[whoisListStart:]
	whoisListEnd := strings.Index(ListStart, `<div id='title'>更多详细信息`)
	ListEnd := ListStart[:whoisListEnd]
	whoisList := delStringToNum(delIClass(delAHerf(ListEnd)), "&#12288;", 2)
	w.Write([]byte(whoisList))
	w.Write([]byte("<h2>详细信息</h2>"))
	//取信息
	infoHtmlData := ListStart[whoisListEnd:]
	whoisInfoStart := strings.Index(infoHtmlData, `<pre>`)
	whoisInfoEnd := strings.Index(infoHtmlData, "最新查询记录")
	Info := infoHtmlData[whoisInfoStart : whoisInfoEnd-63]
	w.Write([]byte("<div>"))
	w.Write([]byte(Info))
	w.Write([]byte("</div>"))
	w.Write([]byte("</body></html>"))
}
