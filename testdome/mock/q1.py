import math


def sol(a, b, c, sign):
    return (-1*b + sign*math.sqrt(b**2 - 4*a*c)) / (2*a)


def find_roots(a, b, c):
    x1 = sol(a, b, c, 1)
    x2 = sol(a, b, c, -1)
    return (x1, x2)


print(find_roots(2, 10, 8))
