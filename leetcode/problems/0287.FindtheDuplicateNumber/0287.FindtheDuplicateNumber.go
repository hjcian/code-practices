package findtheduplicatenumber

// ====== 解題思路 ======
// array 長度為 n+1，裡面只會含有 1~n 的數字。舉例，n=4
// array 長度就是 5，然後只有以下的可能（故意由小排到大）
// [1,1,2,3,4]
// [1,2,2,3,4]
// [1,2,3,3,4]
// [1,2,3,4,4]
// [1,1,1,2,3] ...
// 觀察，想一下，發現可以利用數值當作 index，跳著尋找、檢查元素值
// 假設 array 為 [3, 1, 3, 4, 2]
// step 1： （index 初始值為 0）
// index = 0，發現值為 3，那我知道下一步要去 index 3
// step 2：
// index = 3，發現值為 4，那我知道下一步要去 index 4
// step 3：
// index = 4，發現值為 2，那我知道下一步要去 index 2
// step 4：
// index = 2，發現值為 3，那我知道下一步要去 index 3
// step 5：
// index = 3，發現值為 4，那我知道下一步要去 index 4 ... 發現我會回到 step 3，進入無窮迴圈！
// 所以，看來我得在每次去下一步前，將目前的元素值 asssign 成非 1~n 的數值，此例我選擇 -1
// 就能在 step 5 時發現值為 -1，然後檢查目前的 index 值，就是重複的數值了

func findDuplicate(nums []int) int {
	idx := 0
	for {
		if nums[idx] == -1 {
			return idx
		}
		idx, nums[idx] = nums[idx], -1
	}
}
