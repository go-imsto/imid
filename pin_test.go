package imid

import (
	"fmt"
	"hash/crc64"
	"testing"

	"github.com/liut/baseconv"
)

// TestPin ...
func TestPin(t *testing.T) {

	crc64q := crc64.MakeTable(crc64.ISO)
	for _, cs := range []struct {
		name string
		n    int
	}{
		{"5B", 5},
		{"100B", 100},
		{"4KB", 4e3},
		{"20KB", 20e3},
		{"80KB", 80e3},
		{"10MB", 10e6},
	} {
		input := make([]byte, cs.n)
		for i := range input {
			input[i] = byte(i)
		}
		id := crc64.Checksum(input, crc64q)
		p := NewPin(id, EtPNG)
		// b := p.ID.Bytes()
		// t.Logf("id len %d", len(b))
		b := p.Bytes()
		lid, err := baseconv.Convert(fmt.Sprintf("%x", b), 16, 36)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%4s id: %20d %14x %14s \t%s\t%x,%s", cs.name, p.ID, p.ID.Bytes(), p.ID, p.Path(), b, lid)
		s := p.String()
		p2, err := ParsePin(s)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("parsed OK %q", p2)
		}

		// dt0km71q2c0rc dq0jk71n2c0oc
	}
}

// TestExt ...
func TestExt(t *testing.T) {
	for _, cs := range []struct {
		name string
		ext  Ext
	}{
		{"a.png", EtPNG},
		{"jpeg", EtJPEG},
	} {
		ext := ParseExt(cs.name)
		t.Logf("%8s %4s === %4s, %v", cs.name, ext, cs.ext, ext == cs.ext)
	}
}
