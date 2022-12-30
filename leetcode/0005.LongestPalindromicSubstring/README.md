# Question

# Thoughts

# Answers
## Two pointers
> https://www.youtube.com/watch?v=XYQecbcd6_c

這段程式碼是使用一個暴力枚舉的方法來解決 Longest Palindromic Substring 問題的。

首先，它會枚舉每一個字符作為回文字串的中心，然後嘗試向左右延伸，看看最多能延伸多少。每當枚舉到一個字符時，它會先向右延伸，直到遇到不同的字符為止，這樣就會形成一個由相同字符組成的子字符串，比如 "aaa"。然後它會向左延伸，同時檢查左右兩邊的字符是否相同。如果相同，就繼續向左延伸；如果不同，則停止延伸，並記錄下當前的回文字串長度，並與目前已經記錄的最長回文字串長度進行比較，如果更長就更新最長回文字串的起始位置和長度。

最後，它會返回由最長回文字串的起始位置和長度所組成的子字符串。

這個演算法的時間複雜度是 O(n^2)，因為它需要枚舉每一個字符，並且對於每一個字符都要枚舉左右兩側的字符，檢查它們是否相同。

```python
class Solution:
    def longestPalindrome(self, s: str) -> str:
        n = len(s)
        longest_palindrome_start, longest_palindrome_len = 0, 1

        for i in range(0, n):
            right = i
            while right < n and s[i] == s[right]:
                right += 1
            # s[i, right - 1] inclusive are equal characters e.g. "aaa"

            # while s[left] == s[right], s[left, right] inclusive is palindrome e.g. "baaab"
            left = i - 1
            while left >= 0 and right < n and s[left] == s[right]:
                left -= 1
                right += 1

            # s[left + 1, right - 1] inclusive is palindromic
            palindrome_len = right - left - 1
            if palindrome_len > longest_palindrome_len:
                longest_palindrome_len = palindrome_len
                longest_palindrome_start = left + 1

        return s[longest_palindrome_start: longest_palindrome_start + longest_palindrome_len]
```

## Dynamic Programming
> 想不通，看不懂
>

這段程式碼是使用動態規劃的方法來解決 Longest Palindromic Substring 問題的。

首先，它使用一個二維布林陣列 dp 來記錄字符串中的子字符串是否是回文字串。它的大小是 n x n，其中 n 是字符串的長度。

然後，它會枚舉每一個子字符串的結尾字符，並倒序枚舉每一個子字符串的起始字符。对於每一個子字符串，它會檢查起始字符和結尾字符是否相同。如果相同，則再檢查起始字符和結尾字符之間的子字符串是否是回文字串。如果是，則將當前子字符串的回文狀態設置為 True，並更新當前最長回文字串的起始位置和長度。

最後，它會返回由最長回文字串的起始位置和長度所組成的子字符串。

這個演算法的時間複雜度是 O(n^2)，因為它需要枚舉每一個子字符串的起始和結尾位置，並且對於每一個子字符串都要進行一次檢查。

```python
class Solution:
    def longestPalindrome(self, s: str) -> str:
        n = len(s)
        dp = [[False] * n for _ in range(n)]
        for i in range(n):
            dp[i][i] = True
        longest_palindrome_start, longest_palindrome_len = 0, 1

        for end in range(0, n):
            for start in range(end - 1, -1, -1):
                # print('start: %s, end: %s' % (start, end))
                if s[start] == s[end]:
                    if end - start == 1 or dp[start + 1][end - 1]:
                        dp[start][end] = True
                        palindrome_len = end - start + 1
                        if longest_palindrome_len < palindrome_len:
                            longest_palindrome_start = start
                            longest_palindrome_len = palindrome_len
        return s[longest_palindrome_start: longest_palindrome_start + longest_palindrome_len]
```
