package gifcmd

import (
	"fmt"
	"sort"
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

// EnumSetCSV is a `flag.Value` for comma-separated enum arguments.
// Only distinct values are returned.
type EnumSetCSV struct {
	Choices []string

	Value map[string]bool
	Texts []string
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *EnumSetCSV) Help() string {
	separator := ","
	return fmt.Sprintf("%q-separated list of values from %v", separator, fv.Choices)
}

// Values returns a string slice of specified values.
func (fv *EnumSetCSV) Values() (out []string) {
	for v := range fv.Value {
		out = append(out, v)
	}
	sort.Strings(out)
	return
}

// Set is flag.Value.Set
func (fv *EnumSetCSV) Set(v string) error {
	separator := ","
	if fv.Value == nil {
		fv.Value = make(map[string]bool)
	}
	parts := strings.Split(v, separator)
	for _, part := range parts {
		part = strings.TrimSpace(part)
		var ok bool
		var value string
		for _, c := range fv.Choices {
			if strings.EqualFold(c, part) {
				value = c
				ok = true
				break
			}
		}
		if !ok {
			return fmt.Errorf(`"%s" must be one of [%s]`, v, strings.Join(fv.Choices, " "))
		}
		fv.Value[value] = true
		fv.Texts = append(fv.Texts, part)
	}
	return nil
}

func (fv *EnumSetCSV) String() string {
	return strings.Join(fv.Values(), ",")
}
