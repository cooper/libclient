package libclient

import "time"

var (
	Process *ProcessManagerClient
)

// one function that does it all for convenience
func Loop() {

}

// Processmanager loop
func Launch(data map[string]string) {
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
