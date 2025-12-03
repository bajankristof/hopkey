package hopkey

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"

	"golang.design/x/hotkey"
)

type Binding struct {
	Exec      string
	Key       Key
	Modifiers []Modifier
}

func (b Binding) String() string {
	mods := ""
	for i, mod := range b.Modifiers {
		if i > 0 {
			mods += "+"
		}
		mods += mod.String()
	}
	return fmt.Sprintf("%s+%s -> %s", mods, b.Key.String(), b.Exec)
}

// Go registers the hotkey and starts listening for events
// (in a goroutine) until the context is cancelled.
func (b Binding) Go(ctx context.Context) error {
	key := hotkey.Key(b.Key)
	mods := []hotkey.Modifier{}
	for _, mod := range b.Modifiers {
		mods = append(mods, hotkey.Modifier(mod))
	}

	hk := hotkey.New(mods, key)
	err := hk.Register()
	if err != nil {
		return fmt.Errorf("failed to register binding %s: %w", b, err)
	}

	slog.Info("binding registered", "binding", b.String())
	go func() {
		for {
			select {
			case <-hk.Keydown():
				slog.Info("binding triggered", "binding", b.String())
				err := b.exec()
				if err != nil {
					slog.Error(err.Error())
					continue
				}
			case <-ctx.Done():
				hk.Unregister()
				return
			}
		}
	}()

	return nil
}

// exec runs the command specified in the binding.
func (b *Binding) exec() error {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "sh"
	}

	cmd := exec.Command(shell, "-c", b.Exec)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute binding %s: %w\n%s", b, err, string(output))
	}

	return nil
}
