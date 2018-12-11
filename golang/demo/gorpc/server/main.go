package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Params struct {
	Width, Height int
}

type Rect struct{}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	fmt.Println("ret:", *ret)
	return nil
}
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	fmt.Println("ret:", *ret)
	return nil
}

func chkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	rect := new(Rect)

	//注册一个rect服务
	rpc.Register(rect)

	//把服务处理绑定到Http协议上
	// rpc.HandleHTTP()
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//基于TCP的RPC
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080")
	chkError(err)
	tcplisten, err2 := net.ListenTCP("tcp", tcpaddr)
	chkError(err2)
	for {
		conn, err3 := tcplisten.Accept()
		if err3 != nil {
			continue
		}

		//基于tcp RPC
		//go rpc.ServeConn(conn)

		//jsonprc RPC使用json编码，而不是gob编码
		go jsonrpc.ServeConn(conn)
	}
}
