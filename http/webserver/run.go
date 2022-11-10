package webserver

import (
	"OwlGramServer/consts"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

func (g Context) Run() {
	if consts.IsDebug {
		log.Println("Running Debug WebServer!")
		log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", 8765), g.handler))
	} else {
		log.Println("Running WebServer!")
		log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", 9602), g.handler))
	}
}
