package main

import (
	"fmt"
	"log"
	"encoding/json"
	"sync"
	
	"github.com/valyala/fasthttp"
)


type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PersonList struct {
	Person []*Person    `json:"person_list"`
}

func doGet(url string, wg *sync.WaitGroup) (*Person, error) {
	defer wg.Done()

	client := &fasthttp.Client{}
	req := fasthttp.AcquireRequest() //空のrequestを作る
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)
	
	resp := fasthttp.AcquireResponse() //空のresponseを作る
	err := client.Do(req, resp) // request	
	fasthttp.ReleaseRequest(req)

	var p Person
	err = json.Unmarshal(resp.Body(), &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}


func Handler(ctx *fasthttp.RequestCtx) {
	serverList := [2]string{
		"localhost:8181",
		"localhost:8282",
	}
	
	var wg sync.WaitGroup
	mutex := sync.Mutex{}

	path := ctx.URI().Path()
	query := ctx.URI().QueryString()
	
	fmt.Println("path: ", string(path))
	fmt.Println("query: ", string(query))

	personList := PersonList{}
	for _, host := range serverList {
		fmt.Println("host", host)
		wg.Add(1)
		go func(h string) {
			uri := fasthttp.AcquireURI()
			uri.SetPathBytes(path)
			uri.SetQueryStringBytes(query)
			uri.SetHost(h)
			
			fmt.Println("url", uri.String())
			r, err := doGet(uri.String(), &wg)
			if err != nil {
				fmt.Printf("debug print %s\n", err)
				return
			}
			mutex.Lock()
			personList.Person = append(personList.Person, r)
			mutex.Unlock()			
		}(host)
	}
	wg.Wait()


	ctx.Response.Header.Set("Content-Type", "application/json")
	s, _ := json.Marshal(personList)

	ctx.SetBody(s)
	fmt.Println("return!")
	return
}


func main() {
	log.Fatal(fasthttp.ListenAndServe("127.0.0.1:8080", Handler))
}
