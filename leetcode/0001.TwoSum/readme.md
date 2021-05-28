# Question
https://leetcode.com/problems/two-sum/

陣列裡有兩個數字加起來等於目標值
且假設陣列裡只會有一組唯一的答案

# Thoughts
**原思路**

走一遍陣列 (*O(n)*)，會需要「查找」該陣列其他元素，找到兩者相加等於目標值的元素索引值

暴力解的思路是用 *O(n)* 的方式「查找」，造成複雜度變成 *O(n<sup>2</sup>)*：
```
for i in [0, len(array)):
    for j in [i+1, len(array)):
        if array[i]+array[j] == target:
            return (i, j)
```

但如果是先建立一個 lookup table ，讓查詢變成 *O(1)*，就可以讓複雜度變成:
- 建立 lookup table *O(n)*
- +走一遍陣列 *O(n)*
  - +每一個元素會查找一次，共 n 次 *n * O(1)*
- =故總複雜度變成 *2O(n) + nO(1)*

**最速解賞析**

將建立 lookup table 的操作放在走一遍陣列的 for loop 內，當查不到的時候就將當前的元素放到 lookup table 中
有機會 early return 增加效率

```go
func twoSum(nums []int, target int) []int {
    seen := map[int]int{}
    for i := 0; i < len(nums); i++ {
        v, ok := seen[target-nums[i]]
        if ok {
            return []int{v, i}
        }

        seen[nums[i]] = i
    }
    return []int{}
}
```