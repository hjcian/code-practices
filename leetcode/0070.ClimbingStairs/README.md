# Question
https://leetcode.com/problems/climbing-stairs/

# Thoughts

動態規劃，藉由將大問題拆成小問題來解

與 divide and concur 不同的是，動態規劃指的是小問題有重複的解可以運用，而不用重複計算

n=1，只有 1 種走法
- 1

n=2，只有 2 種走法
- 1+1
- 2

n=3，有趣了，要反過來想：
- 假設成功踩完 n=3 的最後一步走 1 步，表示前面有兩階要走，而根據前面已知的結果，兩階總共有 2 種走法
- 又假設最後一步走的是 2 步，表示前面有一階要走，而一階總共只有 1 種走法
- 故 n=3 總共有 2 + 1 = 3 種走法

# related problem
[509. Fibonacci Number](../0509.FibonacciNumber/README.md)