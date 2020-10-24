def lcs(a, b, i, j):
    if i == len(a) or j == len(b):
        return 0

    if a[i] == b[j]:
        res = 1 + lcs(a, b, i+1, j+1)
    else:
        res = max(
            lcs(a, b, i+1, j),
            lcs(a, b, i, j+1),
        )
    # print(a[i], b[j], res)
    return res


# levenshtein
def ld(a, b, i, j):
    if len(a) == i or len(b) == j:
        # 表示其中一個字串已經用罄
        # 回傳剩餘的字串的長度，兩者剩餘的字串 slice 取大
        # 到達此處，表示問題可能已經變成在問：
        # "abc" vs. "" 要編輯幾次，故答案是 3 次
        return max(
            len(a[i:]),
            len(b[j:])
        )

    if a[i] == b[j]:
        # 根據 Levenshtein Distance 的定義，如果目前考慮的位置字元相等，
        # 表示直接考慮 small problem 即可
        return ld(a, b, i+1, j+1)

    # 實際上要找到的子問題是：
    res = 1 + min(
        ld(a, b, i+1, j),
        ld(a, b, i, j+1),
        ld(a, b, i+1, j+1),
    )

    return res


def levenshtein(a, b):
    return ld(a, b, 0, 0)


if __name__ == "__main__":
    a = "testasdtest"
    b = "asdasd"
    res = lcs(a, b, 0, 0)
    # print(res)

    a = "sitting"
    b = "kitten"
    print(ld(a, b, 0, 0))
