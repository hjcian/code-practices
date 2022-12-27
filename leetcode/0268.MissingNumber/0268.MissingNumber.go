package missingnumber

func missingNumber(arr []int) int {
	hashMap := make(map[int]struct{})

	for i := range arr {
		hashMap[arr[i]] = struct{}{}
	}

	for i := range arr {
		_, ok := hashMap[i]
		if !ok {
			return i
		}
	}
	return len(arr)
}

func missingNumber_Sum_formula(arr []int) int {
	formulaSum := (1 + len(arr)) * len(arr) / 2
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return formulaSum - sum
}

func missingNumber_bit_operation(arr []int) int {
	res := len(arr)
	for i := range arr {
		res ^= arr[i]
		res ^= i
	}
	return res
}

// return fmt.Sprintf("%016b", p)
