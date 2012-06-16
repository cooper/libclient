package libclient

type LaunchManagerClient struct {
	*Connection
}

func ConnectLaunchManager() (lm *LaunchManagerClient, err error) {
	conn, err := Connect("/system/socket/LaunchSocket")
	lm = &LaunchManagerClient{conn}
	return
}

func (*LaunchManagerClient) Launch() {
}
