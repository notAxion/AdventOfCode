def part1():
    with open("input", encoding="utf-8") as file:
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


part1()
