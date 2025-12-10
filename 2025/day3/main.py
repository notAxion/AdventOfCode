def part1():
    total = 0

    file = open("input", encoding="utf-8")
    for line in file.readlines():
        line = line.strip()
        # print(line)
        first = 0
        first_index = 0
        for i in range(len(line) - 1):
            # print(f"{num},", end="")
            num = int(line[i])
            if num > first:
                first = num
                first_index = i

        second = 0
        for i in range(first_index + 1, len(line)):
            num = int(line[i])
            if num > second:
                second = num
        # print(f"{first}{second}")
        total += (first * 10) + second

    print(total)


def part2():
    total = 0

    file = open("input", encoding="utf-8")
    for line in file.readlines():
        line = line.strip()
        line_total = 0
        # print(line)

        last_index = -1
        for digit in range(12):
            digit_num = 0
            for i in range(last_index + 1, len(line) - (11 - digit)):
                # print(f"{num},", end="")
                num = int(line[i])
                if num > digit_num:
                    digit_num = num
                    last_index = i

            line_total = (line_total * 10) + digit_num

        total += line_total

    print(total)


part1()
part2()
