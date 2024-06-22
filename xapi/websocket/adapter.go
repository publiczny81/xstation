package websocket

import "golang.org/x/net/websocket"

const (
	protocol = "tcp"
)

type Connection struct {
	*websocket.Conn
}

func (c *Connection) Send(msg any) error {
	return websocket.JSON.Send(c.Conn, msg)
}

func (c *Connection) Receive(msg any) error {
	return websocket.JSON.Receive(c.Conn, msg)
}

func Connect(url_, origin string) (c *Connection, err error) {
	var (
		conn *websocket.Conn
	)

	if conn, err = websocket.Dial(url_, protocol, origin); err != nil {
		return
	}
	c = &Connection{
		Conn: conn,
	}
	return
}
