def search_substr(full_text, search_text, allow_overlap=True):
    if not full_text or not search_text:
        return 0

    # these init. can reduce excution time during while loop
    limit = len(full_text) 
    step = allow_overlap and 1 or len(search_text)
    
    start_pos = 0
    count = 0
    while start_pos < limit:
        matched_pos = full_text.find(search_text, start_pos)

        if matched_pos == -1: 
            return count # early return to reduce instruction
        
        count += 1
        start_pos = matched_pos + step

import unittest
from collections import namedtuple

class TestObject(unittest.TestCase):
    def test_cases(self):
        Test = namedtuple("Test", "assert_equals")(self.assertEqual)
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

import re
def re_search_substr(full_text, search_text, allow_overlap=True):
    if not (full_text and search_text): return 0
    return len(re.findall(f"({allow_overlap and '?=' or ''}{search_text})", full_text))
    

def benchmark():
    import timeit

    template = '''
{0}('aa_bb_cc_dd_bb_e', 'bb')
{0}('aaacccbbbcccc', 'cc')
{0}('aaabbbaaa', 'bb',False)
'''

    print("Start C style (string manipulation)")
    res = timeit.timeit(template.format("search_substr"), setup="from __main__ import search_substr")
    print(f"{res:.5f} s")
    print("Start Pythonic style (leverage re module)")
    res = timeit.timeit(template.format("re_search_substr"), setup="from __main__ import re_search_substr")
    print(f"{res:.5f} s")

if __name__ == '__main__':
    benchmark()
    unittest.main()