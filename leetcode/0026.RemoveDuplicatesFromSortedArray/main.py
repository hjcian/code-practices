from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        # Runtime: 80 ms, faster than 81.85% of Python3 online submissions for Remove Duplicates from Sorted Array.
        # Memory Usage: 16.1 MB, less than 10.87% of Python3 online submissions for Remove Duplicates from Sorted Array.
        bound = len(nums) - 1
        idx = 1
        for i in range(bound):
            if nums[i] != nums[i+1]:
                nums[idx] = nums[i+1]
                idx += 1
        return idx

    def removeDuplicatesByPop(self, nums: List[int]) -> int:
        # Runtime: 96 ms, faster than 40.49% of Python3 online submissions for Remove Duplicates from Sorted Array.
        # Memory Usage: 15.9 MB, less than 53.14% of Python3 online submissions for Remove Duplicates from Sorted Array.
        i = 1
        while i < len(nums):
            if nums[i-1] == nums[i]:
                nums.pop(i)
            else:
                i += 1
        return i
