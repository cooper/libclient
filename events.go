package libclient

// LaunchManager

var LaunchEventHandlers = make(map[string]func(conn *Connection, name string, params map[string]interface{}))

func createLaunchEventHandlers() {
}

// ProcessManager

var ProcessEventHandlers = make(map[string]func(conn *Connection, name string, params map[string]interface{}))

func createProcessEventHandlers() {
	ProcessEventHandlers["ping"] = pingHandler
}

func pingHandler(conn *Connection, _ string, _ map[string]interface{}) {
	conn.Send("pong", nil)
}
