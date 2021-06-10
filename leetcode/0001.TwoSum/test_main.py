from typing import List

import unittest
try:
    from .main import Solution  # need for linter
except:
    from main import Solution  # need for this file is main


solution = Solution()


class Args():
    def __init__(self, nums: List[int], target: int, expect: List[int]):
        self.nums = nums
        self.target = target
        self.expect = expect


class TestBasicFunction(unittest.TestCase):
    cases = [
        Args([2, 7, 11, 15], 9, [0, 1]),
        Args([3, 2, 4], 6, [1, 2]),
        Args([3, 3], 6, [0, 1]),
    ]

    def test(self):

        for case in self.cases:
            ret = solution.twoSum(case.nums, case.target)
            self.assertListEqual(sorted(ret), sorted(case.expect))

            ret = solution.bestAnswer(case.nums, case.target)
            self.assertListEqual(sorted(ret), sorted(case.expect))


if __name__ == '__main__':
    unittest.main()
