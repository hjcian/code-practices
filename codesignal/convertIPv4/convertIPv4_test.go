package convertipv4

import (
	"encoding/binary"
	"fmt"
	"net"
	"reflect"
	"strconv"
	s "strings"
	"testing"
)

func solution(a string) string {
	var z []int
	for _, p := range s.Split(a, ".") {
		j, f := strconv.Atoi(p)
		if f != nil || j>>8 != 0 {
			return ""
		}
		z = append(z, j)
	}
	if len(z) != 4 {
		return ""
	}

	l, m := z[2]<<8|z[3], z[0]<<8|z[1]
	return fmt.Sprintf("0:0:0:0:0:ffff:%x:%x", m, l)
}

func Test_convertipv4(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []string
	}{
		{"Example", "216.58.200.14", []string{"3627730958", "0:0:0:0:0:ffff:d83a:c80e"}},
		{"Test x0", "0.0.0.0", []string{"0", solution("0.0.0.0")}},
		{"Test x1", "0.0.0.1", []string{"1", solution("0.0.0.1")}},
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
	t.Run("Loop all IPv4", func(t *testing.T) {
		int2ip := func(nn uint32) net.IP {
			ip := make(net.IP, 4)
			binary.BigEndian.PutUint32(ip, nn)
			return ip
		}

		for i := 0; i < 1<<32-1; i++ {
			ip := int2ip(uint32(i)).String()
			got, sol := ipv4Toipv6(ip), solution(ip)
			if got != sol {
				t.Errorf("[%v] Expect '%v', got '%v' \n", ip, sol, got)
				t.FailNow()
			}
		}
	})
}
