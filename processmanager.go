package libclient

import "os"

type ProcessManagerClient struct {
	*Connection
}

func ConnectProcessManager() (lm *ProcessManagerClient, err error) {
	conn, err := Connect("/system/socket/ProcessSocket")
	lm = &ProcessManagerClient{conn}
	return
}

// register to ProcessManager
func (conn *ProcessManagerClient) Register(data map[string]string) {
	flexibleData := make(map[string]interface{}, len(data))
	for key, val := range data {
		flexibleData[key] = val
	}
	flexibleData["pid"] = os.Getpid()
	conn.Send("register", flexibleData)
}

// Process an application
func (*ProcessManagerClient) Process() {
}
