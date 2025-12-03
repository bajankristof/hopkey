package hopkey

import "errors"

type Key int

var ErrUnknownKey = errors.New("unknown key")

const (
	KeyA              = 0x00
	KeyS              = 0x01
	KeyD              = 0x02
	KeyF              = 0x03
	KeyH              = 0x04
	KeyG              = 0x05
	KeyZ              = 0x06
	KeyX              = 0x07
	KeyC              = 0x08
	KeyV              = 0x09
	KeyB              = 0x0B
	KeyQ              = 0x0C
	KeyW              = 0x0D
	KeyE              = 0x0E
	KeyR              = 0x0F
	KeyY              = 0x10
	KeyT              = 0x11
	Key1              = 0x12
	Key2              = 0x13
	Key3              = 0x14
	Key4              = 0x15
	Key6              = 0x16
	Key5              = 0x17
	KeyEqual          = 0x18
	Key9              = 0x19
	Key7              = 0x1A
	KeyMinus          = 0x1B
	Key8              = 0x1C
	Key0              = 0x1D
	KeyRightBracket   = 0x1E
	KeyO              = 0x1F
	KeyU              = 0x20
	KeyLeftBracket    = 0x21
	KeyI              = 0x22
	KeyP              = 0x23
	KeyL              = 0x25
	KeyJ              = 0x26
	KeyQuote          = 0x27
	KeyK              = 0x28
	KeySemicolon      = 0x29
	KeyBackslash      = 0x2A
	KeyComma          = 0x2B
	KeySlash          = 0x2C
	KeyN              = 0x2D
	KeyM              = 0x2E
	KeyPeriod         = 0x2F
	KeyGrave          = 0x32
	KeyReturn         = 0x24
	KeyTab            = 0x30
	KeySpace          = 0x31
	KeyDelete         = 0x33
	KeyEscape         = 0x35
	KeyCapsLock       = 0x39
	KeyFunction       = 0x3F
	KeyF17            = 0x40
	KeyVolumeUp       = 0x48
	KeyVolumeDown     = 0x49
	KeyMute           = 0x4A
	KeyF18            = 0x4F
	KeyF19            = 0x50
	KeyF20            = 0x5A
	KeyF5             = 0x60
	KeyF6             = 0x61
	KeyF7             = 0x62
	KeyF3             = 0x63
	KeyF8             = 0x64
	KeyF9             = 0x65
	KeyF11            = 0x67
	KeyF13            = 0x69
	KeyF16            = 0x6A
	KeyF14            = 0x6B
	KeyF10            = 0x6D
	KeyContextualMenu = 0x6E
	KeyF12            = 0x6F
	KeyF15            = 0x71
	KeyHelp           = 0x72
	KeyHome           = 0x73
	KeyPageUp         = 0x74
	KeyForwardDelete  = 0x75
	KeyF4             = 0x76
	KeyEnd            = 0x77
	KeyF2             = 0x78
	KeyPageDown       = 0x79
	KeyF1             = 0x7A
	KeyLeftArrow      = 0x7B
	KeyRightArrow     = 0x7C
	KeyDownArrow      = 0x7D
	KeyUpArrow        = 0x7E
)

// ParseKey parses a key from its string representation.
func ParseKey(text []byte) (Key, error) {
	switch string(text) {
	case "a":
		return KeyA, nil
	case "s":
		return KeyS, nil
	case "d":
		return KeyD, nil
	case "f":
		return KeyF, nil
	case "h":
		return KeyH, nil
	case "g":
		return KeyG, nil
	case "z":
		return KeyZ, nil
	case "x":
		return KeyX, nil
	case "c":
		return KeyC, nil
	case "v":
		return KeyV, nil
	case "b":
		return KeyB, nil
	case "q":
		return KeyQ, nil
	case "w":
		return KeyW, nil
	case "e":
		return KeyE, nil
	case "r":
		return KeyR, nil
	case "y":
		return KeyY, nil
	case "t":
		return KeyT, nil
	case "1":
		return Key1, nil
	case "2":
		return Key2, nil
	case "3":
		return Key3, nil
	case "4":
		return Key4, nil
	case "6":
		return Key6, nil
	case "5":
		return Key5, nil
	case "=":
		return KeyEqual, nil
	case "9":
		return Key9, nil
	case "7":
		return Key7, nil
	case "-":
		return KeyMinus, nil
	case "8":
		return Key8, nil
	case "0":
		return Key0, nil
	case "]":
		return KeyRightBracket, nil
	case "o":
		return KeyO, nil
	case "u":
		return KeyU, nil
	case "[":
		return KeyLeftBracket, nil
	case "i":
		return KeyI, nil
	case "p":
		return KeyP, nil
	case "l":
		return KeyL, nil
	case "j":
		return KeyJ, nil
	case "'":
		return KeyQuote, nil
	case "k":
		return KeyK, nil
	case ";":
		return KeySemicolon, nil
	case "\\":
		return KeyBackslash, nil
	case ",":
		return KeyComma, nil
	case "/":
		return KeySlash, nil
	case "n":
		return KeyN, nil
	case "m":
		return KeyM, nil
	case ".":
		return KeyPeriod, nil
	case "`":
		return KeyGrave, nil
	case "return":
		return KeyReturn, nil
	case "tab":
		return KeyTab, nil
	case "space":
		return KeySpace, nil
	case "delete":
		return KeyDelete, nil
	case "escape":
		return KeyEscape, nil
	case "capslock":
		return KeyCapsLock, nil
	case "function":
		return KeyFunction, nil
	case "f17":
		return KeyF17, nil
	case "volumeup":
		return KeyVolumeUp, nil
	case "volumedown":
		return KeyVolumeDown, nil
	case "mute":
		return KeyMute, nil
	case "f18":
		return KeyF18, nil
	case "f19":
		return KeyF19, nil
	case "f20":
		return KeyF20, nil
	case "f5":
		return KeyF5, nil
	case "f6":
		return KeyF6, nil
	case "f7":
		return KeyF7, nil
	case "f3":
		return KeyF3, nil
	case "f8":
		return KeyF8, nil
	case "f9":
		return KeyF9, nil
	case "f11":
		return KeyF11, nil
	case "f13":
		return KeyF13, nil
	case "f16":
		return KeyF16, nil
	case "f14":
		return KeyF14, nil
	case "f10":
		return KeyF10, nil
	case "contextualmenu":
		return KeyContextualMenu, nil
	case "f12":
		return KeyF12, nil
	case "f15":
		return KeyF15, nil
	case "help":
		return KeyHelp, nil
	case "home":
		return KeyHome, nil
	case "pageup":
		return KeyPageUp, nil
	case "forwarddelete":
		return KeyForwardDelete, nil
	case "f4":
		return KeyF4, nil
	case "end":
		return KeyEnd, nil
	case "f2":
		return KeyF2, nil
	case "pagedown":
		return KeyPageDown, nil
	case "f1":
		return KeyF1, nil
	case "left":
		return KeyLeftArrow, nil
	case "leftarrow":
		return KeyLeftArrow, nil
	case "right":
		return KeyRightArrow, nil
	case "rightarrow":
		return KeyRightArrow, nil
	case "down":
		return KeyDownArrow, nil
	case "downarrow":
		return KeyDownArrow, nil
	case "up":
		return KeyUpArrow, nil
	case "uparrow":
		return KeyUpArrow, nil
	default:
		return 0, ErrUnknownKey
	}
}

