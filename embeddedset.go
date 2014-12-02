package halgo

import "encoding/json"

// linkSet is represents a set of HAL links. Deserialisable from a single
// JSON hash, or a collection of links.
type embeddedSet []interface{}

func (l embeddedSet) MarshalJSON() ([]byte, error) {
	if len(l) == 1 {
		return json.Marshal(l[0])
	}

	other := make([]interface{}, len(l))
	copy(other, l)

	return json.Marshal(other)
}

func (l *embeddedSet) UnmarshalJSON(d []byte) error {
	var single interface{}
	err := json.Unmarshal(d, &single)
	if err == nil {
		*l = []interface{}{single}
		return nil
	}

	if _, ok := err.(*json.UnmarshalTypeError); !ok {
		return err
	}

	multiple := []interface{}{}
	err = json.Unmarshal(d, &multiple)

	if err == nil {
		*l = multiple
		return nil
	}

	return err
}
