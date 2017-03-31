package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	c := &fasthttp.Client{}

	reqHeader := fasthttp.RequestHeader{}
	reqHeader.SetMethod("GET")
	reqHeader.SetRequestURI("http://127.0.0.1:1323/")
	reqHeader.Set("Content-Type", "application/json")

	req := fasthttp.Request{
		Header: reqHeader,
	}

	resp := new(fasthttp.Response)

	err := c.Do(&req, resp)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(string(resp.Body()[:]))
}
