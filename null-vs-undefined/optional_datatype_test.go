package datatype_test

import (
	"encoding/json"
	"testing"

	"gotest.tools/assert"
)

type TestJSONPayload struct {
	Name datatype.Optional[string] `json:"name"`
}

func TestOptionalDatatypeStates(t *testing.T) {

	t.Run("should marshal an explicitly null field correctly", func(t *testing.T) {
		testPayload := TestJSONPayload{}
		payload := `
			{
				"name": null
			}
		`
		err := json.Unmarshal([]byte(payload), &testPayload)
		assert.NilError(t, err)

		assert.Equal(t, testPayload.Name.Defined, true, "name is defined in the payload")
		assert.Equal(t, true, testPayload.Name.Value == nil, "name was explicitly set to nil (we use equal comparison since the types are different, nil *string vs untyped nil)")
	})

	t.Run("should correctly marshal a a field with value", func(t *testing.T) {
		testPayload := TestJSONPayload{}
		payload := `
			{
				"name": "some value"
			}
		`
		err := json.Unmarshal([]byte(payload), &testPayload)
		assert.NilError(t, err)

		assert.Equal(t, testPayload.Name.Defined, true, "name is defined in the payload")
		assert.Equal(t, "some value", *testPayload.Name.Value, "Value has been marshalled correctly")
	})

	t.Run("should identify when the property is ommitted", func(t *testing.T) {
		testPayload := TestJSONPayload{}
		payload := `
			{
				
			}
		`
		err := json.Unmarshal([]byte(payload), &testPayload)
		assert.NilError(t, err)

		assert.Equal(t, testPayload.Name.Defined, false, "name is NOT defined in the payload")
		assert.Equal(t, true, testPayload.Name.Value == nil, "name is ommitted in the payload, so it has no value")
	})
}

func TestSettingValues(t *testing.T) {
	t.Run("should set the value and keep the 'Defined' state in sync", func(t *testing.T) {
		testPayload := TestJSONPayload{}

		assert.Equal(t, false, testPayload.Name.Defined, "name field is not defined by default")
		assert.Equal(t, true, testPayload.Name.Value == nil, "name field value is nil")

		name := "Batman"
		testPayload.Name.Set(&name)

		assert.Equal(t, true, testPayload.Name.Defined, "name field is now defined")
		assert.Equal(t, name, *testPayload.Name.Value, "name field value was updated")
	})

	t.Run("should also work if we explicitly set a value to nil", func(t *testing.T) {
		testPayload := TestJSONPayload{}

		assert.Equal(t, false, testPayload.Name.Defined, "name field is not defined by default")
		assert.Equal(t, true, testPayload.Name.Value == nil, "name field value is nil")

		testPayload.Name.Set(nil)

		assert.Equal(t, true, testPayload.Name.Defined, "name field is now defined")
		assert.Equal(t, true, testPayload.Name.Value == nil, "name field value is nil")
	})
}

func TestMakingUndefined(t *testing.T) {
	t.Run("should reset the state of the field as 'undefined'", func(t *testing.T) {
		testPayload := TestJSONPayload{}

		name := "Batman"
		testPayload.Name.Set(&name)

		assert.Equal(t, true, testPayload.Name.Defined, "name field is defined")
		assert.Equal(t, name, *testPayload.Name.Value, "name field value is set")

		testPayload.Name.MakeUndefined()

		assert.Equal(t, false, testPayload.Name.Defined, "name field is now marked as undefined")
		assert.Equal(t, true, testPayload.Name.Value == nil, "name field value is nil")
	})

}
