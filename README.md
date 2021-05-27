個人程式碼練習集中存放區

- [Problems & practices](#problems--practices)
  - [Leetcode](#leetcode)
  - [classic data structures & algorithms practices](#classic-data-structures--algorithms-practices)
  - [Codewars](#codewars)
  - [CodeSignal](#codesignal)
- [思路統整](#思路統整)
- [參考資料](#參考資料)
  - [練習順序](#練習順序)
- [Miscellaneous](#miscellaneous)
  - [指令備忘錄](#指令備忘錄)
    - [Go](#go)
    - [folder management](#folder-management)
  - [有趣的議題](#有趣的議題)
  - [Custom badge](#custom-badge)
  - [Partner](#partner)





# Problems & practices
## Leetcode
> ![](https://img.shields.io/badge/LeetCode-Easy-brightgreen)
- [1. Two Sum](leetcode/problems/0001.TwoSum/0001.TwoSum.go)
- [26. Remove Duplicates from Sorted Array](leetcode/problems/0026.RemoveDuplicatesFromSortedArray/0026.RemoveDuplicatesFromSortedArray.go)
- [27. Remove Element](leetcode/problems/0027.RemoveElement/0027.RemoveElement.go)
- [53. Maximum Subarray](leetcode/problems/0053.MaximumSubarray/0053.MaximumSubarray.go)
- [88. Merge Sorted Array](leetcode/problems/0088.MergeSortedArray/0088.MergeSortedArray.go)
- [219. Contains Duplicate II](leetcode/problems/0219.ContainsDuplicateII/0219.ContainsDuplicateII.go)
- [283. Move Zeroes](leetcode/problems/0283.MoveZeroes/0283.MoveZeroes.go)
- [287. Find the Duplicate Number](leetcode/problems/0287.FindtheDuplicateNumber/0287.FindtheDuplicateNumber.go)
- [532. K-diff Pairs in an Array](leetcode/problems/0532.KdiffPairsinanArray/0532.KdiffPairsinanArray.go)
- [561. Array Partition I](leetcode/problems/0561.ArrayPartitionI/0561.ArrayPartitionI.go)
- [566. Reshape the Matrix](leetcode/problems/0566.ReshapetheMatrix/0566.ReshapetheMatrix.go)
- [628. Maximum Product of Three Numbers](leetcode/problems/0628.MaximumProductofThreeNumbers/0628.MaximumProductofThreeNumbers.go)
- [704. BinarySearch](leetcode/problems/0704.BinarySearch/0704.BinarySearch.go)
- [746. Min Cost Climbing Stairs](leetcode/problems/0746.MinCostClimbingStairs/0746.MinCostClimbingStairs.go)👁‍🗨
- [766. Toeplitz Matrix](leetcode/problems/0746.MinCostClimbingStairs/0746.MinCostClimbingStairs.go)👁‍🗨
- [867. Transpose Matrix](leetcode/problems/0867.TransposeMatrix/0867.TransposeMatrix.go)
- [977. Squares of a Sorted Array](leetcode/problems/0977.SquaresofaSortedArray/0977.SquaresofaSortedArray.go)👁‍🗨

> ![](https://img.shields.io/badge/LeetCode-Medium-orange)
- [2. Add Two Numbers](./leetcode/problems/0002.AddTwoNumbers/0002.AddTwoNumbers.go)
- [3. Longest Substring Without Repeating Characters](./leetcode/problems/0003.LongestSubstringWithoutRepeatingCharacters/0003.LongestSubstringWithoutRepeatingCharacters.go)
- [11. Container With Most Water](leetcode/problems/0011.ContainerWithMostWater/0011.ContainerWithMostWater.go)👁‍🗨
- [15. 3Sum](leetcode/problems/0015.3Sum/0015.3Sum.go)👁‍🗨
- [16. 3Sum Closest](leetcode/problems/0016.3SumClosest/0016.3SumClosest.go)👁‍🗨 3-pointers skill
- [19. Remove Nth Node From End of List](./leetcode/problems/0019.RemoveNthNodeFromEndofList/0019.RemoveNthNodeFromEndofList.go)
- [39. Combination Sum](leetcode/problems/0039.CombinationSum/0039.CombinationSum.go) 😞👁‍🗨
- [48. Rotate Image](leetcode/problems/0048.RotateImage/0048.RotateImage.go)👁‍🗨
- 54. Spiral Matrix
- [56. Merge Intervals](leetcode/problems/0056.MergeIntervals/0056.MergeIntervals.go)
- [1143. Longest Common Subsequence](leetcode/problems/1143.LongestCommonSubsequence/1143.LongestCommonSubsequence.go)


> ![](https://img.shields.io/badge/LeetCode-Hard-red)
- [72. Edit Distance](leetcode/problems/0072.EditDistance/0072.EditDistance.go)
  - related: 1143. Longest Common Subsequence ([ref](leetcode/problems/1143.LongestCommonSubsequence/1143.LongestCommonSubsequence.go))

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

# 思路統整

發現題目直覺暴力解需要 *O(n<sup>2</sup>)*，則使用 `map` 的資料結構歷遍一次(*O(n)*)建立查詢效率 *O(1)* 的 lookup table，接著再歷遍一次做比較
- [1. Two Sum](leetcode/problems/0001.TwoSum/)



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
- [Levenshtein Distance](leetcode/problems/0072.EditDistance/levenshteinDistance.py)
  - [good explanation](https://medium.com/@ethannam/understanding-the-levenshtein-distance-equation-for-beginners-c4285a5604f0)
## Custom badge
- https://shields.io/
## Partner
- 🤘 https://github.com/hongtw/coding-life
