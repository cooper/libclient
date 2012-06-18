package libclient

import "time"

var (
	Launch *LaunchManagerClient
)

// one function that does it all for convenience
func Loop() {

}

// launchmanager loop
func LM(data map[string]string) {
	var err error
	for {
		Launch, err = ConnectLaunchManager()

		// wait five seconds before trying again...
		if err != nil {
			time.Sleep(5)
		} else {
			Launch.Register(data)
			Launch.Run()
		}
	}
}
