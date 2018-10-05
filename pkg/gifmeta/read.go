package gifmeta

import (
	"bytes"
	"fmt"
	"io"
)

func Read(r io.Reader, fs ...func(e *Extension) bool) ([]Extension, error) {
	var out []Extension
	if err := skipToken(r, header); err != nil {
		return out, fmt.Errorf("read gif header signature: %v", err)
	}
	if err := skipEitherToken(r, version87a, version89a); err != nil {
		return out, fmt.Errorf("read gif header version: %v", err)
	}
	logicalScreenDescriptor := make([]byte, 7)
	if _, err := io.ReadFull(r, logicalScreenDescriptor); err != nil {
		return out, fmt.Errorf("read gif header: %v", err)
	}
	globalFlags := int(logicalScreenDescriptor[4])
	colorTableField := 1 << 7
	colorTableBitsMask := 3
	if globalFlags&colorTableField != 0 {
		n := 1 << (1 + uint(globalFlags&colorTableBitsMask))
		if _, err := io.ReadFull(r, discard[:3*n]); err != nil {
			return out, fmt.Errorf("read global color table: %v", err)
		}
	}
	singleByte := []byte{0x00}
	i := 0
	for {
		_, err := r.Read(singleByte)
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return out, fmt.Errorf("read gif section type (section %d): %v", i, err)
		}
		switch singleByte[0] {
		case extension:
			e := &Extension{}
			err := e.Read(r)
			if err != nil {
				return out, fmt.Errorf("read gif extension (section %d): %v", i, err)
			}
			ok := true
			for _, f := range fs {
				ok = ok && f(e)
				if !ok {
					break
				}
			}
			if ok {
				out = append(out, *e)
			}
		case imageDescriptor:
			if err := skipImageDescriptor(r); err != nil {
				return out, fmt.Errorf("read gif image descriptor (section %d): %v", i, err)
			}
		case trailer:
			return out, nil
		}
		i++
	}

}

func skipToken(r io.Reader, token []byte) error {
	buf := make([]byte, len(token))
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return err
	}
	if !bytes.Equal(token, buf) {
		return fmt.Errorf("expected: %s, actual: %s", token, buf)
	}
	return nil
}

func skipEitherToken(r io.Reader, token1, token2 []byte) error {
	buf := make([]byte, len(token1))
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return err
	}
	ok := bytes.Equal(token1, buf) || bytes.Equal(token2, buf)
	if !ok {
		return fmt.Errorf("expected: %x or %x, actual: %x", token1, token2, buf)
	}
	return nil
}

func skipBlock(r io.Reader) (uint8, error) {
	lengthByte := []byte{0x00}
	_, err := io.ReadFull(r, lengthByte)
	if err != nil {
		return 0, fmt.Errorf("read length: %v", err)
	}
	n := uint8(lengthByte[0])
	if n == 0 {
		return 0, nil
	}
	_, err = io.ReadFull(r, discard[:n])
	if err != nil {
		return n, fmt.Errorf("read data (length %d): %v", n, err)
	}
	return n, nil
}

func skipImageDescriptorHeader(r io.Reader) error {
	header := make([]byte, 9)
	if _, err := io.ReadFull(r, header); err != nil {
		return err
	}
	imageFlags := int(header[8])
	colorTableFlag := 1 << 7
	colorTableBitsMask := 7
	if imageFlags&colorTableFlag != 0 {
		n := 1 << (1 + uint(imageFlags&colorTableBitsMask))
		if _, err := io.ReadFull(r, discard[:3*n]); err != nil {
			return fmt.Errorf("read color table: %v", err)
		}
	}
	return nil
}

func skipImageDescriptor(r io.Reader) error {
	if err := skipImageDescriptorHeader(r); err != nil {
		return fmt.Errorf("read image descriptor header: %v", err)
	}
	if _, err := io.ReadFull(r, discard[:1]); err != nil {
		return fmt.Errorf("read image descriptor pixel size: %v", err)
	}
	for {
		n, err := skipBlock(r)
		if err != nil {
			return fmt.Errorf("read image descriptor block: %v", err)
		}
		if n == 0 {
			return nil
		}
	}
}
