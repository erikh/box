package executor

import (
	"io"

	"github.com/erikh/box/builder/config"
)

// Hook is a hook used in commit calls
type Hook func(id string) (string, error)

// Executor is an engine for talking to different layering/execution context
// subsystems. It is the meat-and-potatoes of image building.
type Executor interface {
	// LoadConfig loads the configuration into the executor.
	LoadConfig(*config.Config) error

	// Config returns the current *Config for the executor.
	Config() *config.Config

	// ImageID returns the image identifier of the most recent layer.
	ImageID() string

	// Commit commits an entry to the layer list.
	Commit(string, Hook) error

	// CheckCache consults the cache to see if there are any items which fit it.
	CheckCache(string) (bool, error)

	// CopyToContainer copies a tarred up series of files (passed in through the
	// io.Reader handle) to the container where they are untarred.
	CopyToContainer(string, string, io.Reader) error

	// CopyFromContainer copies a series of files in a similar fashion to
	// CopyToContainer, just in reverse.
	CopyFromContainer(string, string) (io.Reader, error)

	// CopyOneFileFromContainer copies a file from the container and returns its content.
	CopyOneFileFromContainer(string) ([]byte, error)

	Create() (string, error)
	Destroy(string) error

	Tag(string) error
	Fetch(string) (string, error)

	RunHook(string) (string, error)
}
