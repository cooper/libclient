package libclient

// LaunchManager

var LaunchEventHandlers = make(map[string]func(conn *Connection, name string, params map[string]interface{}))

// ProcessManager

var ProcessEventHandlers = make(map[string]func(conn *Connection, name string, params map[string]interface{}))
