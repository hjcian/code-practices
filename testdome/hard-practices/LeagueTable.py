from collections import Counter
from collections import OrderedDict
from functools import cmp_to_key


def compare(left, right):
    # compare score
    if left[1] > right[1]:
        return -1
    elif left[1] < right[1]:
        return 1

    # equal score, compare games
    if left[2] < right[2]:
        return -1
    elif left[2] > right[2]:
        return 1

    # equal games, compare name order
    if left[3] < right[3]:
        return -1
    elif left[3] > right[3]:
        return 1

    # all equal
    return 0


class LeagueTable:
    def __init__(self, players):
        self.standings = OrderedDict(
            [(player, Counter()) for player in players])

    def record_result(self, player, score):
        self.standings[player]['games_played'] += 1
        self.standings[player]['score'] += score

    def player_rank(self, rank):
        grades = [(name, self.standings[name]["score"], self.standings[name]
                   ["games_played"], i) for i, name in enumerate(self.standings)]
        return sorted(grades, key=cmp_to_key(compare))[rank-1][0]


if __name__ == "__main__":
    table = LeagueTable(['Mike', 'Chris', 'Arnold'])
    table.record_result('Mike', 2)
    table.record_result('Mike', 3)
    table.record_result('Arnold', 5)
    table.record_result('Chris', 5)
    print(table.player_rank(1))
