package configuration

import (
	"os"
	"reflect"
	"testing"
)

func Test_lookupEnvironmentVariable(t *testing.T) {
	t.Run("given environment variable set, return value", func(t *testing.T) {
		os.Setenv(applicationNameEnvironmentVariable, "foo")
		defer func() {
			os.Unsetenv(applicationNameEnvironmentVariable)
		}()

		want := "foo"
		wantErr := false
		got, err := lookupEnvironmentVariable(applicationNameEnvironmentVariable)
		if (err != nil) != wantErr {
			t.Errorf("lookupEnvironmentVariable() error = %v, wantErr %v", err, wantErr)
			return
		}
		if got != want {
			t.Errorf("lookupEnvironmentVariable() = %v, want %v", got, want)
		}
	})

	t.Run("given environment not variable set, return error", func(t *testing.T) {
		os.Unsetenv(applicationNameEnvironmentVariable)
		wantErr := true
		want := ""
		got, err := lookupEnvironmentVariable(applicationNameEnvironmentVariable)
		if (err != nil) != wantErr {
			t.Errorf("lookupEnvironmentVariable() error = %v, wantErr %v", err, wantErr)
			return
		}
		if got != want {
			t.Errorf("lookupEnvironmentVariable() = %v, want %v", got, want)
		}
	})
}

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		want    Bundle
		wantErr bool
		before  func()
	}{
		{
			name: "return config bundle given all env vars set",
			want: Bundle{
				applicationName: "foo",
				version:         "1",
				lastCommitSha:   "abcd123",
				description:     "bazz",
			},
			before: func() {
				os.Setenv(applicationNameEnvironmentVariable, "foo")
				os.Setenv(versionEnvironmentVariable, "1")
				os.Setenv(lastCommitShaEnvironmentVariable, "abcd123")
				os.Setenv(descriptionEnvironmentVariable, "bazz")
			},
		},
		{
			name: "return error given app name env var not set",
			before: func() {
				os.Setenv(versionEnvironmentVariable, "1")
				os.Setenv(lastCommitShaEnvironmentVariable, "abcd123")
				os.Setenv(descriptionEnvironmentVariable, "bazz")
			},
			wantErr: true,
		},
		{
			name: "return error given version env var not set",
			before: func() {
				os.Setenv(applicationNameEnvironmentVariable, "foo")
				os.Setenv(lastCommitShaEnvironmentVariable, "abcd123")
				os.Setenv(descriptionEnvironmentVariable, "bazz")
			},
			wantErr: true,
		},
		{
			name: "return error given commit sha env var not set",
			before: func() {
				os.Setenv(applicationNameEnvironmentVariable, "foo")
				os.Setenv(versionEnvironmentVariable, "1")
				os.Setenv(descriptionEnvironmentVariable, "bazz")
			},
			wantErr: true,
		},
		{
			name: "return error given description env var not set",
			before: func() {
				os.Setenv(applicationNameEnvironmentVariable, "foo")
				os.Setenv(versionEnvironmentVariable, "1")
				os.Setenv(lastCommitShaEnvironmentVariable, "abcd123")
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer clearEnv()
			tt.before()
			got, err := Load()
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
	}
}

func clearEnv() {
	envVars := []string{
		applicationNameEnvironmentVariable,
		versionEnvironmentVariable,
		descriptionEnvironmentVariable,
		lastCommitShaEnvironmentVariable,
	}

	for _, value := range envVars {
		os.Unsetenv(value)
	}
}

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
