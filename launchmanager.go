package libclient

import (
	"os"
	"time"
)

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

// launchmanager loop
func RunLaunch() {
	var err error
	for {
		Launch, err = ConnectLaunchManager()

		// first of all, this should never happen.
		// wait five seconds before trying again...
		if err != nil {
			time.Sleep(5)
		} else {
			Launch.Register()
			Launch.Run()
		}
	}
}
