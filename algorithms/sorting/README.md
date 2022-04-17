# Sorting algorithms
- [Sorting algorithms](#sorting-algorithms)
  - [Selection sort](#selection-sort)
  - [Insertion sort](#insertion-sort)
  - [Bubble sort](#bubble-sort)
  - [Merge sort](#merge-sort)
  - [Quick sort](#quick-sort)
  - [Heap sort](#heap-sort)


## Selection sort
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
## Insertion sort
簡單來說，插入排序法就是你玩撲克牌時用到的排序法。
1. 讀一個數字
    - 從「未排序好的數字」最左邊中讀取一個數
2. 插入合適位置
    - 把目前讀取的數字，與左邊的數值比大小；若左邊的數值比較大，則兩者互換位置，直到左邊的數字小於目前讀取的數字

complexity: O(n<sup>2</sup>)

ref: [初學者學演算法｜排序法入門：選擇排序與插入排序法](https://medium.com/appworks-school/%E5%88%9D%E5%AD%B8%E8%80%85%E5%AD%B8%E6%BC%94%E7%AE%97%E6%B3%95-%E6%8E%92%E5%BA%8F%E6%B3%95%E5%85%A5%E9%96%80-%E9%81%B8%E6%93%87%E6%8E%92%E5%BA%8F%E8%88%87%E6%8F%92%E5%85%A5%E6%8E%92%E5%BA%8F%E6%B3%95-23d4bc7085ff)

## Bubble sort
Bubble Sort 的方式是從陣列的最前面開始，一次比較陣列中兩兩相鄰的元素，然後根據大小將它們調換順序，大的移到後面
重複一次過程，此時應已確保最大的在最後面，故排除最右邊的元素

complexity: O(n<sup>2</sup>)
ref: [[演算法] 氣泡排序法（Bubble Sort）：利用兩兩元素交換位置達到排序](https://pjchender.blogspot.com/2017/09/bubble-sort.html)

## Merge sort
- 是一種 **divide and conquer** 演算法
- 基本想法是透過 recursive 不斷將 array 二分，直到 array size==1 則直接 return
- 而最小二分後拿到的左邊與右邊，再用一個 `merge()` 方法來合併
- 在 `merge()` 中會額外創建 merged array 供回傳 (memory usage 升高)
- 再額外用變數儲存 left and right array current positions 來走過元素並比較較小的，放進 merged array

**time complexity analysis**
- divide 原始 array 至最小單元，需要 $log(n)$ 次
- 接著從最小單元開始 merge，到底走訪了幾次 elements 呢？
- 思考這個問題可以從最後的結果倒推
- 最後的 n 是由 2 個 n/2 組合而來。而最後的 n 要能夠組起來表示 2 個 n/2 都看過一次，所以是 n 次
- 再往下推， n/2 也是來自 2 個 n/4 組合而來，而要組合 n/2 也需要走過 n/2 次，故在這一層總共也需要 n 次 (因為 2 個 n/2)
- 而我們的 tree 多高？ 就是 log(n) 那麼高
- 所以總共就需要 $nlog(n)$ 這麼多次的走訪

![https://www.khanacademy.org/computing/computer-science/algorithms/merge-sort/a/analysis-of-merge-sort](https://imgur.com/IDLRERH.png)

**is a stable sort**

也就是同樣的元素依舊維持在同樣的位置
需要在 merge 時注意，處理元素值相等時，以左邊的優先即可

complexity: $O(nlog(n))$

## Quick sort

- 也是一個 divide and conquer 演算法
- 利用陣列中最後一個元素（方便）當作 pivot（樞紐），並將陣列中小於 pivot 的元素往左邊放、大於 pivot 的元素往右邊放，然後 pivot 放中間
- 接著再遞迴地重複施以此算法於兩邊的陣列，直到陣列元素只剩下一個元素為止
- 關鍵在以 pivot 分邊的 `partition` 的動作該怎麼做
  - 直覺想到的是創建額外的空間來填（小於就填入左邊、大於則填入右邊）
  - 或是利用 two pointers 來挪動小於 pivot 及大於 pivot 的數值
    - 設計一個 recursive function (nums, start, end)
    - end 指的就是我們的 pivot
    - function 內比較 pivot-1 與 pivot 的大小；
      - 若 pivot-1 較小，則與 start 調換位置，start pointer++
      - 若 pivot-1 較大，則與 pivot 調換位置，pivot pointer--
      - 直到start pointer 與 pivot pointer 相撞為止
      - 最後再呼叫 recursive function 丟進 (nums, start, pivot-1) 與 (nums, pivot+1, end)

## Heap sort

堆積排序(Heap Sort)演算法是利用完全二元樹(Complete Binary Tree)，也就是堆積(Heap)的資料結構來完成排序的演算法。

兩大步驟：
1. 將要排序的陣列，用 `buildHeap(array)` 製作成 **Max Heap** or **Min Heap** ，取決於你要**遞增排序**還是**遞減排序**
2. 接著在 swap 頭尾元素值，完了之後用樣的 `buildHeap(array)` 調整陣列剩餘的部分，重複此步驟即可

**Complete Binary Tree** - 完全二元樹是一種二元樹，它只允許最後一層(最底下那層)的節點數量可以不必填滿
```
     5
   /   \
  1     3
 / \   /
6   2 4
```

**Min Heap** - 最小堆積就是上層節點數值必定會小於下層節點數值的完全二元樹。例如以下結構即為最小堆積
```
     1
   /   \
  2     5
 / \   /
4   6 3
```

**Max Heap** - 最大堆積就是上層節點數值必定會大於下層節點數值的完全二元樹。例如以下結構即為最大堆積
```
     6
   /   \
  5     3
 / \   /
4   2 1
```




ref: [堆積排序(Heap Sort)演算法，利用完全二元樹來排序的演算法](https://magiclen.org/heap-sort/)
