package execex

import (
	"context"
	"os"
	"os/exec"

	"github.com/rstudio/goex/zapex"
)

var (
	// Sh runs a given command with context and returns the stdout
	// bytes and error, if any. The inner command's stdin and
	// stderr are inherited from os.Stdin and os.Stderr.
	Sh ShellOutputterFunc = defaultShellOutputter

	// Run runs a given command with context and returns the error,
	// if any. The inner command's stdin, stdout, and stderr are
	// inherited from os.Stdin, os.Stdout, and os.Stderr.
	Run ShellRunnerFunc = defaultShellRunner

	// ShellDryRun is a global switch to turn off running commands
	ShellDryRun bool
)

type ShellOutputterFunc func(ctx context.Context, exe string, argv ...string) ([]byte, error)

type ShellRunnerFunc func(ctx context.Context, exe string, argv ...string) error

func defaultShellOutputter(ctx context.Context, exe string, argv ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, exe, argv...)

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if logger, ok := zapex.LoggerFromContext(ctx); ok {
		if ShellDryRun {
			logger.Debugw("not actually getting output from command", "cmd", cmd)
			return []byte(""), nil
		}

		logger.Debugw("getting output from command", "cmd", cmd)
	}

	if ShellDryRun {
		return []byte(""), nil
	}

	return cmd.Output()
}

func defaultShellRunner(ctx context.Context, exe string, argv ...string) error {
	cmd := exec.CommandContext(ctx, exe, argv...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	if logger, ok := zapex.LoggerFromContext(ctx); ok {
		if ShellDryRun {
			logger.Debugw("not actually running command", "cmd", cmd)
			return nil
		}

		logger.Debugw("running command", "cmd", cmd)
	}

	if ShellDryRun {
		return nil
	}

	return cmd.Run()
}
