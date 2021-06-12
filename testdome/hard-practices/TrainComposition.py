from collections import deque


class _Train:
    def __init__(self, id):
        self.id = id
        self.prev = None
        self.next = None


class _TrainComposition:

    def __init__(self):
        self.head = None
        self.tail = None

    def _init_train(self, wagonId):
        if self.is_empty():
            self.head = self.tail = _Train(wagonId)
            return True
        return False

    def is_empty(self):
        return self.head is None

    def attach_wagon_from_left(self, wagonId):
        if self._init_train(wagonId):
            return

        train = _Train(wagonId)
        train.next = self.head
        self.head.prev = train
        self.head = train

    def attach_wagon_from_right(self, wagonId):
        if self._init_train(wagonId):
            return

        train = _Train(wagonId)
        self.tail.next = train
        train.prev = self.tail
        self.tail = train

    def detach_wagon_from_left(self):
        if self.is_empty():
            return
        ret = self.head
        self.head = self.head.next
        if self.head:
            self.head.prev = None
        else:  # no train remains
            self.head = self.tail = None

        return ret.id

    def detach_wagon_from_right(self):
        if self.is_empty():
            return
        ret = self.tail
        self.tail = self.tail.prev
        if self.tail:
            self.tail.next = None
        else:  # no train remains
            self.head = self.tail = None
        return ret.id

    def show(self):
        head = self.head
        ret = ''
        while head:
            ret += str(head.id) + " <-> "
            head = head.next
        print(ret and ret[:-5])


class TrainComposition:

    def __init__(self):
        self.deque = deque()

    def attach_wagon_from_left(self, wagonId):
        self.deque.appendleft(wagonId)

    def attach_wagon_from_right(self, wagonId):
        self.deque.append(wagonId)

    def detach_wagon_from_left(self):
        return self.deque.popleft()

    def detach_wagon_from_right(self):
        return self.deque.pop()

    def show(self):
        print(self.deque)


if __name__ == "__main__":
    train = TrainComposition()
    train.attach_wagon_from_left(7)
    train.attach_wagon_from_left(13)
    train.show()

    train.attach_wagon_from_right(99)
    train.attach_wagon_from_right(77)
    train.show()

    print("detach from right:", train.detach_wagon_from_right())  # should print 77
    train.show()

    print("detach from left:", train.detach_wagon_from_left())  # should print 13
    train.show()

    '''
    self implemented
    $ python TrainComposition.py
    13 <-> 7
    13 <-> 7 <-> 99 <-> 77
    detach from right: 77
    13 <-> 7 <-> 99
    detach from left: 13
    7 <-> 99

    builtin package
    $ python TrainComposition.py
    deque([13, 7])
    deque([13, 7, 99, 77])
    detach from right: 77
    deque([13, 7, 99])
    detach from left: 13
    deque([7, 99])
    '''
