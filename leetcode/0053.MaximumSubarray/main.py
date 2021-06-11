from typing import List


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        # Runtime: 76 ms, faster than 19.64% of Python3 online submissions for Maximum Subarray.
        # Memory Usage: 15.2 MB, less than 13.72% of Python3 online submissions for Maximum Subarray.
        localMax = globalMax = nums[0]
        for i in range(1, len(nums)):
            current = nums[i]
            localMax = max(current, localMax+current)
            globalMax = max(globalMax, localMax)
        return globalMax
