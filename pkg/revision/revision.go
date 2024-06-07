package revision

import (
	"os"
)

var (
	ServiceName    string
	ServiceVersion string
	GitCommit      string
)

func init() {
	// Provided these env's are in the Dockerfile
	ServiceName = os.Getenv("SERVICE_NAME")
	ServiceVersion = os.Getenv("SERVICE_VERSION")
	GitCommit = os.Getenv("GIT_COMMIT")
}
