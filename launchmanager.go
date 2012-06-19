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
func (conn *LaunchManagerClient) Register() {
	conn.Send("register", map[string]interface{}{
		"pid": os.Getpid(),
	})
}

// Launch an application
func (*LaunchManagerClient) Launch() {
}
