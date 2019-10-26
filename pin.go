package imagid

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

const (
	binaryVersion byte = 1
)

// Pin ...
type Pin struct {
	ID  IID
	Ext Ext
}

// NewPin ...
func NewPin(id uint64, ext Ext) Pin {
	var p = Pin{ID: IID(id), Ext: Ext(ext)}
	return p
}

// Bytes ...
func (p Pin) Bytes() []byte {
	enc := append([]byte{}, p.ID.Bytes()...)
	return append(enc, p.Ext.Val(), binaryVersion)
}

// String ...
func (p Pin) String() string {
	return p.ID.String() + "." + p.Ext.String()
}

// Path ...
func (p Pin) Path() string {
	s := p.ID.String()
	return s[0:2] + "/" + s[2:] + "." + p.Ext.String()
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (p Pin) MarshalBinary() ([]byte, error) {
	return p.Bytes(), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (p *Pin) UnmarshalBinary(buf []byte) error {
	if len(buf) == 0 {
		return errors.New("Pin.UnmarshalBinary: no data")
	}

	if len(buf) != /*id*/ 8+ /*ext*/ 1+ /*version*/ 1 {
		return errors.New("Pin.UnmarshalBinary: invalid length")
	}

	if buf[len(buf)-1] != binaryVersion {
		return errors.New("Pin.UnmarshalBinary: unsupported version")
	}

	id := int64(buf[7]) | int64(buf[6])<<8 | int64(buf[5])<<16 | int64(buf[4])<<24 |
		int64(buf[3])<<32 | int64(buf[2])<<40 | int64(buf[1])<<48 | int64(buf[0])<<56

	*p = Pin{}
	p.ID = IID(uint64(id))
	p.Ext = Ext(buf[8])
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface.
func (p Pin) MarshalText() ([]byte, error) {
	b := []byte(p.String())
	return b, nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (p *Pin) UnmarshalText(data []byte) (err error) {
	*p = Pin{}
	p, err = ParsePin(string(data))
	return
}

// ParsePin ...
func ParsePin(s string) (p *Pin, err error) {
	arr := strings.Split(s, ".")
	if len(arr) < 2 {
		return nil, errors.New("invalid pin data: '" + s + "'")
	}
	var id IID
	id, err = ParseID(arr[0])
	if err != nil {
		return
	}
	p = &Pin{}
	p.ID = id
	p.Ext = ParseExt(arr[1])
	return
}

// Scan implements of database/sql.Scanner
func (p *Pin) Scan(src interface{}) (err error) {
	switch s := src.(type) {
	case string:
		return p.UnmarshalText([]byte(s))
	case []byte:
		return p.UnmarshalText(s)
	}
	return fmt.Errorf("'%v' is invalid Pin", src)
}

// Value implements of database/sql/driver.Valuer
func (p Pin) Value() (driver.Value, error) {
	return p.String(), nil
}
