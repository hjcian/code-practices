def unique_names(names1, names2):
    return list(set(names1).union(names2))


if __name__ == "__main__":
    names1 = ["Ava", "Emma", "Olivia"]
    names2 = ["Olivia", "Sophia", "Emma"]
    # should print Ava, Emma, Olivia, Sophia
    print(unique_names(names1, names2))
