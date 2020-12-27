package footballfanids

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_footballFanIds(t *testing.T) {
	type args struct {
		root string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test1",
			args{"./test1/files"},
			"./test1/want",
		},
		{
			"Test2",
			args{"./test2/files"},
			"./test2/want",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, err := ioutil.ReadFile(tt.want)
			require.NoError(t, err)

			expected := strings.TrimSpace(string(want))
			if got := strings.TrimSpace(
				footballFanIds(tt.args.root)); got != expected {
				require.Equal(t, expected, got)
			}
		})
	}
}
