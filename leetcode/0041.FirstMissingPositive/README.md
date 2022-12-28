# Related
- 268

# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
- 參考解答，理解一遍

# Approach
<!-- Describe your approach to solving the problem. -->
分析以下 cases，了解值域：
```
case 1: [0, 1, 2], n=3, missing integer=3
case 2: [1, 2, 3], n=3, missing integer=4
case 3: [0, 1, 3], n=3, missing integer=2
case 4: [1, 2, 4], n=3, missing integer=3
```

得知，missing integer 落在 `1~n+1` 之間。

又題目限制必須使用 constant space，表示你必須想辦法利用原始 array，甚至是修改原始 array 的值。

又 space 有限制的 array 題目，常常需要你使用 element value 視為 array index 來做一些魔法。

這裡的魔法是，利用餘數特性，想辦法在指定 index 內做上標記。

```
// 3, 0, 1, 1 (i=0) ->3, 0, 1, 1+4=5
// 3, 0, 1, 5 (i=1) ->3+4=7, 0, 1, 5
// 7, 0, 1, 5 (i=2) ->7, 0+4=4, 1, 5
// 7, 4, 1, 5 (i=3) 用 5%n 找回 1 -> 7, 4+4=8, 1, 5
// 7, 8, 1, 5
```


# Complexity
- Time complexity: $$O(n)$$
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity: $$O(1)$$
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
func firstMissingPositive(nums []int) int {
	nums = append(nums, 0)
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] < 0 || nums[i] >= n {
			nums[i] = 0
		}
	}


	for i := 0; i < n; i++ {
		nums[nums[i]%n] += n
	}
	for i := 1; i < n; i++ {
		if nums[i] < n {
			return i
		}
	}
	return n
}
```