def get_adj(matrix, i, j):
    if 0 <= i < len(matrix) and 0 <= j < len(matrix[0]):
        return matrix[i][j]
    return "."


# max_num_ats of '@' around the point
def check_all_adj(matrix, i, j, max_num_ats):
    # print(f"ran matrix, {i}, {j}")
    num_ats = 0
    adjs = [
        [-1, -1],
        [-1, 0],
        [-1, 1],
        [0, -1],
        [0, 1],
        [1, -1],
        [1, 0],
        [1, 1],
    ]

    for adj in adjs:
        place = get_adj(matrix, i + adj[0], j + adj[1])
        if place != ".":
            num_ats += 1

            if num_ats >= max_num_ats:
                return False
    return True


def part1():
    total_rolls = 0

    file = open("input", encoding="utf-8")
    matrix = []
    for line in file.readlines():
        line = line.strip()
        matrix.append(line)

    # print(matrix)
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            # print(matrix[i][j], end="")
            place = matrix[i][j]
            if place == "@":
                if check_all_adj(matrix, i, j, 4):
                    total_rolls += 1
        #             print("x", end="")
        #         else:
        #             print(matrix[i][j], end="")
        #     else:
        #         print(matrix[i][j], end="")
        # print()

    print(total_rolls)


part1()
