# Question

# Thoughts
- https://leetcode.com/problems/maximum-product-subarray/solutions/203013/c-o-n-time-o-1-space-solution-with-explanation/
- killer 👉 https://www.youtube.com/watch?v=lXVy6YWFcRM
- 前一步的最大值與最小值分別與 nums[i] 相乘，都有機會貢獻出全域最大值

```
   n  max  min best
  -2   -2   -2   -1
  -3    6   -3    6
  -1    3   -6    6
   2    6  -12    6
  -4   48  -24   48
  -3   72 -144   72
  -1  144  -72  144
```