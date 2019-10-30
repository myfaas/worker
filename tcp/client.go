package main

import (
	"net"
	"log"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
    if err != nil {
        log.Println("dial error:", err)
        return
    }
	defer conn.Close()
	log.Println("dial ok")
	
	data := "hey world"
	
	conn.Write([]byte(data))
}