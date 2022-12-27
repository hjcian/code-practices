posted on: https://leetcode.com/problems/missing-number/solutions/2958461/golang-bit-operation-solution-explained-in-chinese/

# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
- 看別人的解說，並複習基本的 XOR 結合律、交換律、恆等律、歸零律、自反，了解原理

# Approach
<!-- Describe your approach to solving the problem. -->
- 在學習完基本的 XOR 規則後，再細細看題目限制：
  - array 長度若為 n，則 array 內的元素最大值為 n
  - array 內的元素不重複
- 假設有個數字 M，`M^n^n = M`，且無論 XOR 運算的順序皆會如此，故
- 接著思考：在題目的限制下你可以想像若 given array 為 `[0 1 2 3 4]`，則想要找到的缺值為 5。又若有另一條 array 為 `[0 1 2 3 4 5]`，則你可以拿 `0` 與兩條 array 每一個元素都做一次 XOR，將運算操作攤平會如下所示：
    ```
    0^(0^0)^(1^1)^(2^2)^(3^3)^(4^4)^5
    = 0^(0)^(0)^(0)^(0)^(0)^5
    = 5
    ```
- 所以，我們就直接對 given array 做一次 interation，將 `0` 與元素值做 XOR，再對 index 做一次 XOR。 iteratate 完了之後，因為 index 只會做到 `n-1`，故需要再對 `n` 補做一個 XOR
- 最終的結果就是 missing number
- 又根據交換律與結合律，可以將最後對 `n` 做 XOR 的操作拿到最前面，與初始值 `0` 做 XOR。因為 `0^len(arr)=len(arr)`，所以 refactored 版本就直接將 `res := len(arr)`


# Complexity
- Time complexity: $$O(n)$$
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity: $$O(1)$$
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```
func missingNumber(arr []int) int {
    res := 0
	for i := range arr {
		res ^= arr[i]
		res ^= i
	}
    res ^= len(arr)
	return res
}
```

refactored:
```
func missingNumber(arr []int) int {
    res := len(arr)
	for i := range arr {
		res ^= arr[i]
		res ^= i
	}
	return res
}
```