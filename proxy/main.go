package main

import (
	"fmt"
	"log"
	"encoding/json"
	"sync"
	
	"github.com/fasthttp/router"
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
		"http://localhost:8181/",
		"http://localhost:8282/",
	}
	var wg sync.WaitGroup
	mutex := sync.Mutex{}

	personList := PersonList{}
	for _, url := range serverList {
		wg.Add(1)
		go func() {
			r, err := doGet(url, &wg)
			if err != nil {
				fmt.Printf("debug print %s\n", err)
			}			
			mutex.Lock()
			personList.Person = append(personList.Person, r)
			mutex.Unlock()			
		}()
	}
	wg.Wait()


	ctx.Response.Header.Set("Content-Type", "application/json")
	// fmt.Println(personList)
	s, _ := json.Marshal(personList)	
	// fmt.Println(string(s))

	ctx.SetBody(s)
	fmt.Println("return!")
	return
}

func main() {
    r := router.New()
    r.GET("/", Handler)
    log.Fatal(fasthttp.ListenAndServe("127.0.0.1:8080", r.Handler))
}
