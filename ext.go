package imid

import (
	"strings"
)

// Ext ...
type Ext byte

// consts
const (
	EtNone Ext = iota
	EtGIF
	EtJPEG
	EtPNG
	EtWebP
)

// String for fmt.Stringer
func (z Ext) String() string {
	switch z {
	case EtGIF:
		return "gif"
	case EtJPEG:
		return "jpeg"
	case EtPNG:
		return "png"
	case EtWebP:
		return "webp"
	}
	return "unknown"
}

// Val ...
func (z Ext) Val() byte {
	return byte(z)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (z Ext) MarshalText() ([]byte, error) {
	b := []byte(z.String())
	return b, nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (z *Ext) UnmarshalText(data []byte) error {
	*z = ParseExt(string(data))
	return nil
}

// ParseExt ...
func ParseExt(s string) Ext {
	if pos := strings.LastIndex(s, "."); pos != -1 && pos < len(s) {
		s = s[pos+1:]
	}
	switch s {
	case "gif":
		return EtGIF
	case "jpeg", "jpg":
		return EtJPEG
	case "png":
		return EtPNG
	case "webp":
		return EtWebP
	}
	return EtNone
}
