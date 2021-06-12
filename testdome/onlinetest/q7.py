from collections import OrderedDict


def first_unique_product(products):
    od = OrderedDict()
    for p in products:
        if p in od:
            od[p] += 1
        else:
            od[p] = 1

    for key in od:
        if od[key] == 1:
            return key

    return None


if __name__ == "__main__":
    # should print "Computer"
    print(first_unique_product(["Apple", "Computer", "Apple", "Bag"]))
