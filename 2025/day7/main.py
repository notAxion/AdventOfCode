def indexreplace(s, index, char):
    return s[:index] + char + s[index + 1:]


def random_beam_go(manifold: list[list[str]], i: int, j: int, hit_map: dict[str, bool]):
    if not (0 <= i < len(manifold)) or not (0 <= j < len(manifold[0])):
        return
    space = manifold[i][j]
    if space == "|":
        return
    # go down
    if space == ".":
        manifold[i] = indexreplace(manifold[i], j, "|")
        # manifold[i][j] = "|"
        random_beam_go(manifold, i + 1, j, hit_map)
    elif space == "^":
        hit_map[f"{i}x{j}"] = True
        # go left
        random_beam_go(manifold, i, j - 1, hit_map)
        # go right
        random_beam_go(manifold, i, j + 1, hit_map)


def part1():
    file = open("input", encoding="utf-8")

    manifold = []

    for line in file.readlines():
        line = line.strip()
        manifold.append(line)

    # for line in manifold:
    #     print(line)

    hit_map = {}

    for j in range(len(manifold[0])):
        if manifold[0][j] == "S":
            # start with going straight down
            random_beam_go(
                manifold,
                1,
                j,
                hit_map,
            )
            # print(0, j)
            break

    # for line in manifold:
    #     print(line)

    # print(hit_map)
    print(len(hit_map))


part1()
