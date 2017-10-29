package gotils

import (
	"testing"
)

type dummyMarshallable struct {
	name string
}

func newDummyMarshallable(name string) dummyMarshallable {
	return dummyMarshallable{name}
}

func (d *dummyMarshallable) jsonSerialize() map[string]interface{} {
	props := make(map[string]interface{})
	props["name"] = d.name
	return props
}

func TestJsonWrapperEncode(t *testing.T) {
	dummy := newDummyMarshallable("test")
	wrapper := newJsonWrapper()

	encoded, err := wrapper.Encode(&dummy)
	if err != nil {
		t.Error("got error: ", err)
	}

	if encoded != "{\"name\":\"test\"}" {
		t.Error("Encoded message is not valid, expected: {\"name\":\"test\"}, got: ", encoded)
	}
}

func TestJsonWrapperDecode(t *testing.T) {
	wrapper := newJsonWrapper()
	testMessage := "{\"name\":\"test\"}"

	props := make(map[string]interface{})
	props["name"] = "test"
	decoded, err := wrapper.Decode(testMessage)

	if err != nil {
		t.Error("got error: ", err)
	}

	if decoded["name"] != props["name"] {
		t.Error("expected name with value: test, got: ", decoded["name"])
	}
}
