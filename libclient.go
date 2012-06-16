package libclient

import "time"

var (
	Launch *LaunchManagerClient
)

// one function that does it all for convenience
func Loop() {
	var err error

	// launchmanager goroutine
	go func() {
		for {
			Launch, err = ConnectLaunchManager()

			// wait five seconds before trying again...
			if err != nil {
				time.Sleep(5)
			} else {
				Launch.Run()
			}
		}
	}()
}
