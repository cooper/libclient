package libclient

import "os"

type ProcessManagerClient struct {
	*Connection
}

func ConnectProcessManager() (pm *ProcessManagerClient, err error) {
	conn, err := Connect("/system/socket/ProcessSocket")
	pm = &ProcessManagerClient{conn}
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
