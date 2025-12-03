package cli

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/bajankristof/hopkey"
	"github.com/urfave/cli/v3"
)

func NewApp() *cli.Command {
	return &cli.Command{
		Name:    "hopkey",
		Usage:   "A hotkey daemon for macOS",
		Version: "0.1.0",
		Commands: []*cli.Command{
			NewInstallCommand(),
			NewReloadCommand(),
			NewUninstallCommand(),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "config file path (default: ~/.config/hopkey.toml)",
				Value:   "",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			configPath := cmd.String("config")
			config, err := hopkey.LoadConfig(configPath)
			if err != nil {
				return err
			}

			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			for _, binding := range config.Bindings {
				err := binding.Go(ctx)
				if err != nil {
					return err
				}
			}

			signals := make(chan os.Signal, 1)
			signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
			<-signals

			return nil
		},
	}
}
