package libclient

import "os"

type LaunchManagerClient struct {
	*Connection
}

func ConnectLaunchManager() (lm *LaunchManagerClient, err error) {
	conn, err := Connect("/system/socket/LaunchSocket")
	lm = &LaunchManagerClient{conn}
	return
}

// register to LaunchManager
func (conn *LaunchManagerClient) Register(data map[string]string) {
	flexibleData := make(map[string]interface{}, len(data))
	for key, val := range data {
		flexibleData[key] = val
	}
	flexibleData["pid"] = os.Getpid()
	conn.Send("register", flexibleData)
}

// launch an application
func (*LaunchManagerClient) Launch() {
}
