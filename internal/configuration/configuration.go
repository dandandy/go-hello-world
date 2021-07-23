package configuration

// These configuration variables are overriden at build time by setting the ldflags.
var (
	// these have default values set
	applicationName = "go-hello-world"
	description     = "A simple Go Web API application"

	// these default to empty string
	version       string
	lastCommitSha string
)

// Expose configuration via this bundle struct.
// Use Getter methods to ensure fields don't get overriden.
type Bundle struct {
	applicationName string
	version         string
	lastCommitSha   string
	description     string
}

func (c *Bundle) GetApplicationName() string {
	return c.applicationName
}

func (c *Bundle) GetVersion() string {
	return c.version
}

func (c *Bundle) GetLastCommitSha() string {
	return c.lastCommitSha
}

func (c *Bundle) GetDescription() string {
	return c.description
}

func Load() Bundle {
	return Bundle{
		applicationName: applicationName,
		version:         version,
		lastCommitSha:   lastCommitSha,
		description:     description,
	}
}
