package main

import (
	"net"
	"log"
)

func handleConn(c net.Conn) {
	defer c.Close()
	var buf = make([]byte, 10)
    log.Println("start to read from conn")
	n, _ := c.Read(buf)
	log.Println(n)
	log.Println(string(buf))
	log.Println("请求已处理")
}

func main() {
	// 监听
	conn, err := net.Listen("tcp", ":8888")
    if err != nil {
        log.Println("error listen:", err)
        return
    }
    defer conn.Close()
	log.Println("listen ok")
	
	// 接收请求
	for {
        c, err := conn.Accept()
        if err != nil {
            log.Println("accept error:", err)
            break
		}
		// 开启一个goroutine去处理请求
        go handleConn(c)
    }

}
