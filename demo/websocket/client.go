package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fasthttp/websocket"
	"sync"
	"sync/atomic"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second // 1分钟超时
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var id int64

// 每个连接对应一个client，client负责该连接数据的I/O
type Client struct {
	// 客户端ID
	id int64
	// socket conn
	conn *websocket.Conn
	// 消息写入处理
	writeCh chan interface{}
	// 锁
	mu sync.Mutex
}

// 创建客户端
func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		id:      atomic.AddInt64(&id, 1),
		conn:    conn,
		writeCh: make(chan interface{}, 256),
	}
}

// 启动服务
func (c *Client) startServe() {
	// 接收消息
	go c.runReader()
	// 发送消息
	go c.runWriter()
}

type TestRequest struct {
	Msg string `json:"msg"`
}

// 接收消息
func (c *Client) runReader() {
	// 设置读取最大消息数量
	c.conn.SetReadLimit(maxMessageSize)
	// 设置超时时间 1分钟
	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		fmt.Println("设置超时时间错误：", err.Error())
	}
	c.conn.SetPongHandler(func(string) error {
		return c.conn.SetReadDeadline(time.Now().Add(pongWait))
	})
	// 阻塞监听读取消息
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			fmt.Println("message:", err)
			break
		}
		var repSocket SocketRep
		repSocket.TaskLog = string(message) + " OK"
		c.writeCh <- repSocket

	}
}

// 关闭客户端
func (c *Client) close() {
	c.mu.Lock()
	defer c.mu.Unlock()
}

// 发送消息
func (c *Client) runWriter() {
	_, cancel := context.WithCancel(context.Background())
	// 定时 Ticker
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		cancel()
		ticker.Stop()
		_ = c.conn.Close()
	}()

	for {
		select {
		// push message 到客户端
		case message := <-c.writeCh:
			err := c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				c.close()
				return
			}

			// 序列化 message
			buf, err := json.Marshal(message)
			if err != nil {
				continue
			}

			// 写入到客户端
			err = c.conn.WriteMessage(websocket.TextMessage, buf)
			if err != nil {
				c.close()
				return
			}

			// 定时 保活
		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			err := c.conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				c.close()
				return
			}
		}
	}
}
