class Song:
    def __init__(self, name):
        self.name = name
        self.next = None

    def next_song(self, song):
        self.next = song

    def is_repeating_playlist(self):
        """
        :returns: (bool) True if the playlist is repeating, False if not.
        """

        nextSong = nextNextSong = self
        while nextSong and nextNextSong:
            # one step pointer
            nextSong = nextSong.next
            if nextSong is None:
                return False

            # two steps pointer
            nextNextSong = nextNextSong.next
            if nextNextSong is None:
                return False
            nextNextSong = nextNextSong.next
            if nextNextSong is None:
                return False

            # check if meet again
            if id(nextSong) == id(nextNextSong):
                return True
        return False


def example1():
    first = Song("Hello")
    second = Song("Eye of the tiger")

    first.next_song(second)
    second.next_song(first)

    print(first.is_repeating_playlist())


def example2():
    first = Song("Hello")
    second = Song("Eye of the tiger")
    third = Song("Foo")
    four = Song("Bar")

    first.next_song(second)
    second.next_song(third)
    third.next_song(four)

    print(first.is_repeating_playlist())


def example3():
    a = Song("a")
    b = Song("b")
    c = Song("c")
    d = Song("d")
    e = Song("e")

    a.next_song(b)
    b.next_song(c)
    c.next_song(d)
    d.next_song(e)
    e.next_song(b)

    print("================ should all True ================")
    print(a.is_repeating_playlist())
    print(b.is_repeating_playlist())
    print(c.is_repeating_playlist())
    print(d.is_repeating_playlist())
    print(e.is_repeating_playlist())
    print("================ should all True ================")


if __name__ == "__main__":
    example1()
    example2()
    example3()
