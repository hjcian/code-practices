# Sorting algorithms

## selection sort
基本來說，選擇排序只需要重複執行兩個步驟，分別是：
1. 找最小值
    - 從「未排序好的數字」中找到**最小值**
2. 丟到左邊
    - 把**最小值**丟到「未排序好的數字」的最左邊，把它標示成已排序好

complexity: O(n<sup>2</sup>)
```
traversal count: n + (n-1) + (n-2) + ... + 3 + 2 + 1
是梯形公式
[(1 + n) * n]/2 = (n^2 + n)/2 ~ O(n2)
```

ref: [初學者學演算法｜排序法入門：選擇排序與插入排序法](https://medium.com/appworks-school/%E5%88%9D%E5%AD%B8%E8%80%85%E5%AD%B8%E6%BC%94%E7%AE%97%E6%B3%95-%E6%8E%92%E5%BA%8F%E6%B3%95%E5%85%A5%E9%96%80-%E9%81%B8%E6%93%87%E6%8E%92%E5%BA%8F%E8%88%87%E6%8F%92%E5%85%A5%E6%8E%92%E5%BA%8F%E6%B3%95-23d4bc7085ff)
## insertion sort
簡單來說，插入排序法就是你玩撲克牌時用到的排序法。
1. 讀一個數字
    - 從「未排序好的數字」最左邊中讀取一個數
2. 插入合適位置
    - 把這個讀取的數與左邊的數值比大小，若左邊的數值比較大則兩者互換位置，直到左邊的數字大於這個讀取的數字

complexity: O(n<sup>2</sup>)

ref: [初學者學演算法｜排序法入門：選擇排序與插入排序法](https://medium.com/appworks-school/%E5%88%9D%E5%AD%B8%E8%80%85%E5%AD%B8%E6%BC%94%E7%AE%97%E6%B3%95-%E6%8E%92%E5%BA%8F%E6%B3%95%E5%85%A5%E9%96%80-%E9%81%B8%E6%93%87%E6%8E%92%E5%BA%8F%E8%88%87%E6%8F%92%E5%85%A5%E6%8E%92%E5%BA%8F%E6%B3%95-23d4bc7085ff)

## bubble sort
Bubble Sort 的方式是從陣列的最前面開始，一次比較陣列中兩兩相鄰的元素，然後根據大小將它們調換順序，大的移到後面
重複一次過程，此時應已確保最大的在最後面，故排除最右邊的元素

complexity: O(n<sup>2</sup>)
ref: [[演算法] 氣泡排序法（Bubble Sort）：利用兩兩元素交換位置達到排序](https://pjchender.blogspot.com/2017/09/bubble-sort.html)

## merge sort
是一個 divide and conquer 演算法
基本上透過 recursive 不斷將 array 二分，直到 array size==1 則直接 return
而最小二分後拿到的左邊與右邊，再用一個 `merge()` 方法來合併
在 `merge()` 中會額外創建 merged array 供回傳
再額外用變數儲存 left and right array current positions 來走過元素並比較較小的，放進 merged array

**time complexity analysis**
- divide 原始 array 至最小單元，需要的步驟就是 log(n) 了
- 接著從最小單元開始 merge，到底走訪的幾次 elements？
- 思考這個問題可以從最後的結果倒推
- 最後的 n 是由 2 個 n/2 組合而來。而最後的 n 要能夠組起來表示 2 個 n/2 都看過一次，所以是 n 次
- 再往下推， n/2 也是來自 2 個 n/4 組合而來，而要組合 n/2 也需要走過 n/2 次，故在這一層總共也需要 n 次 (因為 2 個 n/2)
- 而我們的 tree 多高？ log(n) 這麼高
- 所以總共就需要 nlog(n) 這麼多次的走訪

![https://www.khanacademy.org/computing/computer-science/algorithms/merge-sort/a/analysis-of-merge-sort](https://imgur.com/IDLRERH.png)

**is a stable sort**
也就是同樣的元素依舊維持在同樣的位置
需要在 merge 時注意，處理元素值相等時，以左邊的優先即可

complexity: O(nlog(n))

## quick sort


## heap sort