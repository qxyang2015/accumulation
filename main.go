package main

import (
	"fmt"
	"github.com/qxyang2015/accumulation/demo/websocket"
)

func main() {
	fmt.Println("start")
	port := ":9010"
	path := "/websocket"
	websocket.SocketServerRun(port, path)
	fmt.Println("end")
}
