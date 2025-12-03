package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/bajankristof/hopkey/internal/cli"
	"golang.design/x/hotkey/mainthread"
)

func main() {
	mainthread.Init(func() {
		app := cli.NewApp()
		err := app.Run(context.Background(), os.Args)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
	})
}
