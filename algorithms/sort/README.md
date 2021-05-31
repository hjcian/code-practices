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
## insertion sort
簡單來說，插入排序法就是你玩撲克牌時用到的排序法。
1. 讀一個數字
    - 從「未排序好的數字」最左邊中讀取一個數
2. 插入合適位置
    - 把這個讀取的數與左邊的數值比大小，若左邊的數值比較大則兩者互換位置，直到左邊的數字大於這個讀取的數字

complexity: O(n<sup>2</sup>)

## bubble sort

## merge sort

## quick sort

## heap sort