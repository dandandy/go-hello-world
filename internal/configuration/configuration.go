package configuration

import (
	"fmt"
	"os"
)

const (
	applicationNameEnvironmentVariable = "APPLICATION_NAME"
	versionEnvironmentVariable         = "VERSION"
	lastCommitShaEnvironmentVariable   = "LAST_COMMIT_SHA"
	descriptionEnvironmentVariable     = "DESCRIPTION"
)

const (
	environmentErrorString = "environment variable %s not set"
)

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

func Load() (Bundle, error) {
	name, err := lookupEnvironmentVariable(applicationNameEnvironmentVariable)
	if err != nil {
		return Bundle{}, err
	}

	sha, err := lookupEnvironmentVariable(lastCommitShaEnvironmentVariable)
	if err != nil {
		return Bundle{}, err
	}

	version, err := lookupEnvironmentVariable(versionEnvironmentVariable)
	if err != nil {
		return Bundle{}, err
	}

	description, err := lookupEnvironmentVariable(descriptionEnvironmentVariable)
	if err != nil {
		return Bundle{}, err
	}

	return Bundle{
		applicationName: name,
		version:         version,
		lastCommitSha:   sha,
		description:     description,
	}, nil
}

func lookupEnvironmentVariable(variableName string) (string, error) {
	value, ok := os.LookupEnv(variableName)
	if !ok {
		return "", fmt.Errorf(environmentErrorString, variableName)
	}
	return value, nil
}
