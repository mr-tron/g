package g

import (
	"reflect"
	"testing"
)

func TestUniqBool(t *testing.T) {
	input := []bool{}
	output := UniqBool(input)
	if !reflect.DeepEqual(input, output) {
		t.Fatal("invalid result")
	}
	input = []bool{true, true, true}
	output = UniqBool(input)
	if !reflect.DeepEqual(output, []bool{true}) {
		t.Fatal("invalid result")
	}
	input = []bool{true, true, false, true}
	output = UniqBool(input)
	if !reflect.DeepEqual(output, []bool{true, false}) {
		t.Fatal("invalid result")
	}
}

func TestUniqString(t *testing.T) {
	input := []string{}
	output := UniqString(input)
	if !reflect.DeepEqual(input, output) {
		t.Fatal("invalid result")
	}
	input = []string{"1", "2", "ololo", "1", "ololo"}
	output = UniqString(input)
	if !reflect.DeepEqual(output, []string{"1", "2", "ololo"}) {
		t.Fatal("invalid result")
	}
}
