個人程式碼練習集中存放區

- [Problems & practices](#problems--practices)
  - [Leetcode](#leetcode)
  - [classic data structures & algorithms practices](#classic-data-structures--algorithms-practices)
  - [Codewars](#codewars)
  - [CodeSignal](#codesignal)
- [類似觀念集中整理](#類似觀念集中整理)
  - [Binary search variants](#binary-search-variants)
- [基礎知識/DSA](#基礎知識dsa)
  - [algorithms](#algorithms)
  - [data structures](#data-structures)
  - [未分類](#未分類)
    - [Bitwise operation](#bitwise-operation)
    - [Miscellaneous](#miscellaneous)
- [參考資料](#參考資料)
  - [練習順序](#練習順序)
- [Miscellaneous](#miscellaneous-1)
  - [指令備忘錄](#指令備忘錄)
    - [Go](#go)
    - [folder management](#folder-management)
  - [有趣的議題](#有趣的議題)
  - [Custom badge](#custom-badge)
  - [Partner](#partner)


# Problems & practices
## Leetcode
> ![Easy](https://img.shields.io/badge/LeetCode-Easy-brightgreen)
>
- [1. Two Sum](leetcode/0001.TwoSum/README.md) - go, python
- [20. Valid Parentheses](leetcode/0020.ValidParentheses/README.md) - python
- [26. Remove Duplicates from Sorted Array](leetcode/0026.RemoveDuplicatesFromSortedArray/README.md) - go, python
- [27. Remove Element](leetcode/0027.RemoveElement/README.md) - go, python
- [35. Search Insert position](leetcode/0035.SearchInsertPosition/README.md) - go
- [53. Maximum Subarray](leetcode/0053.MaximumSubarray/README.md) - go, python
- (DP) [70. Climbing Stairs](leetcode/0070.ClimbingStairs/README.md) - go
- [88. Merge Sorted Array](leetcode/0088.MergeSortedArray/0088.MergeSortedArray.go)
- (tree) [94. Binary Tree Inorder Traversal](leetcode/0094.BinaryTreeInorderTraversal/README.md) - go (recursive, iterative)
- (tree) [144. Binary Tree Preorder Traversal](leetcode/0144.BinaryTreePreorderTraversal/README.md) - go (recursive, iterative)
- ⚠️ (tree) [145. Binary Tree Postorder Traversal](leetcode/0145.BinaryTreePostorderTraversal/README.md) - go... iterative 還想不通 😫
- [219. Contains Duplicate II](leetcode/0219.ContainsDuplicateII/0219.ContainsDuplicateII.go)
- [283. Move Zeroes](leetcode/0283.MoveZeroes/0283.MoveZeroes.go)
- [287. Find the Duplicate Number](leetcode/0287.FindtheDuplicateNumber/0287.FindtheDuplicateNumber.go)
- [338. Counting Bits](leetcode/0338.CountingBits/)
- (DP) [509. Fibonacci Number](leetcode/0509.FibonacciNumber/README.md) - go
- [532. K-diff Pairs in an Array](leetcode/0532.KdiffPairsinanArray/0532.KdiffPairsinanArray.go)
- [561. Array Partition I](leetcode/0561.ArrayPartitionI/0561.ArrayPartitionI.go)
- [566. Reshape the Matrix](leetcode/0566.ReshapetheMatrix/0566.ReshapetheMatrix.go)
- [628. Maximum Product of Three Numbers](leetcode/0628.MaximumProductofThreeNumbers/0628.MaximumProductofThreeNumbers.go)
- [704. BinarySearch](leetcode/0704.BinarySearch/0704.BinarySearch.go)
- [746. Min Cost Climbing Stairs](leetcode/0746.MinCostClimbingStairs/0746.MinCostClimbingStairs.go)👁‍🗨
- [766. Toeplitz Matrix](leetcode/0746.MinCostClimbingStairs/0746.MinCostClimbingStairs.go)👁‍🗨
- [867. Transpose Matrix](leetcode/0867.TransposeMatrix/0867.TransposeMatrix.go)
- [977. Squares of a Sorted Array](leetcode/0977.SquaresofaSortedArray/0977.SquaresofaSortedArray.go)👁‍🗨

> ![Medium](https://img.shields.io/badge/LeetCode-Medium-orange)
>
- [2. Add Two Numbers](./leetcode/0002.AddTwoNumbers/0002.AddTwoNumbers.go)
- [3. Longest Substring Without Repeating Characters](./leetcode/0003.LongestSubstringWithoutRepeatingCharacters/0003.LongestSubstringWithoutRepeatingCharacters.go)
- [11. Container With Most Water](leetcode/0011.ContainerWithMostWater/0011.ContainerWithMostWater.go)👁‍🗨
- [15. 3Sum](leetcode/0015.3Sum/)👁‍🗨
- [16. 3Sum Closest](leetcode/0016.3SumClosest/0016.3SumClosest.go)👁‍🗨 3-pointers skill
- [19. Remove Nth Node From End of List](./leetcode/0019.RemoveNthNodeFromEndofList/0019.RemoveNthNodeFromEndofList.go)
- [33. Search in Rotated Sorted Array](leetcode/0033.SearchinRotatedSortedArray/README.md) - python
- [39. Combination Sum](leetcode/0039.CombinationSum/0039.CombinationSum.go) 😞👁‍🗨
- [48. Rotate Image](leetcode/0048.RotateImage/0048.RotateImage.go)👁‍🗨
- 54. Spiral Matrix
- [56. Merge Intervals](leetcode/0056.MergeIntervals/0056.MergeIntervals.go)
- [75. Sort Colors](leetcode/0075.SortColors/README.md) - go
- [1143. Longest Common Subsequence](leetcode/1143.LongestCommonSubsequence/1143.LongestCommonSubsequence.go)


> ![Hard](https://img.shields.io/badge/LeetCode-Hard-red)
>
- [72. Edit Distance](leetcode/0072.EditDistance/0072.EditDistance.go)
  - related: 1143. Longest Common Subsequence ([ref](leetcode/1143.LongestCommonSubsequence/1143.LongestCommonSubsequence.go))

## classic data structures & algorithms practices
- [Sortings](algorithms/sorting.go)
- [Bloom Filter](./data-structures/bloomfilter.go) - a naive proof of concept
  - [資料結構大便當：Bloom Filter](https://medium.com/@Kadai/%E8%B3%87%E6%96%99%E7%B5%90%E6%A7%8B%E5%A4%A7%E4%BE%BF%E7%95%B6-bloom-filter-58b0320a346d)
  - usages: [crypto/hash](https://gobyexample.com/sha1-hashes), [Go: Convert byte array to big.Int](https://stackoverflow.com/questions/24757814/golang-convert-byte-array-to-big-int/36944328), [Go: modulus using math big package](https://stackoverflow.com/questions/24098959/golang-modulus-using-math-big-package), [Go: `big.Int` SetBit()/Bit()](https://stackoverflow.com/a/53681508/8694937)

## Codewars
- [Return substring instance count 2](codewars/Returnsubstringinstancecount2/Returnsubstringinstancecount2.py)
- [Simple Pig Latin](codewars/SimplePigLatin/SimplePigLatin.js)
- [Pick peaks](./codewars/Pickpeaks/Pickpeaks.go)
- [Interlaced Spiral Cipher](./codewars/InterlacedSpiralCipher/InterlacedSpiralCipher.go)

## CodeSignal
- [convertIPv4](codesignal/convertIPv4/convertIPv4.go)
- [christmasTree](codesignal/christmasTree/christmasTree.go)
- happyNumber: [go](codesignal/happyNumber/happyNumber.go), [ts](codesignal/happyNumber/happyNumber.ts)
- videoPart: [go](codesignal/videoPart/videoPart.go), [ts](codesignal/videoPart/videoPart.ts)
- factorial_fun: [ts](codesignal/factorial_fun/factorial_fun.ts)
- isBeautifulString: [js](codesignal/isBeautifulString/isBeautifulString.js)
- [composeRanges](codesignal/composeRanges/composeRanges.js)

# 類似觀念集中整理

## Binary search variants
  - [SortedSearch - less than](https://www.testdome.com/questions/python/sorted-search/40608?visibility=3&skillId=9&orderBy=Difficulty) - [python](testdome/hard-practices/SortedSearch.py)
  - [33. Search in Rotated Sorted Array](leetcode/0033.SearchinRotatedSortedArray/README.md) - python


# 基礎知識/DSA
## algorithms
- [sorting](algorithms/sorting/README.md)
- [searching](algorithms/searching/README.md)

## data structures

**tree**

tree pre/in/post-order traversal 的定義： https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder/

![tree](https://media.geeksforgeeks.org/wp-content/cdn-uploads/2009/06/tree12.gif)

```
   1
 2   3
4 5
```

- 所謂的 pre/in/post-order traversal 其實是深度優先（Depth First Traversals）搜尋：
  - (a) **In**order (Left, **Root**, Right) : [4 2 5 1 3]
  - (b) **Pre**order (**Root**, Left, Right) : [1 2 4 5 3]
  - (c) **Post**order (Left, Right, **Root**) : [4 5 2 3 1]
- 而廣度優先（Breadth First or Level Order Traversal）搜尋則會得到： 1 2 3 4 5


## 未分類
### Bitwise operation
**Brian Kernighan’s Algorithm ([count set bits in an integer](https://www.geeksforgeeks.org/count-set-bits-in-an-integer/))**
- 用來計算一個整數的二進位表示法裡有多少的 `1`
- 很神奇地，做一個 while loop，n > 0，並且將 n 與 (n-1) 做幾次 bitwise &，就能知道有幾個 `1`
```
cnt = 0
while (n > 0):
  n = n & (n-1)
  cnt++
```

### Miscellaneous

發現題目直覺暴力解需要 *O(n<sup>2</sup>)*，則使用 `map` 的資料結構歷遍一次(*O(n)*)建立查詢效率 *O(1)* 的 lookup table，接著再歷遍一次做比較
- [1. Two Sum](leetcode/0001.TwoSum/)


# 參考資料
## 練習順序
- 👉 https://leetcode.com/problemset/all/?topicSlugs=array
- 🚫 https://books.halfrost.com/leetcode/ChapterTwo/Array/

# Miscellaneous
## 指令備忘錄
### Go
- test/benchmark (https://my.oschina.net/solate/blog/3034188)
  - `go test -bench=. -run=none -benchmem`
    - `-bench=.` 指的是當前路徑
    - `-run=none` run 原本是用來匹配想要執行的單元測試。不去設定會全跑，若想要全部都不跑就指定一個一定不存在的 pattern (none)
    - `-benchmem` 打開記憶體配置量量測

### folder management
- `python createFolder.py -c -t "convertIPv4"` create a codesignal problem
- `python createFolder.py -l -t "283. Move Zeroes"` create a leetcode problem
- `npm run ts codesignal/happyNumber/happyNumber.ts`

## 有趣的議題
- [bitwise-ops: add overflow detect](interesting-problems/bitwise-ops-addOk/add_overflow_detect.go)
- [Levenshtein Distance](leetcode/0072.EditDistance/levenshteinDistance.py)
  - [good explanation](https://medium.com/@ethannam/understanding-the-levenshtein-distance-equation-for-beginners-c4285a5604f0)

## Custom badge
- https://shields.io/

## Partner
- 🤘 https://github.com/hongtw/coding-life
