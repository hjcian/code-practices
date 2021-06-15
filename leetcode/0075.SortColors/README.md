# Question
https://leetcode.com/problems/sort-colors/

# Thoughts

剛好題目三個數字，小中大

那麼就就兩個 pointers: `zeroIdx`, `twoIdx` 指向目前「小的位置」與「大的位置」，起始分別為 `-1` 與 `len(nums)`

走訪元素時，檢查目前元素是 0 還是 2
- 0 的話，`zeroIdx++`，並且與之互換 `swap(nums[i], nums[zeroIdx])`，然後 `i++` 繼續走
- 2 的話，`twoIdx--`，並且與之互換 `swap(nums[i], nums[twoIdx])`，此時 `i` 不用 `++`，因為互換過來的元素還沒檢查

故考慮以下陣列（`zeroIdx`, `twoIdx` 分別以 `a` 與 `b` 簡稱）迭代展示思路：

> *iter 0*

```go
 a               b
-1 0 1 2 3 4 5 6 7
  [2,0,2,0,1,1,0]
   i
```

> *iter 1*

檢查 nums[i] 為 2

`b--`
```go
 a             b
-1 0 1 2 3 4 5 6 7
  [2,0,2,0,1,1,0]
   i
```

`swap(nums[i], nums[b])`
```go
 a             b
-1 0 1 2 3 4 5 6 7
  [0,0,2,0,1,1,2]
   i
```

i 不用 ++

> *iter 2*

檢查 nums[i] 為 0

`a++`
```go
   a           b
-1 0 1 2 3 4 5 6 7
  [0,0,2,0,1,1,2]
   i
```

`swap(nums[i], nums[a])`
```go
   a           b
-1 0 1 2 3 4 5 6 7
  [0,0,2,0,1,1,2]
   i
```

`i++`
```go
   a           b
-1 0 1 2 3 4 5 6 7
  [0,0,2,0,1,1,2]
     i
```

> *iter3*

以此類推，迴圈終止條件為 `i < b`