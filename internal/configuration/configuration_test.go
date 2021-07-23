package configuration

import (
	"testing"
)

func TestBundle_GetBundleValues(t *testing.T) {
	app := "foo"
	ver := "ver"
	sha := "abcd123"
	desc := "bazz"
	c := &Bundle{
		applicationName: app,
		version:         ver,
		lastCommitSha:   sha,
		description:     desc,
	}

	if got := c.GetApplicationName(); got != app {
		t.Errorf("Bundle.GetApplicationName() = %v, want %v", got, app)
	}

	if got := c.GetVersion(); got != ver {
		t.Errorf("Bundle.GetVersion() = %v, want %v", got, ver)
	}

	if got := c.GetLastCommitSha(); got != sha {
		t.Errorf("Bundle.GetLastCommitSha() = %v, want %v", got, sha)
	}

	if got := c.GetDescription(); got != desc {
		t.Errorf("Bundle.GetDescription() = %v, want %v", got, desc)
	}
}
