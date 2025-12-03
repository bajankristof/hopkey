# hopkey

A hotkey daemon for macOS, written in Go.

## Features

- Define global hotkeys to execute commands or scripts.
- Installable as a launch agent for automatic startup.
- Simple configuration using a TOML file.
- As lightweight as it gets.

## Installation

***THIS DOES NOT WORK YET***

```bash
brew install bajankristof/tap/hopkey
```

### Notarization

**By using hopkey, you acknowledge that it has not been notarized by Apple.**
The Homebrew installation process includes a step to remove the quarantine attribute,
but you may still encounter security warnings when running the application for the first time.
To bypass these warnings, you can go to the "Security & Privacy" settings in macOS and allow the application to run.

## Configuration

Create a configuration file at `~/.config/hopkey.toml`:

```toml
[[bindings]]
exec = "open -a Terminal"
key = "T"
modifiers = ["cmd", "alt"]

[[bindings]]
exec = "osascript -e 'display notification \"Hello, World!\" with title \"Notification\"'"
key = "N"
modifiers = ["ctrl", "shift"]
```

## Usage

```bash
hopkey
# or
hopkey --config /path/to/your/config.toml
```

### Install Launch Agent
```bash
hopkey i
# or
hopkey install
# or
hopkey i --config /path/to/your/config.toml
```

### Uninstall Launch Agent
```bash
hopkey u
# or
hopkey uninstall
```

### Reload Configuration
```bash
hopkey r
# or
hopkey reload
```
