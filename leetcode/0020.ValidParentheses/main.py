from typing import List


class Solution:
    openBrackets = {c: i for i, c in enumerate('([{')}
    closeBrackets = {c: i for i, c in enumerate(')]}')}

    def isValid(self, s: str) -> bool:
        # Runtime: 32 ms, faster than 61.86% of Python3 online submissions for Valid Parentheses.
        # Memory Usage: 14.3 MB, less than 33.49% of Python3 online submissions for Valid Parentheses.
        openQueue = []

        for i, c in enumerate(s):
            if len(openQueue) == 0 and c in self.closeBrackets:
                return False

            if c in self.openBrackets:
                openQueue.append(i)

            elif self.openBrackets[s[openQueue[-1]]] != self.closeBrackets[c]:
                # c is close bracket, check if same pair
                return False
            else:
                openQueue.pop()

        return len(openQueue) == 0

    pairs = {"(": ")", "{": "}", "[": "]"}

    def moreConcise(self, s: str) -> bool:
        # Runtime: 24 ms, faster than 95.25% of Python3 online submissions for Valid Parentheses.
        # Memory Usage: 14.2 MB, less than 63.08% of Python3 online submissions for Valid Parentheses.
        stack = []
        for c in s:
            if c in self.pairs.keys():
                # push open bracket
                stack.append(c)
            elif len(stack) == 0:
                # encounter close but stack is empty
                return False
            elif self.pairs[stack.pop()] != c:
                # stack has open bracket but not consistent with close one
                return False
        return stack == []
