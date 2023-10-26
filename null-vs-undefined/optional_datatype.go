package datatype

import "encoding/json"

// Optional datatype allows for a tri state field, which is useful for PATCH merge operations.
// Incoming payloads may contain fields that can either be:
// a) set with a value that updates the current state's field
// b) ommitted, as to be ignored by the patch handler
// c) explicitely set to "nil", meaning that the value should be cleaned up/deleted from the current state's field
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

// Set is a utility function that helps setting the value while keeping the "Defined" state in sync
func (o *Optional[T]) Set(value *T) {
	o.Defined = true
	o.Value = value
}

// MakeUndefined will reset the field to its undefined state
func (o *Optional[T]) MakeUndefined() {
	o.Defined = false
	o.Value = nil
}
