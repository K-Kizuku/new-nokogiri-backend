package websocket

import (
	"encoding/json"
	"log"
	"nokogiri-api/db"
	"sync"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024
)

// var (
// 	newline = []byte{'\n'}
// 	space   = []byte{' '}
// )

var upgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *fasthttp.RequestCtx) bool {
		return true
	},
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	room *Room

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	id string

	mu sync.Mutex
}

func (c *Client) readPump() {
	defer func() {
		c.room.unregister <- c
		c.conn.Close()
	}()
	// var v V
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		// err := c.conn.ReadJSON(v)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		// message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		dat := db.Post{}

		if err := json.Unmarshal(message, &dat); err != nil {
			log.Printf("error: %v", err)
		}
		// data, _ := json.Marshal(&dat)
		// c.room.broadcast <- []byte(data)
		c.room.mu.Lock()
		c.room.senddata.Room = 1
		pos := db.Pointer{PointerX: dat.Pos.PointerX, PointerY: dat.Pos.PointerY}
		c.room.senddata.Pos[c.id] = pos
		c.room.senddata.Score[c.id] = dat.Score
		// GenerateQ(c.room.question["test"])
		c.room.senddata.Question = c.room.question["test"]
		data, _ := json.Marshal(&c.room.senddata)
		c.room.mu.Unlock()
		c.room.broadcast <- []byte(data)
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	c.conn.WriteMessage(websocket.TextMessage, []byte(c.id))
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			// c.conn.WriteJSON()
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				// w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (room *Room) ServeWs(ctx *fiber.Ctx) error {
	err := upgrader.Upgrade(ctx.Context(), func(conn *websocket.Conn) {
		u, _ := uuid.NewRandom()
		client := &Client{room: room, conn: conn, send: make(chan []byte, 512), id: u.String()}
		client.room.register <- client
		log.Println("tets")
		log.Printf("uuid: %v", u.String())

		go client.writePump()
		client.readPump()
	})

	if err != nil {
		log.Println(err)
	}
	return nil
}
