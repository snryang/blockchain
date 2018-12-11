package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Params struct {
	Width, Height int
}

func main() {
	//链接远程rpc服务 HTTP
	//rpc, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")

	//连接远程RPC服务 TCP
	//rpc, err := rpc.Dial("tcp", "127.0.0.1:8080")

	//链接远程RPC服务 jsonrpc
	rpc, err := jsonrpc.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		log.Fatal(err)
	}
	ret := 0
	//调用远程方法
	err2 := rpc.Call("Rect.Area", Params{50, 100}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(ret)
	err3 := rpc.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println(ret)
}
