package hopkey

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

var plist = template.Must(
	template.
		New("plist").
		Parse(
			`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
		<key>Label</key>
		<string>com.bajankristof.hopkey</string>
		<key>ProgramArguments</key>
		<array>
				<string>{{.ExecPath}}</string>
				{{if .ConfigPath}}
				<string>-c</string>
				<string>{{.ConfigPath}}</string>
				{{end}}
		</array>
		<key>RunAtLoad</key>
		<true/>
		<key>KeepAlive</key>
		<true/>
</dict>
</plist>
`,
		),
)

type Daemon struct {
	ExecPath   string
	ConfigPath string
	agentPath  string
}

// Install creates and loads the LaunchAgent plist.
func (d *Daemon) Install() error {
	p, err := d.AgentPath()
	if err != nil {
		return err
	}

	f, err := os.Create(p)
	if err != nil && errors.Is(err, os.ErrExist) {
		f, err = os.OpenFile(p, os.O_WRONLY|os.O_TRUNC, 0644)
	}
	if err != nil {
		return fmt.Errorf("failed to write launch agent plist: %w", err)
	}
	defer f.Close()

	err = plist.Execute(f, d)
	if err != nil {
		return fmt.Errorf("failed to write launch agent plist: %w", err)
	}

	return d.Reload()
}

// Uninstall unloads and removes the LaunchAgent plist.
func (d *Daemon) Uninstall() error {
	err := d.Unload()
	if err != nil {
		return err
	}

	p, err := d.AgentPath()
	if err != nil {
		return err
	}

	err = os.Remove(p)
	if err != nil {
		return fmt.Errorf("failed to remove launch agent plist: %w", err)
	}

	return nil
}

// Reload unloads and loads the LaunchAgent plist.
func (d *Daemon) Reload() error {
	err := d.Unload()
	if err != nil {
		return err
	}

	return d.Load()
}

// Load loads the LaunchAgent plist.
func (d *Daemon) Load() error {
	p, err := d.AgentPath()
	if err != nil {
		return err
	}

	cmd := exec.Command("launchctl", "load", p)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to load launch agent: %s: %w", string(output), err)
	}

	return nil
}

// Unload unloads the LaunchAgent plist.
func (d *Daemon) Unload() error {
	p, err := d.AgentPath()
	if err != nil {
		return err
	}

	cmd := exec.Command("launchctl", "unload", p)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to unload launch agent: %s: %w", string(output), err)
	}

	return nil
}

// AgentPath returns the path to the LaunchAgent plist file.
func (d *Daemon) AgentPath() (string, error) {
	if d.agentPath != "" {
		return d.agentPath, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine user home directory: %w", err)
	}

	d.agentPath = filepath.Join(
		home,
		"Library",
		"LaunchAgents",
		"io.bajankristof.hopkey.plist",
	)

	return d.agentPath, nil
}
