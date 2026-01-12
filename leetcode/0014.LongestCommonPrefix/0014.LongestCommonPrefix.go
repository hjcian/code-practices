package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	end := 0
	next := true
	for next {
		for i := 0; i < len(strs)-1; i++ {

			if end == len(strs[i]) || end == len(strs[i+1]) || strs[i][end] != strs[i+1][end] {
				next = false
				end--
				break
			}
		}
		if next {
			end++
		}
	}
	return strs[0][0 : end+1]
}
