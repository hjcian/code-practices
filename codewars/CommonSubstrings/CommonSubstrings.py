def displayTable(table):
    for row in table:
        print(row)
    print()


# Explaination Material
#    123456
#    abcdef
# z
#     1 2 3 4 5 6
#     a b c d e f
# 1 z
# 2 b
# 3 c
# 4 d
# 5 f

# Answer plate
#     1 2 3 4 5 6
#     a b c d e f
# 1 z 0 0 0 0 0 0
# 2 b 0 1 0 0 0 0
# 3 c 0 0 2 0 0 0
# 4 d 0 0 0 3 0 0
# 5 f 0 0 0 0 0 1
def substring_test(word1, word2):
    # your code here
    if not word1 or not word2:
        return False
    word1 = word1.lower()
    word2 = word2.lower()

    word1Len = len(word1)
    word2Len = len(word2)
    table = [[0] * (word2Len + 1) for _ in range(word1Len + 1)]

    # displayTable(table)

    for m in range(1, word1Len + 1):
        for n in range(1, word2Len + 1):
            if word1[m-1] == word2[n-1]:
                table[m][n] = 1 + table[m-1][n-1]
                if table[m][n] == 2:
                    return True
            else:
                table[m][n] = 0
        # displayTable(table)

    return False

def assert_equals(no, got, want):
    if got != want:
        print(f"[{no}] expect {want}, got {got}")
        # sys.exit(1)


if __name__ == "__main__":
    print("Basic tests")
    assert_equals("0", substring_test("otmhethi", "the"), True)
    assert_equals("1", substring_test("Something", "Home"), True)
    assert_equals("2", substring_test("Something", "Fun"), False)
    assert_equals("3", substring_test("Something", ""), False)
    assert_equals("4", substring_test("", "Something"), False)
    assert_equals("5", substring_test("BANANA", "banana"), True)
    assert_equals("6", substring_test("test", "lllt"), False)
    assert_equals("7", substring_test("", ""), False)
    assert_equals("8", substring_test("1234567", "541265"), True)
    assert_equals("9", substring_test("supercalifragilisticexpialidocious", "SoundOfItIsAtrocious"), True)
    assert_equals("10", substring_test("LoremipsumdolorsitametconsecteturadipiscingelitAeneannonaliquetligulautplaceratorciSuspendissepotentiMorbivolutpatauctoripsumegetaliquamPhasellusidmagnaelitNullamerostellustemporquismolestieaornarevitaediamNullaaliquamrisusnonviverrasagittisInlaoreetultricespretiumVestibulumegetnullatinciduntsempersemacrutrumfelisPraesentpurusarcutempusnecvariusidultricesaduiPellentesqueultriciesjustolobortisrhoncusdignissimNuncviverraconsequatblanditUtbibendumatlacusactristiqueAliquamimperdietnuncsempertortorefficiturviverra", "thisisalongstringtest"), True)
    assert_equals("11", substring_test("Yw490ENGODnh6JuFucv", "wU8VVua"), False)
