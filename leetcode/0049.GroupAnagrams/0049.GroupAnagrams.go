package groupanagrams

import "slices"

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string, 0)

	for _, str := range strs {
		strByte := []byte(str)
		slices.Sort(strByte)
		groups[string(strByte)] = append(groups[string(strByte)], str)
	}

	ret := make([][]string, 0)
	for _, group := range groups {
		ret = append(ret, group)
	}
	return ret
}
