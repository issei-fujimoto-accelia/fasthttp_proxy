# readme

proxyが複数のserverに対してrequestをforward。responseを集約してclientに返す

`go run ./proxy/main.go `

`go run ./server/main.go 8181`

`go run ./server/main.go 8282`


---

proxyは8080で起動する。8080にrequest

```
$ curl localhost:8080

{"person_list":[{"name":"hoge","age":1},{"name":"hoge","age":1}]}
```


- pathとquery stringはそのまま各サーバーにfowardされる
- ndjson形式は対応してないのでTODO
