# from typing import List

import unittest
try:
    from .main import Solution  # need for linter
except:
    from main import Solution  # need for this file is main


solution = Solution()


class Args():
    def __init__(self, title: str, s: str, expect: bool):
        self.title = title
        self.s = s
        self.expect = expect


class TestBasicFunction(unittest.TestCase):
    cases = [
        Args("ex1", "()", True),
        Args("ex2", "()[]{}", True),
        Args("ex3", "(]", False),
        Args("ex4", "([)]", False),
        Args("ex5", "{[]}", True),
        Args("7/91 tests passed", "[", False),
        Args("9/91 tests passed", "(){}}{", False),
        Args("85/91 tests passed", "]", False),
    ]

    def test(self):

        for case in self.cases:
            ret = solution.isValid(case.s)
            self.assertEqual(ret, case.expect, case.title)

            ret = solution.moreConcise(case.s)
            self.assertEqual(ret, case.expect, case.title)


if __name__ == '__main__':
    unittest.main()
