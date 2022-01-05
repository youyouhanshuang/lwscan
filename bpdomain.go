package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/thedevsaddam/gojsonq"
)

// type Redomain struct {
// 	Status string
// 	Domain string
// 	Ip     string
// }

//子域名爆破
type bpdHandler struct{}

func (m *bpdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	setdm := r.PostFormValue("domain")
	log.Println("正在爆破\"" + setdm + "\"的子域名,请稍等...")
	w.Write([]byte(html1))
	w.Write([]byte(`<title>子域名爆破<title>`))
	w.Write([]byte(html2))
	fmt.Fprintln(w, "")
	//r.ParseForm()
	httpFlusher(w)
	fmt.Fprintln(w, "目标域名:"+r.PostFormValue("domain"))
	w.Write([]byte(`<br>`))
	fmt.Fprintln(w, "域名列表:")
	//fmt.Fprintln(w, "状态码--------域名-----------IP地址--------")
	w.Write([]byte(`<table class="hovertable">
	<tr>
		<th>状态码</th>
		<th>域名</th>
		<th>IP地址</th>
	</tr>
	`))

	//发送GET
	durl, _ := url.Parse("https://phpinfo.me/domain")
	q := durl.Query()
	q.Set("domain", setdm)
	f, _ := os.Open("wwwroot/dict/dic.txt") //将目录所读取的设置为参数q的value
	defer f.Close()
	re := bufio.NewReader(f)
	for {
		aa, err := readLine(re)
		if err != nil {
			break
		}
		//fmt.Println(string(aa))
		q.Set("q", aa)
		durl.RawQuery = q.Encode()
		res, err := http.Get(durl.String())
		if err != nil {
			log.Fatal(err)
			return
		}
		result, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
		if strings.Contains(string(result), "status\":200,") { //对返回的状态码进行判断
			jsonX := string(result)
			domain := gojsonq.New().JSONString(jsonX).Find("domain")
			status := gojsonq.New().JSONString(jsonX).Find("status")
			ip := gojsonq.New().JSONString(jsonX).Find("ip")
			//str := "---"
			Domain := fmt.Sprint(domain) //interface转string
			Status := fmt.Sprint(status)
			Ip := fmt.Sprint(ip)
			Table := "<tr onmouseover=\"this.style.backgroundColor='#ffff66';\" onmouseout=\"this.style.backgroundColor='#d4e3e5';\"><td>" + Status + "</td><td>" + Domain + "</td><td>" + Ip + "</td></tr>"
			w.Write([]byte(Table))
			httpFlusher(w)
		}

	}
	w.Write([]byte(`</table></body></html>`))
}
