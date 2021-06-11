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
        Args("ex1", [-2, 1, -3, 4, -1, 2, 1, -5, 4], 6),
        Args("ex2", [1], 1),
        Args("ex3", [5, 4, -1, 7, 8], 23),
    ]

    def test_inplace(self):
        for case in self.cases:
            nums = copy.deepcopy(case.nums)

            result = solution.maxSubArray(nums)
            self.assertEqual(result, case.expect, case.title)


if __name__ == '__main__':
    unittest.main()
