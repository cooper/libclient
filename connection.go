package libclient

import (
	"bufio"
	"encoding/json"
	"net"
	"os"
)

type Connection struct {
	path     string
	socket   net.Conn
	incoming *bufio.Reader
	outgoing *bufio.Writer
}

func Connect(path string) (conn *Connection, err error) {

	// check if file exists. if not, bail.
	_, err = os.Lstat(path)
	if err != nil {
		return
	}

	// resolve the address
	addr, err := net.ResolveUnixAddr("unix", path)
	if err != nil {
		return
	}

	// connect
	unixConn, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		return
	}

	conn = &Connection{
		path:     path,
		socket:   unixConn,
		incoming: bufio.NewReader(unixConn),
		outgoing: bufio.NewWriter(unixConn),
	}

	return
}

// send a JSON event
func (conn *Connection) Send(command string, params map[string]interface{}) bool {
	b, err := json.Marshal(params)
	if err != nil {
		return false
	}
	_, err = conn.outgoing.Write(b)
	if err != nil {
		return false
	}
	return true
}

// read data from a connection continuously
func (conn *Connection) Run() {
	for {
		line, _, err := conn.incoming.ReadLine()
		if err != nil {
			return
		}
		conn.handleEvent(line)
	}
}

// handle a JSON event
func (conn *Connection) handleEvent(data []byte) bool {
	var i interface{}
	err := json.Unmarshal(data, &i)
	if err != nil {
		return false
	}

	// should be an array.
	c := i.([]interface{})

	command := c[0].(string)
	params := c[1].(map[string]interface{})

	// if a handler for this command exists, run it
	if EventHandlers[command] != nil {
		EventHandlers[command](conn, command, params)
	}

	return true
}
