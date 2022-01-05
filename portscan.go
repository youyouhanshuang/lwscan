package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/thedevsaddam/gojsonq"
	"golang.org/x/net/websocket"
)

//portscan's Handler
type ptsHandler struct{}

func (m *ptsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var origin = "http://coolaf.com"
	var url = "ws://coolaf.com:9010/tool/ajaxport"
	Ptsdomain := r.PostFormValue("ptsdomain")
	Ptsport := r.PostFormValue("ptsport")
	log.Println("正在扫描\"" + Ptsdomain + "\"的端口,请稍等...")
	w.Write([]byte(html1))
	w.Write([]byte(`<title>端口扫描<title>`))
	w.Write([]byte(html2))
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "目标地址:"+Ptsdomain)
	w.Write([]byte(`<br />`))
	fmt.Fprintln(w, "端口列表:")
	w.Write([]byte(`<table class="hovertable">
	<tr>
		<th>目标域名</th>
		<th>IP</th>
		<th>端口号</th>
		<th>状态</th>
		<th>可能运行服务</th>
	</tr>
	`))
	ws, err := websocket.Dial(url, "", origin) //websocket拨号
	if err != nil {
		log.Fatal(err)
	}
	rePortList := readPortString(Ptsport) //解析需要爆破的端口
	httpFlusher(w)
	f, _ := os.Open("wwwroot/dict/servicedic.txt")
	defer f.Close()
	re := bufio.NewReader(f)
	var byteService []string
	for {
		services, err := readLine(re)
		if err != nil {
			break
		}
		sStringlist := strings.Split(services, "    ") //以","切割,并存入数组
		byteService = append(byteService, sStringlist...)
	}
	for i := 0; i < len(rePortList); i++ {
		Port := strconv.Itoa(rePortList[i])
		message := []byte("{\"ip\":\"" + Ptsdomain + "\",\"port\":\"" + Port + "\"}")
		_, err = ws.Write(message)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("Send: %s\n", message)

		var msg = make([]byte, 512)
		m1, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("Receive: %s\n", msg[:m1])
		jsonX := string((msg[:m1]))
		psStatus := gojsonq.New().JSONString(jsonX).Find("Status")
		strStatus := fmt.Sprint(psStatus)
		if strStatus == "2" {
			continue
		}
		psHost := gojsonq.New().JSONString(jsonX).Find("Host")
		psIp := gojsonq.New().JSONString(jsonX).Find("Ip")
		psPort := gojsonq.New().JSONString(jsonX).Find("Port")
		strHost := fmt.Sprint(psHost)
		strIp := fmt.Sprint(psIp)
		strPort := fmt.Sprint(psPort)
		strServer := takeServiceString(strPort, byteService)
		Table := "<tr onmouseover=\"this.style.backgroundColor='#ffff66';\" onmouseout=\"this.style.backgroundColor='#d4e3e5';\"><td>" + strHost + "</td><td>" + strIp + "</td><td>" + strPort + "</td><td>开启</td><td>" + strServer + "</td></tr>"
		w.Write([]byte(Table))
		//fmt.Fprintln(w, psHost, str, psIp, str, psPort, str, psStatus)
		// fmt.Fprint(w, psHost)
		// fmt.Fprint(w, "---")
		// fmt.Fprint(w, psIp)
		// fmt.Fprint(w, "---")
		// fmt.Fprint(w, psPort)
		// fmt.Fprint(w, "---")
		// fmt.Fprintln(w, "开启")
		httpFlusher(w)
	}
	w.Write([]byte(`</table></body></html>`))

}
