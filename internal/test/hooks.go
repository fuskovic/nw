package test

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

// WithNw is a pre-test hook that asserts the nw binary is globally installed before running the test.
// It's intended to be used in the cmd/nw pkg.
func WithNw(t *testing.T, name string, fn func(t *testing.T)) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		// make sure nw is installed
		cmd := exec.Command("which", "nw")
		require.NoError(t, cmd.Run(), "nw not installed")
		fn(t)
	})
}
