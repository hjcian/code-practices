# Question
https://leetcode.com/problems/counting-bits/



# Thoughts

簡單的 solution 就是走一遍，每次都用 while loop + bitwise and 累加 `1` 的數量

進階的做法是得知道 `Brian Kernighan’s Algorithm` ([leetcode ref](https://leetcode.com/problems/counting-bits/discuss/79538/Simple-Python-Solution/742084))

且歸納出此 table：
```
i= 1 ( 1 &  0 =  0)     1 &     0 =     0  n_ones =  1
i= 2 ( 2 &  1 =  0)    10 &     1 =     0  n_ones =  1
i= 3 ( 3 &  2 =  2)    11 &    10 =    10  n_ones =  2
i= 4 ( 4 &  3 =  0)   100 &    11 =     0  n_ones =  1
i= 5 ( 5 &  4 =  4)   101 &   100 =   100  n_ones =  2
i= 6 ( 6 &  5 =  4)   110 &   101 =   100  n_ones =  2
i= 7 ( 7 &  6 =  6)   111 &   110 =   110  n_ones =  3
i= 8 ( 8 &  7 =  0)  1000 &   111 =     0  n_ones =  1
i= 9 ( 9 &  8 =  8)  1001 &  1000 =  1000  n_ones =  2
i=10 (10 &  9 =  8)  1010 &  1001 =  1000  n_ones =  2
i=11 (11 & 10 = 10)  1011 &  1010 =  1010  n_ones =  3
i=12 (12 & 11 =  8)  1100 &  1011 =  1000  n_ones =  2
i=13 (13 & 12 = 12)  1101 &  1100 =  1100  n_ones =  3
i=14 (14 & 13 = 12)  1110 &  1101 =  1100  n_ones =  3
i=15 (15 & 14 = 14)  1111 &  1110 =  1110  n_ones =  4
i=16 (16 & 15 =  0) 10000 &  1111 =     0  n_ones =  1
i=17 (17 & 16 = 16) 10001 & 10000 = 10000  n_ones =  2
i=18 (18 & 17 = 16) 10010 & 10001 = 10000  n_ones =  2
i=19 (19 & 18 = 18) 10011 & 10010 = 10010  n_ones =  3
i=20 (20 & 19 = 16) 10100 & 10011 = 10000  n_ones =  2
```

可以觀察到 `n_ones(i=15) = n_ones(15 & 14) + 1` 這樣的規律，就能用 DP 的方式在 single pass 找完
