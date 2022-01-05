package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/bitly/go-simplejson"
)

//NsLookup
type nslHandler struct{}

func (m *nslHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	setdm := r.PostFormValue("domain")
	log.Println("正在对\"" + setdm + "\"进行NsLookUp查询,请稍等...")
	w.Write([]byte(html1))
	w.Write([]byte(`<title>NSLookUp<title>`))
	w.Write([]byte(html2))
	fmt.Fprintln(w, "")
	//r.ParseForm()
	fmt.Fprintln(w, "目标域名:"+r.PostFormValue("domain"))
	w.Write([]byte(`<br>`))
	fmt.Fprintln(w, "NsLookUp:")
	w.Write([]byte(`<table class="hovertable">
	<tr>
		<td>记录</td>
		<td>内容</td>
	</tr>
	`))
	data := make(url.Values)
	data["ip"] = []string{setdm}
	nslres, err := http.PostForm("http://coolaf.com/tool/ajaxnslook", data)
	nslres.Header.Add("Content-Type", "application/x-www-form-urlencoded;")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer nslres.Body.Close()
	body, _ := ioutil.ReadAll(nslres.Body)
	httpFlusher(w)
	js, err := simplejson.NewJson(body) //使用simplejson处理json
	if err != nil {
		panic(err.Error())
	}
	//输出A记录
	rowa, _ := js.Get("data").Get("a").Array() //取A记录
	if rowa != nil {
		//Rowa := fmt.Sprint(rowa)
		TableA1 := "<tr onmouseover=\"this.style.backgroundColor='#ffff66';\" oAnmouseout=\"this.style.TbackgroundColor='#d4e3e5';\"><td>A记录</td><td>"
		TableA2 := "</td></tr>"
		w.Write([]byte(TableA1))
		for z := 0; z < len(rowa); z++ {
			w.Write([]byte(fmt.Sprint(rowa[z]) + "<br>"))
		}
		w.Write([]byte(TableA2))
	}
	//输出CNAME
	rowc, _ := js.Get("data").Get("c").String() //取CNAME
	TableC1 := "<tr onmouseover=\"this.style.backgroundColor='#ffff66';\" oAnmouseout=\"this.style.TbackgroundColor='#d4e3e5';\"><td>CNAME</td><td>"
	TableC2 := "</td></tr>"
	w.Write([]byte(TableC1))
	w.Write([]byte(rowc + "<br>"))
	w.Write([]byte(TableC2))

	//输出MX
	rowmx, _ := js.Get("data").Get("mx").Array() //取MX
	if rowmx != nil {
		//Rowa := fmt.Sprint(rowa)
		TableMX1 := "<tr onmouseover=\"this.style.backgroundColor='#ffff66';\" oAnmouseout=\"this.style.TbackgroundColor='#d4e3e5';\"><td>MX</td><td>"
		TableMX2 := "</td></tr>"
		w.Write([]byte(TableMX1))
		for _, mxrow := range rowmx { //遍历MX数组
			if each_map, ok := mxrow.(map[string]interface{}); ok {
				//取host
				if reHost, ok := each_map["Host"].(string); ok {
					w.Write([]byte((reHost) + "   "))
				}
				//取Pref
				if rePref, ok := each_map["Pref"].(json.Number); ok {
					w.Write([]byte("Pref:" + (rePref) + "<br>"))
				}
			}
		}
		w.Write([]byte(TableMX2))
	}
	//输出NS
	rowns, _ := js.Get("data").Get("ns").Array() //取NS
	if rowns != nil {
		//Rowa := fmt.Sprint(rowa)
		TableNs1 := "<tr onmouseover=\"this.style.backgroundColor='#ffff66';\" oAnmouseout=\"this.style.TbackgroundColor='#d4e3e5';\"><td>NS</td><td>"
		TableNs2 := "</td></tr>"
		w.Write([]byte(TableNs1))
		for _, nsrow := range rowns { //遍历NS数组
			if each_map, ok := nsrow.(map[string]interface{}); ok {
				//取Host
				if reHost, ok := each_map["Host"].(string); ok {
					w.Write([]byte((reHost) + "<br>"))
				}
			}
		}
		w.Write([]byte(TableNs2))
	}

	//输出TXT
	rowtxt, _ := js.Get("data").Get("txt").Array() //取TXT
	if rowtxt != nil {
		TableT1 := "<tr onmouseover=\"this.style.backgroundColor='#ffff66';\" oAnmouseout=\"this.style.TbackgroundColor='#d4e3e5';\"><td>TXT</td><td>"
		TableT2 := "</td></tr>"
		w.Write([]byte(TableT1))
		for z := 0; z < len(rowtxt); z++ {
			w.Write([]byte(fmt.Sprint(rowtxt[z]) + "<br>"))
		}
		w.Write([]byte(TableT2))
	}
	httpFlusher(w)
	w.Write([]byte(`</table></body></html>`))
}
