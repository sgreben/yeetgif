package gifcmd

import (
	"encoding/json"
	"fmt"
)

// JSON is a `flag.Value` for JSON arguments.
type JSON struct {
	Value interface{}
	Text  string
}

// Set is flag.Value.Set
func (fv *JSON) Set(v string) error {
	fv.Text = v
	return json.Unmarshal([]byte(v), fv.Value)
}

func (fv *JSON) String() string {
	return fmt.Sprint(fv.Value)
}
