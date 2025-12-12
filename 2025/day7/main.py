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


def calculate_universes(
    manifold: list[list[str]], i: int, j: int, beam_map: dict[str, int]
) -> int:
    if not (0 <= i < len(manifold)) or not (0 <= j < len(manifold[0])):
        return 1
    space = manifold[i][j]
    # if space == "|":
    #     return
    universes = 0
    # go down
    if space == ".":
        key = f"{i + 1}x{j}"
        if (saved_unis := beam_map.get(key)) is not None:
            return saved_unis
        # manifold[i] = indexreplace(manifold[i], j, "|")
        # manifold[i][j] = "|"
        side_unis = calculate_universes(manifold, i + 1, j, beam_map)
        beam_map[key] = side_unis
        universes += side_unis
    elif space == "^":
        # beam_map[f"{i}x{j}"] = True
        # go left
        side_unis = calculate_universes(manifold, i, j - 1, beam_map)
        universes += side_unis
        # go right
        side_unis = calculate_universes(manifold, i, j + 1, beam_map)
        universes += side_unis

    return universes


def part2():
    file = open("input", encoding="utf-8")

    manifold = []

    for line in file.readlines():
        line = line.strip()
        manifold.append(line)

    # for line in manifold:
    #     print(line)

    beam_map = {}
    total_universes = 0

    for j in range(len(manifold[0])):
        if manifold[0][j] == "S":
            # start with going straight down
            total_universes = calculate_universes(
                manifold,
                1,
                j,
                beam_map,
            )
            # print(0, j)
            break

    # for line in manifold:
    #     print(line)

    # print(hit_map)
    print(total_universes)


part1()
part2()
