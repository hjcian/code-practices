from typing import List
import copy

import unittest
try:
    from .main import Solution  # need for linter
except:
    from main import Solution  # need for this file is main


solution = Solution()


class Args():
    def __init__(self, title: str, nums: List[int], val: int, expect: List[int]):
        self.title = title
        self.nums = nums
        self.val = val
        self.expect = expect


class TestBasicFunction(unittest.TestCase):
    cases = [
        Args("ex1", [3, 2, 2, 3], 3, [2, 2]),
        Args("ex2", [0, 1, 2, 2, 3, 0, 4, 2], 2, [0, 1, 4, 0, 3]),
        Args("2 / 113 test cases passed.", [], 0, []),
        Args("5 / 113 test cases passed.", [2], 3, [2]),
        Args("oppsite: 5 / 113 test cases passed.", [3], 3, []),
        Args("6 / 113 test cases passed.", [3, 3], 3, []),
    ]

    def test_inplace(self):
        for case in self.cases:
            nums = copy.deepcopy(case.nums)

            length = solution.removeElement(nums, case.val)

            got = nums[:length]
            self.assertListEqual(sorted(got), sorted(case.expect))

    def test_conciseSolution(self):
        for case in self.cases:
            nums = copy.deepcopy(case.nums)

            length = solution.conciseSolution(nums, case.val)

            got = nums[:length]
            self.assertListEqual(sorted(got), sorted(case.expect))


if __name__ == '__main__':
    unittest.main()
