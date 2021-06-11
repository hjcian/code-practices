# Question

https://leetcode.com/problems/search-in-rotated-sorted-array/

There is an integer array `nums` sorted in ascending order (with distinct values).

此陣列會選定一個位置，並做 rotate，例如：
`[0,1,2,4,5,6,7]` 會變轉成
`[4,5,6,7,0,1,2]`

限制：
陣列長度介於 1~5000
但元素值域介於 -10000~10000
陣列元素值皆不重複
陣列保證一定經過 rotate

# Thoughts

## Python

可能會看到以下陣列
元素最大值假設已知=150
陣列總長<150
目標值為 120
[51 ... 80 ... 100 ... 150 ... 50]