func (k Key) String() string {
	switch k {
	case KeyA:
		return "a"
	case KeyS:
		return "s"
	case KeyD:
		return "d"
	case KeyF:
		return "f"
	case KeyH:
		return "h"
	case KeyG:
		return "g"
	case KeyZ:
		return "z"
	case KeyX:
		return "x"
	case KeyC:
		return "c"
	case KeyV:
		return "v"
	case KeyB:
		return "b"
	case KeyQ:
		return "q"
	case KeyW:
		return "w"
	case KeyE:
		return "e"
	case KeyR:
		return "r"
	case KeyY:
		return "y"
	case KeyT:
		return "t"
	case Key1:
		return "1"
	case Key2:
		return "2"
	case Key3:
		return "3"
	case Key4:
		return "4"
	case Key6:
		return "6"
	case Key5:
		return "5"
	case KeyEqual:
		return "="
	case Key9:
		return "9"
	case Key7:
		return "7"
	case KeyMinus:
		return "-"
	case Key8:
		return "8"
	case Key0:
		return "0"
	case KeyRightBracket:
		return "]"
	case KeyO:
		return "o"
	case KeyU:
		return "u"
	case KeyLeftBracket:
		return "["
	case KeyI:
		return "i"
	case KeyP:
		return "p"
	case KeyL:
		return "l"
	case KeyJ:
		return "j"
	case KeyQuote:
		return "'"
	case KeyK:
		return "k"
	case KeySemicolon:
		return ";"
	case KeyBackslash:
		return "\\"
	case KeyComma:
		return ","
	case KeySlash:
		return "/"
	case KeyN:
		return "n"
	case KeyM:
		return "m"
	case KeyPeriod:
		return "."
	case KeyGrave:
		return "`"
	case KeyReturn:
		return "return"
	case KeyTab:
		return "tab"
	case KeySpace:
		return "space"
	case KeyDelete:
		return "delete"
	case KeyEscape:
		return "escape"
	case KeyCapsLock:
		return "capslock"
	case KeyFunction:
		return "function"
	case KeyF17:
		return "f17"
	case KeyVolumeUp:
		return "volumeup"
	case KeyVolumeDown:
		return "volumedown"
	case KeyMute:
		return "mute"
	case KeyF18:
		return "f18"
	case KeyF19:
		return "f19"
	case KeyF20:
		return "f20"
	case KeyF5:
		return "f5"
	case KeyF6:
		return "f6"
	case KeyF7:
		return "f7"
	case KeyF3:
		return "f3"
	case KeyF8:
		return "f8"
	case KeyF9:
		return "f9"
	case KeyF11:
		return "f11"
	case KeyF13:
		return "f13"
	case KeyF16:
		return "f16"
	case KeyF14:
		return "f14"
	case KeyF10:
		return "f10"
	case KeyContextualMenu:
		return "contextualmenu"
	case KeyF12:
		return "f12"
	case KeyF15:
		return "f15"
	case KeyHelp:
		return "help"
	case KeyHome:
		return "home"
	case KeyPageUp:
		return "pageup"
	case KeyForwardDelete:
		return "forwarddelete"
	case KeyF4:
		return "f4"
	case KeyEnd:
		return "end"
	case KeyF2:
		return "f2"
	case KeyPageDown:
		return "pagedown"
	case KeyF1:
		return "f1"
	case KeyLeftArrow:
		return "left"
	case KeyRightArrow:
		return "right"
	case KeyDownArrow:
		return "down"
	case KeyUpArrow:
		return "up"
	default:
		return "unknown"
	}
}

func (k *Key) UnmarshalText(text []byte) error {
	key, err := ParseKey(text)
	*k = key
	return err
}
