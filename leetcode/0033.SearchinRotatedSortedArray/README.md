# Question

https://leetcode.com/problems/search-in-rotated-sorted-array/

There is an integer array `nums` sorted in ascending order (with distinct values).

此陣列會選定一個位置，並做 rotate，例如：
`[0,1,2,4,5,6,7]` 會變轉成
`[4,5,6,7,0,1,2]`

限制：
- 陣列長度介於 1~5000
- 但元素值域介於 -10000~10000
- 陣列元素值皆不重複
- 陣列保證一定經過 rotate

# Thoughts

## Python

可能會看到以下陣列
- `[51 ... 80 ... 100 ... 150 ... 50]`，其中
  - 元素最大值假設已知=150
  - 陣列總長<150
  - 目標值為 120
- 此時至少有個規則是，以中央伍為準，左邊或右邊一定有一邊滿足 left < mid 或 mid < right 的條件
- 此時在滿足的那一邊，檢查 target 是否在該邊的範圍內，若在的話則用這半邊繼續查找；否則用另一半邊
- pseudo code (還有些 edge case 需檢查，避免 index out of range 之類的):
    ```shell
    while left < right:
        if mid == target:
            return mid

        if left < mid:
            if left < target < mid:
                right = mid
            else:
                left = mid+1
        else if mid < right:
            if mid < target < right:
                left = mid+1
            else:
                right = mid
    ```