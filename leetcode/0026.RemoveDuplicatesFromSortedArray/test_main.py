from typing import List
import copy

import unittest
try:
    from .main import Solution  # need for linter
except:
    from main import Solution  # need for this file is main


solution = Solution()


class Args():
    def __init__(self, title: str, nums: List[int], expect: List[int]):
        self.title = title
        self.nums = nums
        self.expect = expect


class TestBasicFunction(unittest.TestCase):
    cases = [
        Args("ex1", [1, 1, 2], [1, 2]),
        Args("ex2", [0, 0, 1, 1, 1, 2, 2, 3, 3, 4], [0, 1, 2, 3, 4]),
    ]

    def test_inplace(self):
        for case in self.cases:
            nums = copy.deepcopy(case.nums)

            length = solution.removeDuplicates(nums)
            self.assertNotEqual(
                length, 0, "should has at least one unique element")

            for i in range(length):
                self.assertEqual(
                    nums[i], case.expect[i], "{}-th element".format(i))

    def test_pop(self):
        for case in self.cases:
            nums = copy.deepcopy(case.nums)

            length = solution.removeDuplicatesByPop(nums)
            self.assertNotEqual(
                length, 0, "should has at least one unique element")

            for i in range(length):
                self.assertEqual(
                    nums[i], case.expect[i], "{}-th element".format(i))


if __name__ == '__main__':
    unittest.main()
