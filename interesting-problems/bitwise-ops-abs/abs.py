import unittest

# more about Python internal integer representation
# https://www.codementor.io/@arpitbhayani/how-python-implements-super-long-integers-12icwon5vk

WHATEVER_YOU_WANT = 94879487
def myAbs(num):
    mask = num >> WHATEVER_YOU_WANT
    return (num ^ mask) - mask

def myAbs2(num):
    mask = num >> WHATEVER_YOU_WANT
    return (num + mask) ^ mask

class Mytest(unittest.TestCase):
    def test_abs(self):
        self.assertEqual(myAbs(3), 3)
        self.assertEqual(myAbs(-3), 3)
        self.assertEqual(myAbs(-123456), 123456)

    def test_abs2(self):
        self.assertEqual(myAbs2(3), 3)
        self.assertEqual(myAbs2(-3), 3)
        self.assertEqual(myAbs2(-123456), 123456)

if __name__ == "__main__":
    unittest.main()