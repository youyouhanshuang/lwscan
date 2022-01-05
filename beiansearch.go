package main

import (
	"net/http"
	"net/url"
	"strings"
)

//备案查询
type basHandler struct{}

func (m *basHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(htmlMode + "<h4>备案信息</h4>"))
	domain := url.QueryEscape(r.PostFormValue("domain"))
	durl := "https://icplishi.com/" + domain + "/"
	httpFlusher(w)
	data := sendGetData(durl)
	if data == nil {
		w.Write([]byte("查询失败,请检查参数及网络设置"))
		return
	}
	htmlData := (string(data))
	//取第一个表格
	firstTableStart := strings.Index(htmlData, `<table>`)
	fTable := htmlData[firstTableStart:]
	firstTableEnd := strings.Index(fTable, `</table>`)
	fTableEnd := delAHerf(fTable[:firstTableEnd+8])
	w.Write([]byte("<div>"))
	w.Write([]byte(fTableEnd))
	w.Write([]byte("</div>"))
	//取第二个表格(若存在)
	secondHtmlData := fTable[firstTableEnd+8:]
	if strings.Contains(secondHtmlData, "<table>") {
		secondTableStart := strings.Index(secondHtmlData, `<table>`)
		sTable := secondHtmlData[secondTableStart:]
		secondTableEnd := strings.Index(sTable, `</table>`)
		sTableEnd := delAHerf(sTable[:secondTableEnd+8])
		w.Write([]byte("<h4>历史备案信息</h4>"))
		w.Write([]byte("<div>"))
		w.Write([]byte(sTableEnd))
		w.Write([]byte("</div>"))
	}
	w.Write([]byte("</body></html>"))
}
