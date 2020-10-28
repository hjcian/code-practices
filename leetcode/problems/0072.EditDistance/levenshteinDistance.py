import sys


def displayTable(table):
    for row in table:
        print(row)
    print()


def assert_equals(got, want):
    if got != want:
        print(f"got {got}, want {want}")
        sys.exit(1)


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


# levenshtein distance
def ld_recursive(a, b, i, j):
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
        return ld_recursive(a, b, i+1, j+1)

    # 實際上要找到的子問題是：
    res = 1 + min(
        ld_recursive(a, b, i+1, j),
        ld_recursive(a, b, i, j+1),
        ld_recursive(a, b, i+1, j+1),
    )

    return res


def levenshtein_recursive(a, b):
    return ld_recursive(a, b, 0, 0)


def ld_table(a, b, i, j, table, indent=""):
    if i < 0 or j < 0:
        # 表示某一層的小問題，有一邊被切成空字串了，所以回傳非空的那一邊的長度
        # 就是兩邊的 index+1，得到還有剩的那邊的字串長度
        print(indent, a, b, i, j, "... bottom case")
        return max(i+1, j+1)

    if table[i][j] != -1:
        # 接著處理是否這一層是一個小問題
        # 檢查是否 hit cache metrix，就知道了
        print(indent, a, b, i, j, "... cache hit")
        return table[i][j]

    if a[i] == b[j]:
        # 此位置的兩字元相等，退縮成拿到 substring 比較的結果，作為這一層的解答
        res = ld_table(a, b, i-1, j-1, table, indent=indent+"  ")
    else:
        # 此位置的兩字元不相等，edit distance 先自己 + 1
        # 接著再比較三種情況的小問題，誰比較小，則加上原本的 1
        res = 1 + min(
            ld_table(a, b, i-1, j, table, indent=indent+"  "),
            ld_table(a, b, i, j-1, table, indent=indent+"  "),
            ld_table(a, b, i-1, j-1, table, indent=indent+"  "),
        )

    table[i][j] = res  # 存下這一次的結果，給下一個 for loop 加速用
    print(indent, a[:i+1], b[:j+1], i, j, "... min. edit dist.:", res)

    return res


def levenshtein_table_cache(a, b):
    table = [[-1]*(len(b)) for _ in range(len(a))]

    # displayTable(table)
    for i, _ in enumerate(a):
        for j, _ in enumerate(b):
            ld_table(a, b, i, j, table, indent="")
            print("-"*50)
        print("="*100)
    # displayTable(table)

    return table[-1][-1]


def levenshtein_table_search_on_the_fly(word1, word2):
    # assum given: ("happy", "ros")
    word1Len = len(word1)
    word2Len = len(word2)
    table = [[0] * (word2Len + 1) for _ in range(word1Len + 1)]

    # 注意(2)
    # 初始值如何設定，可以觀察，每一個字在比的時候，最大編輯距離就是全換掉
    # 例如 r 與 happy 逐字比較的時候，最多最多就是 position number：
    #     r o s
    # 1 h 1
    # 2 a 2
    # 3 p 3
    # 4 p 4
    # 5 y 5
    # 故輔助用的 table 初始為全零之後，在最外層以 index 值填入，得到新的 table 如下：
    #     r o s
    #   0 1 2 3
    # h 1 0 0 0
    # a 2 0 0 0
    # p 3 0 0 0
    # p 4 0 0 0
    # y 5 0 0 0
    # 此時觀察 routine 中的 min() 就不會被 table[m-1][] or table[][n-1] 給誤導了
    # displayTable(table)
    for i in range(word1Len + 1):
        table[i][0] = i
    for j in range(word2Len + 1):
        table[0][j] = j

    # displayTable(table)
    # [0, 1, 2, 3]
    # [1, 0, 0, 0]
    # [2, 0, 0, 0]
    # [3, 0, 0, 0]
    # [4, 0, 0, 0]
    # [5, 0, 0, 0]
    for m in range(1, word1Len + 1):
        for n in range(1, word2Len + 1):
            if word1[m - 1] == word2[n - 1]:
                table[m][n] = table[m - 1][n - 1]
            else:
                # 注意(1)
                # m-1, n-1 在 table 中是沒意義的值，若 table 使用 0 作為初值的話
                # 計算 min() 時會被誤導
                # 故需要在前面就先給適當的初始值輔助運算
                # 被誤導的現象會影響 m=1 以及 n=1 這兩個貼近輔助線的運算時
                table[m][n] = 1 + min(
                    table[m - 1][n],
                    table[m][n - 1],
                    table[m - 1][n - 1],
                    )

    # displayTable(table)
    # [0, 1, 2, 3]
    # [1, 1, 2, 3]
    # [2, 2, 2, 3]
    # [3, 3, 3, 3]
    # [4, 4, 4, 4]
    # [5, 5, 5, 5]

    return table[word1Len][word2Len]


def levenshtein_mem_eff(word1, word2):
    word1Len = len(word1)
    word2Len = len(word2)
    if word1Len < word2Len:
        # table 的 spaces 為 2 * rows，所以 word2 越短空間越省
        return levenshtein_mem_eff(word2, word1)

    table = [
        [0] * (word2Len + 1)
        for _ in range(2)
        ]

    for i in range(2):
        table[i][0] = i
    for j in range(word2Len + 1):
        table[0][j] = j

    # displayTable(table)

    for m in range(1, word1Len + 1):
        table[m % 2][0] = m  # for next loop
        for n in range(1, word2Len + 1):
            if word1[m - 1] == word2[n - 1]:
                table[m % 2][n] = table[(m - 1) % 2][n - 1]
            else:
                table[m % 2][n] = 1 + min(
                    table[(m - 1) % 2][n],
                    table[m % 2][n - 1],
                    table[(m - 1) % 2][n - 1],
                    )
        # print("m:", m)
        # displayTable(table)

    return table[word1Len % 2][-1]


if __name__ == "__main__":

    # levenshtein = levenshtein_recursive
    # levenshtein = levenshtein_table
    # levenshtein = levenshtein_table_search_on_the_fly
    levenshtein = levenshtein_mem_eff

    assert_equals(levenshtein("sitten", "sitting"), 2)
    assert_equals(levenshtein("kitten", "sitting"), 3)
    assert_equals(levenshtein("book", "back"), 2)
    assert_equals(levenshtein("book", "book"), 0)
    assert_equals(levenshtein("qlzcfayxiz", "vezkvgejzb"), 9)
    assert_equals(levenshtein("nayvyedosf", "sjxen"), 9)
    assert_equals(levenshtein("sjxen", "sjxen"), 0)
    assert_equals(levenshtein("peter", "peter"), 0)
    assert_equals(levenshtein("horse", "ros"), 3)
    assert_equals(levenshtein("happy", "ros"), 5)
