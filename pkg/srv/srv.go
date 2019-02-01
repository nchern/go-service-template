package srv

import "log"

// Start runs the api server
func Start(name string) {
	log.Printf("Noop %s srv is running", name)
}

// Stop performs all necessary operations for graceful service tear down.
func Stop() {
	log.Println("Stopping...")
}
