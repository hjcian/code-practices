from typing import List


class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        # Runtime: 56 ms, faster than 6.72% of Python3 online submissions for Remove Element.
        # Memory Usage: 14.3 MB, less than 45.59% of Python3 online submissions for Remove Element.
        print()
        i = 0
        j = len(nums) - 1
        print("original:", nums, ", val:", val)
        while i <= j:
            print("i:", i, nums, ", j:", j)
            if nums[i] == val:
                while i != j and nums[j] == val:
                    j -= 1
                if i == j:
                    break

                print("     swap i & j:", i, j)
                print("     (before)", nums)
                nums[i], nums[j] = nums[j], nums[i]
                print("      (after)", nums)

            i += 1
        print("final i:", i, ", nums:", nums)
        return i

    def conciseSolution(self, nums: List[int], val: int) -> int:
        # Runtime: 40 ms, faster than 15.15% of Python3 online submissions for Remove Element.
        # Memory Usage: 14.3 MB, less than 45.59% of Python3 online submissions for Remove Element.
        i = 0
        for num in nums:
            if num != val:
                nums[i] = num
                i += 1
        return i
