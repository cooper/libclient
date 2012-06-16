package libclient

import (
	"os"
	"net"
	"bufio"
)

type Connection struct {
	path     string
	conn     net.Conn
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
		path: path,
		conn: unixConn,
		incoming: bufio.NewReader(unixConn),
		outgoing: bufio.NewWriter(unixConn),
	}

	return
}
