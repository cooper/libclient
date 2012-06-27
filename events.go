package libclient

// LaunchManager

var LaunchEventHandlers map[string]func(*Connection, string, map[string]interface{})

// create LaunchManager event handlers.
func createLaunchEventHandlers() {
	LaunchEventHandlers = make(map[string]func(*Connection, string, map[string]interface{}))
}

// ProcessManager

var ProcessEventHandlers map[string]func(*Connection, string, map[string]interface{})

// create ProcessManager event handlers.
func createProcessEventHandlers() {
	ProcessEventHandlers = make(map[string]func(*Connection, string, map[string]interface{}))
	ProcessEventHandlers["ping"] = pingHandler
}

// keep the connection alive.
func pingHandler(conn *Connection, _ string, _ map[string]interface{}) {
	conn.Send("pong", nil)
}
