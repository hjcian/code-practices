def count_numbers(sorted_list, less_than):
    left = 0
    right = len(sorted_list)
    while left < right:
        mid = (left + right) // 2  # floor division
        if sorted_list[mid] == less_than:
            return mid
        if right - left == 1:
            # last iteration, left == mid, but mid is not equal to less_than
            # so return left position
            if less_than > sorted_list[left]:
                return right
            return left
        if less_than < sorted_list[mid]:
            right = mid
        else:
            left = mid + 1
    return left


if __name__ == "__main__":
    sorted_list = [1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 20, 21, 22, 23]
    print(count_numbers(sorted_list, 4))  # should print 2
    print(count_numbers(sorted_list, 7))  # should print 3
    print(count_numbers(sorted_list, 8))  # should print 4
