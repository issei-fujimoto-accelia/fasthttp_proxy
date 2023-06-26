# readme

proxyが複数のserverに対してrequestをforwardresponseを集約してclientに返す

`go run ./proxy/main.go `
`go run ./server/main.go 8181`
`go run ./server/main.go 8282`


```
$ curl localhost:8080
{"person_list":[{"name":"hoge","age":1},{"name":"hoge","age":1}]}
```
