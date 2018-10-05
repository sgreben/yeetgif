package gifmeta

import (
	"bytes"
	"fmt"
	"io"
)

type extensionType byte
type ExtensionType extensionType

// Extension types
const (
	Comment        ExtensionType = 0xFE
	PlainText      ExtensionType = 0x01
	Application    ExtensionType = 0xFF
	GraphicControl ExtensionType = 0xF9
)

type Extension struct {
	Type   ExtensionType
	Blocks [][]byte
}

func (e *Extension) String() string {
	buf := &bytes.Buffer{}
	for _, block := range e.Blocks {
		buf.Write(block)
	}
	return buf.String()
}

func Blocks(b []byte) (out [][]byte) {
	for i := 0; i < len(b); i += 255 {
		j := i + 255
		if j > len(b) {
			j = len(b)
		}
		out = append(out, b[i:j])
	}
	return
}

func (e Extension) Write(w io.Writer) error {
	buf := &bytes.Buffer{}
	buf.WriteByte(byte(e.Type))
	for i, block := range e.Blocks {
		n := len(block)
		if n < 1 {
			return fmt.Errorf("comment block #%d too short (%d bytes): %q", i, n, string(block))
		}
		if n > 255 {
			return fmt.Errorf("comment block #%d too long (%d bytes): %q", i, n, string(block))
		}
		buf.WriteByte(uint8(n))
		buf.Write(block)
	}
	buf.WriteByte(0x00)
	_, err := io.Copy(w, buf)
	return err
}

func readBlock(r io.Reader) ([]byte, error) {
	lengthByte := []byte{0x00}
	_, err := io.ReadFull(r, lengthByte)
	if err != nil {
		return nil, fmt.Errorf("read length byte: %v", err)
	}
	n := uint8(lengthByte[0])
	if n == 0 {
		return nil, nil
	}
	buf := make([]byte, n)
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return nil, fmt.Errorf("read body: %v", err)
	}
	return buf, nil
}

func (e *Extension) readBlocks(r io.Reader) error {
	for {
		block, err := readBlock(r)
		if err != nil {
			return fmt.Errorf("read gif extension block %d: %v", len(e.Blocks), err)
		}
		if len(block) == 0 {
			return nil
		}
		e.Blocks = append(e.Blocks, block)
	}
}
func (e *Extension) Read(r io.Reader) error {
	singleByte := []byte{0x00}
	_, err := io.ReadFull(r, singleByte)
	if err != nil {
		return fmt.Errorf("read gif extension type: %v", err)
	}
	e.Type = ExtensionType(singleByte[0])
	switch e.Type {
	case GraphicControl:
		block := make([]byte, 6)
		_, err := io.ReadFull(r, block)
		if err != nil {
			return fmt.Errorf("read gif graphic control block: %v", err)
		}
		e.Blocks = append(e.Blocks, block)
	case Comment, Application:
		return e.readBlocks(r)
	case PlainText:
		block := make([]byte, 13)
		_, err := io.ReadFull(r, block)
		if err != nil {
			return fmt.Errorf("read gif plaintext block: %v", err)
		}
		e.Blocks = append(e.Blocks, block)
		return e.readBlocks(r)
	}
	return nil
}
