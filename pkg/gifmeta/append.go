package gifmeta

import (
	"bytes"
	"io"
)

func Append(w io.Writer, r io.Reader, append ...Extension) error {
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, r)
	io.CopyN(w, buf, int64(buf.Len()-1))
	for _, e := range append {
		if _, err := w.Write(extensionToken); err != nil {
			return err
		}
		if err := e.Write(w); err != nil {
			return err
		}
	}
	_, err := w.Write([]byte{trailer})
	return err
}
