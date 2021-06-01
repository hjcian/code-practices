# Searching algorithms

## linear search
- 沒什麼好說的，就是走過一遍找到它

## binary search
- 前提是數列已排序好
- divide and conquer 的思維
- 並且直接猜中間那個數值是否為目標
- 若非，則比大小決定下一次要看左半邊還右半邊
- 直到找到為止

## interpolation search (內插搜尋)
- 是 binary search 的變體
- 適用情境是你的資料除了已排序好，元素值還存在著某種已知的線性關係 (linear, quadratic, cubic, ...etc.)
- 此時你在猜 target 時就可以不用 middle position，而是使用那條線性關係的方程式來猜
- 此調整有機會在你的 date set 中加速查找效率

故若假設資料呈現**直線關係**，也就是資料為 `uniformly distributed`，故三點間的斜率相等，可導出 `guessed x`：
```
x = index
y = value
left = (x1, y1)
target = (x, target)
right = (x2, y2)

guessed x := x1 + (x2-x1)*(target-y1)/(y2-y1)
```

complexity:
完美均勻分布，可達到 O(1)，內插一次就得到
平均來說可達到 O(log(log(n)))
分佈根本不均勻，最差的情況則是 O(n)


ref:
- [Understanding The Complexity Of Interpolation Search](https://www.cadmo.ethz.ch/education/lectures/HS18/SAADS/reports/17.pdf)
- [Data Structure - Interpolation Search](https://www.tutorialspoint.com/data_structures_algorithms/interpolation_search_algorithm.htm)
- [插補搜尋(Interpolation Search)演算法，運用資料近似線來輔助搜尋的演算法](https://magiclen.org/interpolation-search)