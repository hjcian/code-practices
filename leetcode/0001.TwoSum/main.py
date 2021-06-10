from typing import List


class Solution:
    # Runtime: 60 ms, faster than 39.29% of Python3 online submissions for Two Sum.
    # Memory Usage: 15.4 MB, less than 8.90% of Python3 online submissions for Two Sum.
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        table = {num: i for i, num in enumerate(nums)}

        for i in range(len(nums)):
            if (target-nums[i]) in table and table[target-nums[i]] != i:
                return [i, table[target-nums[i]]]

    def bestAnswer(self, nums: List[int], target: int) -> List[int]:
        # Runtime: 60 ms, faster than 39.29% of Python3 online submissions for Two Sum.
        # Memory Usage: 15.4 MB, less than 18.13% of Python3 online submissions for Two Sum.
        table = {}
        for i in range(len(nums)):
            diff = target - nums[i]
            if diff in table:
                return i, table[diff]
            table[nums[i]] = i
