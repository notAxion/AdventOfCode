def part1(filepath: str):
    with open(filepath, encoding="utf-8") as file:
        password = 0
        dial = 50
        for line in file.readlines():
            # print(f"line: {line.strip()}")
            line = line.strip()
            char = line[0]
            # print(char)
            num = int(line[1:])

            match char:
                case "L":
                    dial -= num
                case "R":
                    dial += num
            if dial % 100 == 0:
                password += 1

        print(password)


def part2(filepath: str):
    with open(filepath, encoding="utf-8") as file:
        password = 0
        dial = 50
        for line in file.readlines():
            # print(f"line: {line.strip()}")
            line = line.strip()
            char = line[0]
            # print(char)
            num = int(line[1:])

            olddial = dial
            match char:
                case "L":
                    dial -= num

                    actual_dial = (-1 * olddial) % 100
                    x_dial = actual_dial + num
                    if x_dial >= 100:
                        hunds = x_dial // 100
                        password += hunds

                case "R":
                    dial += num

                    actual_dial = olddial % 100
                    x_dial = actual_dial + num
                    if x_dial >= 100:
                        hunds = x_dial // 100
                        password += hunds

        print(password)


part1("input")
part2("input")

# i hope this kinda explains the mindset of how i solved part 2
# ❯ bpython
# bpython version 0.25 on top of Python 3.13.5 /usr/bin/python3
# >>> 95 - 105
# -10
# >>> 95 - 196
# -101
# >>> 95 - 390
# -295
# >>> 95 + 10
# 105
# >>> 95 + 101
# 196
# >>> 95 + 390
# 485
# >>> -95 + 96
# 1
# >>> -5 + 10
# 5
# >>> -5 + 101
# 96
# >>> -5 + 390
# 385
# >>> -95 % 100
# 5
# >>> -5 % 100
# 95
# >>> -305 % 100
# 95
# >>> 105 // 100
# 1
# >>> 196 // 100
# 1
# >>> 485 // 100
# 4
# >>> 400 // 100
# 4
# >>> 5 - 10
# -5
# >>> 5 - 101
# -96
# >>> 5 - 390
# -385
# >>> -95 - 10
# -105
# >>> -95 - 101
# -196
# >>> -95 - 390
# -485
# >>> -95 % 100
# 5
# >>> (-1 * -95) % 100
# 95
# >>> (-1 * 5) % 100
# 95
