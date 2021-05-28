package containerwithmostwater

// 解題思路１
// 用遞迴嘗試所有可能的組合，但會因為 test case 28 input 過多造成 Time Limit Exceeded
// test case 28: [76,155,15,188,180,154,84,34,187,142,22,5,27,183,111,128,50,58,2,112,179,2,100,111,115,76,134,120,118,103,31,146,58,198,134,38,104,170,25,92,112,199,49,140,135,160,20,185,171,23,98,150,177,198,61,92,26,147,164,144,51,196,42,109,194,177,100,99,99,125,143,12,76,192,152,11,152,124,197,123,147,95,73,124,45,86,168,24,34,133,120,85,81,163,146,75,92,198,126,191]
/*
func max(a, b, c int) int {
	ret := a
	if b > ret {
		ret = b
	}
	if c > ret {
		ret = c
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calcMaxArea(height []int, left, right int) int {
	if left == right {
		return 0
	}
	leftMax := calcMaxArea(height, left+1, right)
	current := min(height[left], height[right]) * (right - left)
	rightMax := calcMaxArea(height, left, right-1)
	return max(leftMax, current, rightMax)
}

func maxArea(height []int) int {
	return calcMaxArea(height, 0, len(height)-1)
}
*/

// 解題思路２
// 利用雙指針的特性，動態規劃
// 左右兩端比較誰比較矮，以矮的那邊為準計算面積
// 接著只需要移動矮邊的指針，期待在右邊較高的前提下，移動左邊看看是不是能拿到比較好的面積值
// 按照此邏輯就能夠只在 O(N) 的時間內，透過局部最佳解找到全局最佳解

func maxArea(height []int) int {
	left, right, maxArea := 0, len(height)-1, 0
	for left < right {
		w := right - left
		h := 0
		if height[left] < height[right] {
			h = height[left]
			left++ // 移動此邊，期待下一次的迭代能比這一次好
		} else {
			h = height[right]
			right-- // 移動此邊，期待下一次的迭代能比這一次好
		}
		area := w * h // compute current area
		if area > maxArea {
			maxArea = area // replace if get a better value
		}
	}
	return maxArea
}
