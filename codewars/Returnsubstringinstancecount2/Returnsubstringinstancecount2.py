import unittest

def search_substr(full_text, search_text, allow_overlap=True):
    if not full_text or not search_text:
        return 0

    start_pos = 0
    count = 0
    while start_pos < len(full_text):
        matched_pos = full_text.find(search_text, start_pos)

        if matched_pos == -1: break
        else:
            count += 1
            start_pos = matched_pos
            if allow_overlap:
                start_pos += 1
            else:
                start_pos += len(search_text)
    return count

class MimicTestObject(object): pass

class TestObject(unittest.TestCase):
    def test_cases(self):
        Test = MimicTestObject()
        Test.assert_equals = self.assertEqual

        print("Basic tests")
        print("Should handle matches the simple way")
        Test.assert_equals(search_substr('aa_bb_cc_dd_bb_e', 'bb'), 2)
        Test.assert_equals(search_substr('aaabbbcccc', 'bbb'), 1)
        Test.assert_equals(search_substr('aaacccbbbcccc', 'cc'), 5)
        Test.assert_equals(search_substr('aaa', 'aa'), 2)
        print("Should handle non-overlapping cases")
        Test.assert_equals(search_substr('aaa', 'aa',False), 1)
        Test.assert_equals(search_substr('aaabbbaaa', 'bb',False), 1)
        print("Should handle empty strings on both sides")
        Test.assert_equals(search_substr('a', ''), 0)
        Test.assert_equals(search_substr('', 'a'), 0)
        Test.assert_equals(search_substr('', ''), 0)
        Test.assert_equals(search_substr('', '',False), 0)

if __name__ == '__main__':
    unittest.main()
