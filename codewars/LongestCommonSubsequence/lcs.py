import unittest
import pprint

def lcs(x, y):
    hashtable = {}
    for idx, c in enumerate(x):
        c_indices = hashtable.get(c, [])
        c_indices.append(idx)
        hashtable[c] = c_indices

    lcs = ""
    leftPos = -1
    for c in y:
        if c in hashtable: # hash table lookup complexity is O(1)
            available_indices = [ idx for idx in hashtable[c] if idx > leftPos ]
            if available_indices: # means this character can be used
                lcs += c
                leftPos = available_indices[0]
                hashtable[c] = available_indices[1:] # save for speed up

    return lcs

class TestLCS(unittest.TestCase):
    def test_lcs(self):
        # self.assertEqual(lcs("a", "b"), "")
        # self.assertEqual(lcs("abcdef", "abc"), "abc")
        # self.assertEqual(lcs("132535365", "123456789"), "12356")
        self.assertEqual(lcs("132535365", "1234567893"), "12356")
        # self.assertEqual(lcs("finaltest", "zzzfinallyzzz"), "final")

if __name__ == "__main__":
    unittest.main()

