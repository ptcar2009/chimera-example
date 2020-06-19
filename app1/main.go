package main

import (
	"log"

	"github.com/valyala/fasthttp"
	"gitlab.inspr.com/chimera/client-golang"
)

func handle(ctx *fasthttp.RequestCtx) {
	body := ctx.Request.Body()
	err := writer.WriteMessage(string(body), "chan1")
	if err != nil {
		ctx.SetStatusCode(500)
		ctx.SetBody([]byte(err.Error()))
	} else {
		ctx.SetStatusCode(200)

		ctx.SetBody([]byte("REQUEST SENT\n"))
	}
}

var writer client.Writer

func main() {
	var err error
	writer, err = client.NewWriter()
	if err != nil {
		panic(err)
	}
	log.Fatalln(fasthttp.ListenAndServe(":80", handle))
}
