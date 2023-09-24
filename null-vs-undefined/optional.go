package main

import (
	"encoding/json"
	"fmt"
)

type ParsedJSON struct {
	Name Optional[string] `json:"name"`
}

type Optional[T any] struct {
	Defined bool
	Value   *T
}

// UnmarshalJSON is implemented by deferring to the wrapped type (T).
// It will be called only if the value is defined in the  JSON payload.
func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	o.Defined = true
	return json.Unmarshal(data, &o.Value)
}

func main() {
	dataWithNull := ParsedJSON{}

	err := json.Unmarshal([]byte("{\"name\": null}"), &dataWithNull)
	if err != nil {
		fmt.Printf("Failed on property null %v", err)
	}

	dataWithValue := ParsedJSON{}
	err = json.Unmarshal([]byte("{\"name\": \"Test\"}"), &dataWithValue)
	if err != nil {
		fmt.Printf("Failed on property null %v", err)
	}

	dataWithUndefined := ParsedJSON{}
	err = json.Unmarshal([]byte("{}"), &dataWithUndefined)
	if err != nil {
		fmt.Printf("Failed on property undefined %v", err)
	}

	fmt.Printf("%#v,\n %#v,\n %#v", dataWithNull, dataWithValue, dataWithUndefined)
}

type NullableString string

func (b *NullableString) UnmarshalJSON(data []byte) error {
	var result string
	if string(data) == "null" {
		*b = ""
		return nil
	}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return err
	}
	c := NullableString(result)
	*b = c
	return nil
}
