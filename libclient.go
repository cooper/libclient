package libclient

import "time"

var (
	Process *ProcessManagerClient
	Launch  *LaunchManagerClient
)

// one function that does it all for convenience
func Loop() {

}

// processmanager loop
func RunProcess(data map[string]string) {
	var err error
	for {
		Process, err = ConnectProcessManager()

		// wait five seconds before trying again...
		if err != nil {
			time.Sleep(5)
		} else {
			Process.Register(data)
			Process.Run()
		}
	}
}

// launchmanager loop
func RunLaunch() {
	var err error
	for {
		Launch, err = ConnectLaunchManager()

		// wait five seconds before trying again...
		if err != nil {
			time.Sleep(5)
		} else {
			Launch.Register()
			Launch.Run()
		}
	}
}
