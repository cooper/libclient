package libclient

var EventHandlers = make(map[string]func(conn *Connection, name string, params map[string]interface{}))
