package execex_test

import (
	"context"
	"strings"
	"testing"

	"github.com/rstudio/goex/os/execex"
	"github.com/stretchr/testify/require"
)

func TestSh(t *testing.T) {
	r := require.New(t)

	out, err := execex.Sh(context.Background(), "echo", "magenta mountain")
	r.Nil(err)
	r.Equal("magenta mountain", strings.TrimSpace(string(out)))

	execex.ShellDryRun = true

	t.Cleanup(func() {
		execex.ShellDryRun = false
	})

	out, err = execex.Sh(context.Background(), "echo", "magenta mountain")
	r.Nil(err)
	r.Equal("", strings.TrimSpace(string(out)))

	out, err = execex.Sh(context.Background(), "unlikely2work", "\x00\x00\x00\x00")
	r.Nil(err)
	r.Equal("", strings.TrimSpace(string(out)))
}

func TestRun(t *testing.T) {
	r := require.New(t)

	err := execex.Run(context.Background(), "echo", "grim reaper", ",", "the")
	r.Nil(err)

	execex.ShellDryRun = true

	t.Cleanup(func() {
		execex.ShellDryRun = false
	})

	err = execex.Run(context.Background(), "unlikely2work", "\x00\x00\x00\x00")
	r.Nil(err)
}
