package gifcmd

import (
	"fmt"
	"strconv"
)

// Float is a `flag.Value` for a float argument.
// The `BitSize` field is used for parsing when set.
type Float struct {
	Value float64
	Text  string
}

// Set is flag.Value.Set
func (fv *Float) Set(v string) error {
	bitSize := 64
	n, err := strconv.ParseFloat(v, bitSize)
	if err == nil {
		fv.Value = n
		fv.Text = v
	}
	return err
}

func (fv *Float) String() string {
	return fmt.Sprint(fv.Value)
}
