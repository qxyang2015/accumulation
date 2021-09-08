package websocket

import (
	"github.com/fasthttp/websocket"
	"log"
	"net/http"
)

func SocketServerRun(port, path string) {
	http.HandleFunc(path, WebSocket)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("listen出现错误", err)
	}
}

// Web Socket 执行方法
func WebSocket(w http.ResponseWriter, r *http.Request) {
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	//将hhtp协议升级到socket协议
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	// 每次访问path 创建一个Client
	NewClient(conn).startServe()
}

type SocketRep struct {
	TaskLog string `json:"taskLog"`
}
