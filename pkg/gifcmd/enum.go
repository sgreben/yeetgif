package gifcmd

import (
	"fmt"
	"strings"
)

// Enum is a `flag.Value` for one-of-a-fixed-set string arguments.
// The value of the `Choices` field defines the valid choices.
// If `CaseSensitive` is set to `true` (default `false`), the comparison is case-sensitive.
type Enum struct {
	Choices []string

	Value string
	Text  string
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *Enum) Help() string {
	return fmt.Sprintf("one of %v", fv.Choices)
}

// Set is flag.Value.Set
func (fv *Enum) Set(v string) error {
	fv.Text = v
	for _, c := range fv.Choices {
		if strings.EqualFold(c, v) {
			fv.Value = c
			return nil
		}
	}
	return fmt.Errorf(`"%s" must be one of [%s]`, v, strings.Join(fv.Choices, " "))
}

func (fv *Enum) String() string {
	return fv.Value
}
