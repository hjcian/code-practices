
def find_nexts(matrix, come_from, current):
    print()
    print("comefrom:", come_from)
    print(" current:", current)
    positions = []
    if current[0] > 0 and matrix[current[0]-1][current[1]]:
        positions.append((current[0]-1, current[1]))

    if current[0] < len(matrix)-1 and matrix[current[0]+1][current[1]]:
        positions.append((current[0]+1, current[1]))

    if current[1] > 0 and matrix[current[0]][current[1]-1]:
        positions.append((current[0], current[1]-1))

    if current[1] < len(matrix[0])-1 and matrix[current[0]][current[1]+1]:
        positions.append((current[0], current[1]+1))

    positions = [pos for pos in positions
                 if come_from is None or pos != come_from]
    print("next:", positions)
    return positions


def find_roads(matrix, come_from, current, destination):
    # return next two possible roads and reachable
    positions = find_nexts(matrix, come_from, current)
    if len(positions) == 0:
        return False

    if len([1 for pos in positions if pos == destination]) > 0:
        return True

    return len([find_roads(matrix, current, pos, destination)
               for pos in positions]) > 0


def route_exists(from_row, from_column, to_row, to_column, map_matrix):
    if (from_row, from_column) == (to_row, to_column):
        return True

    return find_roads(map_matrix,
                      come_from=None,
                      current=(from_row, from_column),
                      destination=(to_row, to_column))


if __name__ == '__main__':
    map_matrix = [
        [True, False, False],
        [True, True, False],
        [False, True, True]
    ]

    print(route_exists(0, 0, 2, 2, map_matrix))
    print(route_exists(2, 2, 2, 2, map_matrix))
    print(route_exists(0, 0, 0, 2, map_matrix))


# def route_exists(from_row, from_column, to_row, to_column, map_matrix):
#     visited = [[False for i in range(len(map_matrix[0]))]
#                for j in range(len(map_matrix))]

#     def route_exists2(from_row, from_column, to_row, to_column, map_matrix):
#         if (from_row < 0 or from_row >= len(map_matrix) or
#             from_column < 0 or from_column >= len(map_matrix[0]) or
#                 visited[from_row][from_column] == True):
#             return False
#         if map_matrix[from_row][from_column]:
#             visited[from_row][from_column] = True
#             # map_matrix[from_row][from_column]=False #traversed
#             if (from_row == to_row and from_column == to_column):
#                 return True
#             return (route_exists2(from_row-1, from_column, to_row, to_column, map_matrix) or
#                     route_exists2(from_row, from_column-1, to_row, to_column, map_matrix) or
#                     route_exists2(from_row, from_column+1, to_row, to_column, map_matrix) or
#                     route_exists2(from_row+1, from_column, to_row, to_column, map_matrix))

#     if route_exists2(from_row, from_column, to_row, to_column, map_matrix):
#         return True
#     else:
#         return False
