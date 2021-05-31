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

## merge sort

## quick sort

## heap sort