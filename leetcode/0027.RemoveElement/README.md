# Question
https://leetcode.com/problems/remove-element/

# Thoughts

## Python

***自己的原始思路***

用兩個指針，`i` 表示非 `val` 的位置，`j` 為 `val` 的位置

初始值直接指到頭與尾

while loop 從頭開始，若遇到 `nums[i]` 為 `val`，則反向 loop `j` 試著找到 `val` 的位置，然後交換值 (swap)

並且 `i += 1` 進入下一個迭代

這之中需要處理的例外是反向 loop 時，如果 `i == j` 表示雙指針已交會，且 `nums[i]` 非 `val`，那麼應直接 break 出去回傳 `i`

***參考解答 (最平凡解) - `conciseSolution`***

用一個指針 `i` 指著非 `val` 的值的尾端就夠了。初始值為 `0`，表示在還沒看過之前都是 `val`

搭配 loop，遇到非 `val` 的值，則指派給 `nums[i]`，然後 `i+=1`

此時會知道 `nums[0:i]` 都存放著非 `val` 值