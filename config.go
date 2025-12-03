package hopkey

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Bindings []Binding
}

// LoadConfig loads the configuration from the given path.
// If the path is empty, it defaults to ~/.config/hopkey.toml.
func LoadConfig(path string) (*Config, error) {
	var config Config

	if path == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("could not determine user home directory: %w", err)
		}

		path = filepath.Join(home, ".config", "hopkey.toml")
	}

	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to load config from %s: %w", path, err)
	}

	return &config, nil
}
