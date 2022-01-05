package main

import (
	"log"
	"net/http"
)

func main() {
	bpd := bpdHandler{}   //bpdomain's handler
	pts := ptsHandler{}   //portscan's handler
	nsl := nslHandler{}   //nsLookup's handler
	wis := wisHandler{}   //whois's handler
	bas := basHandler{}   //basearch's handler
	pis := pisHandler{}   //pingscan's handeler
	cms := cmsHandler{}   //cmsscan's handeler
	cms1 := cms1Handler{} //cmsscan's handeler
	wis1 := wis1Handler{} //whois_api1's handler
	server := http.Server{
		Addr: ":1433",
		//Handler: nil, // DefaultServeMux
	}
	http.Handle("/bpdomain", &bpd)
	http.Handle("/portscan", &pts)
	http.Handle("/nslookup", &nsl)
	http.Handle("/whois", &wis)
	http.Handle("/whois1", &wis1)
	http.Handle("/basearch", &bas)
	http.Handle("/pingscan", &pis)
	http.Handle("/cmsscan", &cms)
	http.Handle("/cms_scan", &cms1)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "wwwroot"+r.URL.Path)
	})
	log.Println(" ██╗     ██╗    ██╗███████╗ ██████╗ █████╗ ███╗   ██╗")
	log.Println(" ██║     ██║    ██║██╔════╝██╔════╝██╔══██╗████╗  ██║")
	log.Println(" ██║     ██║███╗██║╚════██║██║     ██╔══██║██║╚██╗██║")
	log.Println(" ███████╗╚███╔███╔╝███████║╚██████╗██║  ██║██║ ╚████║")
	log.Println(" ╚══════╝ ╚══╝╚══╝ ╚══════╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═══╝")
	log.Println("")
	log.Println("")
	log.Println("")
	log.Println(" Welcome to LWScan!")
	log.Println(" Starting HTTP server on port 1433....")
	log.Println(" Please visit the website on \"http://localhost:1433\"....")
	log.Println("                                              ------Support for 风御安全_悠悠寒霜儿")
	server.ListenAndServe()
}
