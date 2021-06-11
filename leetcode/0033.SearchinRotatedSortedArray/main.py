from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        # Runtime: 32 ms, faster than 97.76% of Python3 online submissions for Search in Rotated Sorted Array.
        # Memory Usage: 14.8 MB, less than 20.40% of Python3 online submissions for Search in Rotated Sorted Array.
        if len(nums) == 0:
            return -1

        left = 0
        right = len(nums)
        while left < right:
            mid = (left + right) // 2

            # because we use floor division (//), so the `mid`
            # will close to `left` if remain two elements only.
            # check `mid` is same as check `left` if only one element remaining.
            if nums[mid] == target:
                return mid
            elif right - left == 1:
                return -1

            if nums[left] < nums[mid]:
                if nums[left] <= target < nums[mid]:
                    right = mid
                else:
                    left = mid+1
            else:
                if nums[mid] < target <= nums[right-1]:
                    left = mid + 1
                else:
                    right = mid
        return -1
