// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package job

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseRUNNER_TYPEFromString Parse RUNNER_TYPE from string
func ParseRUNNER_TYPEFromString(str string) (RUNNER_TYPE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := RUNNER_TYPE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown RUNNER_TYPE: %s", str)
	}

	return RUNNER_TYPE(v), nil
}

// Equal type compare
func (t RUNNER_TYPE) Equal(target RUNNER_TYPE) bool {
	return t == target
}

// IsIn todo
func (t RUNNER_TYPE) IsIn(targets ...RUNNER_TYPE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t RUNNER_TYPE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *RUNNER_TYPE) UnmarshalJSON(b []byte) error {
	ins, err := ParseRUNNER_TYPEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseRUN_MODEFromString Parse RUN_MODE from string
func ParseRUN_MODEFromString(str string) (RUN_MODE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := RUN_MODE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown RUN_MODE: %s", str)
	}

	return RUN_MODE(v), nil
}

// Equal type compare
func (t RUN_MODE) Equal(target RUN_MODE) bool {
	return t == target
}

// IsIn todo
func (t RUN_MODE) IsIn(targets ...RUN_MODE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t RUN_MODE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *RUN_MODE) UnmarshalJSON(b []byte) error {
	ins, err := ParseRUN_MODEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParsePARAM_VALUE_TYPEFromString Parse PARAM_VALUE_TYPE from string
func ParsePARAM_VALUE_TYPEFromString(str string) (PARAM_VALUE_TYPE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := PARAM_VALUE_TYPE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown PARAM_VALUE_TYPE: %s", str)
	}

	return PARAM_VALUE_TYPE(v), nil
}

// Equal type compare
func (t PARAM_VALUE_TYPE) Equal(target PARAM_VALUE_TYPE) bool {
	return t == target
}

// IsIn todo
func (t PARAM_VALUE_TYPE) IsIn(targets ...PARAM_VALUE_TYPE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t PARAM_VALUE_TYPE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *PARAM_VALUE_TYPE) UnmarshalJSON(b []byte) error {
	ins, err := ParsePARAM_VALUE_TYPEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
