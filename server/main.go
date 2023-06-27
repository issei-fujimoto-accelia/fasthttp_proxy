package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
var hoge = ""

func Handler(ctx *fasthttp.RequestCtx) {
	// fmt.Println("query: ", string(ctx.QueryArgs().Peek("test")))
	fmt.Println("path: ", string(ctx.URI().Path()))
	fmt.Println("query: ", string(ctx.URI().QueryString()))


	ctx.Response.Header.Set("Content-Type", "application/json")


	p := Person{
		Name: hoge,
		Age: 1,
	}
	s, _ := json.Marshal(p)
	ctx.SetBody(s)
	fmt.Println("return!")
	return
}



func main() {
	port := os.Args[1]
	addr := fmt.Sprintf("127.0.0.1:%s", port)
	fmt.Println("listen... ", addr)
	hoge = addr
	
	r := router.New()
	r.GET("/", Handler)
	r.GET("/test", Handler)
	log.Fatal(fasthttp.ListenAndServe(addr, r.Handler))
}
