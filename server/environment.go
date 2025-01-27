package server

import (
	"os"
)

// Defines the basic interface that all environments need to implement so that
// a server can be properly controlled.
type Environment interface {
	// Returns the name of the environment.
	Type() string

	// Determines if the environment is currently active and running a server process
	// for this specific server instance.
	IsRunning() (bool, error)

	// Starts a server instance. If the server instance is not in a state where it
	// can be started an error should be returned.
	Start() error

	// Stops a server instance. If the server is already stopped an error should
	// not be returned.
	Stop() error

	// Determines if the server instance exists. For example, in a docker environment
	// this should confirm that the container is created and in a bootable state. In
	// a basic CLI environment this can probably just return true right away.
	Exists() (bool, error)

	// Terminates a running server instance using the provided signal. If the server
	// is not running no error should be returned.
	Terminate(signal os.Signal) error

	// Creates the necessary environment for running the server process. For example,
	// in the Docker environment create will create a new container instance for the
	// server.
	Create() error

	// Attaches to the server console environment and allows piping the output to a
	// websocket or other internal tool to monitor output. Also allows you to later
	// send data into the environment's stdin.
	Attach() error

	// Follows the output from the server console and will begin piping the output to
	// the server's emitter.
	FollowConsoleOutput() error

	// Sends the provided command to the running server instance.
	SendCommand(string) error

	// Reads the log file for the process from the end backwards until the provided
	// number of bytes is met.
	Readlog(int64) ([]string, error)

	// Polls the given environment for resource usage of the server when the process
	// is running.
	EnableResourcePolling() error

	// Disables the polling operation for resource usage and sets the required values
	// to 0 in the server resource usage struct.
	DisableResourcePolling() error
}
