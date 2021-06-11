# Question
https://leetcode.com/problems/remove-duplicates-from-sorted-array/


# Thoughts

## Python

關鍵提示：list 資料類型 (python) 是 reference

***自己原始思路（pop element from list）- `removeDuplicatesByPop()`***

在 loop 中改變陣列長度時，每一次 iteration 取 `len` 都能動態取得正確的長度

此時可直接用一個 while loop 搭配一個指針從 `i = 1` 開始走
- 不從 0 是因為若只有一個元素，他一定是為一個 unique 元素

而 while loop 的條件需要考慮陣列只有一個元素時的 `i` 與 `len(nums)` 的關係，當只有一個元素的時候我們就不走了，所以條件是 `while i < len(nums)`

接著在迭代中檢查 i 與 i+1 是否為相同元素
- 若相同，則移除 i 元素(⚠️)，此時保留指針還是 i 值走下一個迭代即可
  - (⚠️) 原本用陣列串接 `nums = nums[0:i] + nums[i+1:]` 會導致 reallocate 記憶體
  - 直接使用 `nums.pop(i)`，底層的 linked list 直接 O(1) 拔除元素，才不會導致外部與內部的 `nums` 變成指向不同的記憶體區塊(head)
- 若不相同，則 `i+=1` 往後走，走完後也把所有相同的元素的 pop 掉了

***參考解答的思路 - `removeDuplicates()`***

由於題目要求回傳一個長度值(`L`)，外部評分時則檢查 `nums` 直到 `L` 為止是否與答案相符

且題目有說請直接 `modifying the input array in-place with O(1) extra memory`

故最簡單的做法是直接邊走邊修改 list 中的元素