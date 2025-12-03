package hopkey

import "errors"

type Modifier int

var ErrUnknownModifier = errors.New("unknown modifier")

const (
	ModifierCommand = 1 << 8
	ModifierShift   = 1 << 9
	ModifierOption  = 1 << 11
	ModifierControl = 1 << 12
)

// ParseModifier parses a modifier from its string representation.
func ParseModifier(text []byte) (Modifier, error) {
	switch string(text) {
	case "cmd":
		return ModifierCommand, nil
	case "command":
		return ModifierCommand, nil
	case "shift":
		return ModifierShift, nil
	case "opt":
		return ModifierOption, nil
	case "option":
		return ModifierOption, nil
	case "ctrl":
		return ModifierControl, nil
	case "control":
		return ModifierControl, nil
	default:
		return 0, ErrUnknownModifier
	}
}

func (m Modifier) String() string {
	switch m {
	case ModifierCommand:
		return "command"
	case ModifierShift:
		return "shift"
	case ModifierOption:
		return "option"
	case ModifierControl:
		return "control"
	default:
		return "unknown"
	}
}

func (m *Modifier) UnmarshalText(text []byte) error {
	mod, err := ParseModifier(text)
	*m = mod
	return err
}
