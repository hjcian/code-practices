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
		{"Example", "216.58.200.14", []string{"3627730958", "0:0:0:0:0:ffff:d83a:c80e"}},
		{"Test 2", "2.2.2", []string{}},
		{"Guess1", "256.255.255.255", []string{}},
		{"Guess2", "-1.255.255.255", []string{}},
		{"Guess3", ".255.255.255", []string{}},
	}
	for _, test := range tests {
		got := convertIPv4(test.input)
		if !reflect.DeepEqual(got, test.expect) {
			t.Errorf("[case; %v] Expect '%v', got '%v' \n", test.name, test.expect, got)
		}
	}
}
