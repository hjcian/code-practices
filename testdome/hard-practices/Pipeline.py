def pipeline(*funcs):
    def helper(arg):
        res = arg
        for func in funcs:
            res = func(res)
        return res
    return helper


fun = pipeline(lambda x: x * 3, lambda x: x + 1, lambda x: x / 2)
print(fun(3))  # should print 5.0
