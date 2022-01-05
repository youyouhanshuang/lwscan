package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bitly/go-simplejson"
)

//Pingscan
type pisHandler struct{}

func (m *pisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	domain := r.PostFormValue("domain")
	log.Println("正在进行CDN/ping检测,目标IP:" + domain + "....")
	durl := "http://ping.chinaz.com/iframe.ashx?t=ping&callback=jQuery1113024287815543324764_1640570758322"
	f, _ := os.Open("wwwroot/dict/pingscandic.txt")
	defer f.Close()
	re := bufio.NewReader(f)
	var byteguid []string
	for {
		aa, err := readLine(re)
		if err != nil {
			break
		}
		guidlist := strings.Split(aa, "---")
		byteguid = append(byteguid, guidlist...)
	}
	httpFlusher(w)
	w.Write([]byte(html1))
	w.Write([]byte(`<title>CDN/Ping检测<title>`))
	w.Write([]byte(html2))
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "目标地址:"+domain)
	w.Write([]byte(`<br />`))
	fmt.Fprintln(w, "Ping列表:")
	w.Write([]byte(`<table class="hovertable">
	<tr>
		<th>监测点</th>
		<th>响应IP</th>
		<th>IP归属地</th>
		<th>响应时间</th>
		<th>TTL</th>
	</tr>
	`))
	//i为guid,i+1为对应地区
	for i := 0; i+1 < len(byteguid); i += 2 {
		postdata := "guid=" + byteguid[i] + "&host=" + domain + "&ishost=0&isipv6=0&encode=TJgQKzHWihwusW2YH6sxqTRcZCsfpgwv&checktype=0"
		body := sendPostData(durl, postdata)
		newbody := []byte(deljq(string(body), "(", ")"))
		jsonOne := setJsChar(newbody, `'`, `"`, -1)
		json := regJsonData(jsonOne)
		js, err := simplejson.NewJson(json)
		if err != nil {
			panic(err.Error())
		}
		//取json中的数据
		Jcd := byteguid[i+1]                           //监测点
		Ip, err := js.Get("result").Get("ip").String() //IP
		if err != nil {
			Ip = "-"
		}
		Ipaddress, err := js.Get("result").Get("ipaddress").String() //IP归属地
		if err != nil {
			Ipaddress = "-"
		}
		Responsetime, err := js.Get("result").Get("responsetime").String() //响应时间
		if err != nil {
			Responsetime = "-"
		}
		Ttl, err := js.Get("result").Get("ttl").String() //TTL
		if err != nil {
			Ttl = "-"
		}
		//输出数据
		Table := "<tr onmouseover=\"this.style.backgroundColor='#ffff66';\" onmouseout=\"this.style.backgroundColor='#d4e3e5';\"><td>" + Jcd + "</td><td>" + Ip + "</td><td>" + Ipaddress + "</td><td>" + Responsetime + "</td><td>" + Ttl + "</td></tr>"
		w.Write([]byte(Table))
		httpFlusher(w)
	}
	w.Write([]byte(`</table></body></html>`))
}
