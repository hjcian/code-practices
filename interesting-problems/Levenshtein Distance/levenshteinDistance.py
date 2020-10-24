import sys


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


def assert_equals(got, want):
    if got != want:
        print(f"got {got}, want {want}")
        sys.exit(1)


def showMatrix(matrix):
    for i in matrix:
        print(i)


def ld_matrix(a, b, i, j, matrix, indent=""):
    if i < 0 or j < 0:
        # 表示某一層的小問題，有一邊被切成空字串了，所以回傳非空的那一邊的長度
        # 就是兩邊的 index+1，得到還有剩的那邊的字串長度
        print(indent, a, b, i, j, "... bottom case")
        return max(i+1, j+1)

    if matrix[i][j] != -1:
        # 接著處理是否這一層是一個小問題
        # 檢查是否 hit cache metrix，就知道了
        print(indent, a, b, i, j, "... cache hit")
        return matrix[i][j]

    if a[i] == b[j]:
        # 此位置的兩字元相等，退縮成拿到 substring 比較的結果，作為這一層的解答
        res = ld_matrix(a, b, i-1, j-1, matrix, indent=indent+"  ")
    else:
        # 此位置的兩字元不相等，edit distance 先自己 + 1
        # 接著再比較三種情況的小問題，誰比較小，則加上原本的 1
        res = 1 + min(
            ld_matrix(a, b, i-1, j, matrix, indent=indent+"  "),
            ld_matrix(a, b, i, j-1, matrix, indent=indent+"  "),
            ld_matrix(a, b, i-1, j-1, matrix, indent=indent+"  "),
        )

    matrix[i][j] = res  # 存下這一次的結果，給下一個 for loop 加速用
    print(indent, a[:i+1], b[:j+1], i, j, "... min. edit dist.:", res)

    return res


def levenshtein_matrix(a, b):
    matrix = [[-1]*(len(b)) for _ in range(len(a))]

    # showMatrix(matrix)
    for i, _ in enumerate(a):
        for j, _ in enumerate(b):
            ld_matrix(a, b, i, j, matrix, indent="")
            print("-"*50)
        print("="*100)
    # showMatrix(matrix)

    return matrix[-1][-1]


if __name__ == "__main__":
    # assert_equals(levenshtein("kitten", "sitting"), 3)
    # assert_equals(levenshtein("book", "back"), 2)
    # assert_equals(levenshtein("book", "book"), 0)
    # assert_equals(levenshtein("qlzcfayxiz", "vezkvgejzb"), 9)
    # assert_equals(levenshtein("nayvyedosf", "sjxen"), 9)
    # assert_equals(levenshtein("sjxen", "sjxen"), 0)
    # assert_equals(levenshtein("peter", "peter"), 0)

    assert_equals(levenshtein_matrix("kitten", "sitting"), 3)
    assert_equals(levenshtein_matrix("book", "back"), 2)
    assert_equals(levenshtein_matrix("book", "book"), 0)
    assert_equals(levenshtein_matrix("qlzcfayxiz", "vezkvgejzb"), 9)
    assert_equals(levenshtein_matrix("nayvyedosf", "sjxen"), 9)
    assert_equals(levenshtein_matrix("sjxen", "sjxen"), 0)
    assert_equals(levenshtein_matrix("peter", "peter"), 0)
