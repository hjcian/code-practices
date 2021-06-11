from typing import List
import copy

import unittest
try:
    from .main import Solution  # need for linter
except:
    from main import Solution  # need for this file is main


solution = Solution()


class Args():
    def __init__(self, title: str, nums: List[int], target: int, expect: List[int]):
        self.title = title
        self.nums = nums
        self.target = target
        self.expect = expect


class TestBasicFunction(unittest.TestCase):
    cases = [

        Args("ex1", [4, 5, 6, 7, 0, 1, 2], 0, 4),
        Args("ex2", [4, 5, 6, 7, 0, 1, 2], 3, -1),
        Args("ex3", [1], 0, -1),
        Args("ex3 - found", [0], 0, 0),
        Args("empty array", [], 0, -1),
        Args("two element - found", [1, 2], 1, 0),
        Args("two element - found", [1, 2], 2, 1),
        Args("two element - not found", [1, 2], 0, -1),
        Args("my imagination", [50, 60, 80, 90, 100, 125, 150, 0, 50], 0, 7),
        Args("my imagination", [50, 60, 80, 90, 100, 125, 150, 0, 50], 70, -1),
        Args("10 / 195 test cases passed.", [3, 1], 0, -1),
        Args("168 / 195 test cases passed.", [5, 1, 3], 3, 2),
    ]

    def test_inplace(self):
        for case in self.cases:
            nums = copy.deepcopy(case.nums)

            idx = solution.search(nums, case.target)
            self.assertEqual(idx, case.expect, case.title)


if __name__ == '__main__':
    unittest.main()
