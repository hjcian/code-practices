package analyzingipaddresses

import (
	"io/ioutil"
	"strings"
	"testing"
)

func Test_analyzingipaddresses(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name     string
		args     args
		wantPath string
	}{
		{
			"Test 1",
			args{"./test1/files"},
			"./test1/want",
		},
		{
			"Test 2",
			args{"./test2/files"},
			"./test2/want",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := analyzingipaddresses(tt.args.filepath)
			want, err := ioutil.ReadFile(tt.wantPath)
			check(err)

			if strings.TrimSpace(got) != strings.TrimSpace(string(want)) {
				t.Errorf("expect:\n%s\n-----", string(want))
				t.Errorf("got:\n%s\n-----", got)
			}

		})
	}
}
