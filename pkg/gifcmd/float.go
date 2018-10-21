package gifcmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sgreben/piecewiselinear"
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

// FloatsCSV is a `flag.Value` for comma-separated `float` arguments.
// If `Accumulate` is set, the values of all instances of the flag are accumulated.
// The `BitSize` fields are used for parsing when set.
// The `Separator` field is used instead of the comma when set.
type FloatsCSV struct {
	Accumulate bool

	Values []float64
	Texts  []string
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *FloatsCSV) Help() string {
	return fmt.Sprintf("comma-separated list of 64-bit floats")
}

func (fv *FloatsCSV) PiecewiseLinear(min, max float64) func(float64) float64 {
	f := piecewiselinear.Function{Y: fv.Values}
	f.X = piecewiselinear.Span(min, max, len(f.Y))
	return f.At
}

// Set is flag.Value.Set
func (fv *FloatsCSV) Set(v string) error {
	bitSize := 64
	separator := ","
	if !fv.Accumulate {
		fv.Values = fv.Values[:0]
		fv.Texts = fv.Texts[:0]
	}
	parts := strings.Split(v, separator)
	for _, part := range parts {
		part = strings.TrimSpace(part)
		n, err := strconv.ParseFloat(part, bitSize)
		if err != nil {
			return err
		}
		fv.Values = append(fv.Values, n)
		fv.Texts = append(fv.Texts, part)
	}
	return nil
}

func (fv *FloatsCSV) String() string {
	return fmt.Sprint(fv.Values)
}
