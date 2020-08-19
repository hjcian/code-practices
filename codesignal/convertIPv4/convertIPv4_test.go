package convertipv4

import (
	"reflect"
	"testing"
)

func Test_convertipv4(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []string
	}{
		{"Example", "216.58.200.14", []string{"3627730958", "not"}},
	}
	for _, test := range tests {
		got := convertIPv4(test.input)
		if !reflect.DeepEqual(got, test.expect) {
			t.Errorf("[case; %v] Expect '%v', got '%v' \n", test.name, test.expect, got)
		}
	}
}
