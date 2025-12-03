package cli

import (
	"context"
	"os"

	"github.com/bajankristof/hopkey"
	"github.com/urfave/cli/v3"
)

func NewInstallCommand() *cli.Command {
	return &cli.Command{
		Name:    "install",
		Aliases: []string{"i"},
		Usage:   "Installs the hotkey launch daemon",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			daemon, err := newDaemon(cmd)
			if err != nil {
				return err
			}
			return daemon.Install()
		},
	}
}

func NewReloadCommand() *cli.Command {
	return &cli.Command{
		Name:    "reload",
		Aliases: []string{"r"},
		Usage:   "Reloads the hotkey launch daemon",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			daemon, err := newDaemon(cmd)
			if err != nil {
				return err
			}
			return daemon.Reload()
		},
	}
}

func NewUninstallCommand() *cli.Command {
	return &cli.Command{
		Name:    "uninstall",
		Aliases: []string{"u"},
		Usage:   "Uninstalls the hotkey launch daemon",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			daemon, err := newDaemon(cmd)
			if err != nil {
				return err
			}
			return daemon.Uninstall()
		},
	}
}

func newDaemon(cmd *cli.Command) (hopkey.Daemon, error) {
	execPath, err := os.Executable()
	if err != nil {
		return hopkey.Daemon{}, err
	}

	configPath := cmd.String("config")
	return hopkey.Daemon{
		ExecPath:   execPath,
		ConfigPath: configPath,
	}, nil
}
