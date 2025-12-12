import math
import re


def part1():
    file = open("input", encoding="utf-8")

    matrix = []

    for line in file.readlines():
        line = line.strip()
        nums = re.split(r"\s+", line)
        matrix.append(nums)

    # print(matrix)

    total = 0

    for j in range(len(matrix[0])):
        nums = [int(matrix[i][j]) for i in range(len(matrix) - 1)]
        operator = matrix[len(matrix) - 1][j]
        result = 0
        match operator:
            case "*":
                result = math.prod(nums)
            case "+":
                result = sum(nums)

        total += result

    print(total)


def part2():
    file = open("input", encoding="utf-8")

    num_arr = []
    matrix = []

    last_op = ""
    data = file.read().splitlines()

    """
    from:
        123 328  51 64
         45 64  387 23
          6 98  215 314
        *   +   *   +

    to:
        1  *
        24
        356

        369+
        248
        8

         32*
        581
        175

        623+
        431
          4
    """
    for j in range(len(data[0])):
        line = ""
        for i in range(len(data)):
            line += data[i][j]

        line = line.strip()

        # print(f'line: "{line}"')
        if line == "":
            # print(f"num_arr: {num_arr}")
            matrix.append(num_arr)
            matrix.append(last_op)
            num_arr = []
            continue

        things = re.findall(r"\s*(\d+)\s*([*+])*", line)
        # print(f"things: {things}")
        things_iter = iter(things[0])
        if num := next(things_iter, None):
            num_arr.append(int(num))
        if (oper := next(things_iter, None)) and oper != "":
            last_op = oper

    matrix.append(num_arr)
    matrix.append(last_op)
    num_arr = []

    # print(matrix)

    total = 0

    matrix_iter = iter(matrix)

    while (nums := next(matrix_iter, None)) and nums is not None:
        operator = next(matrix_iter)
        result = 0
        match operator:
            case "*":
                result = math.prod(nums)
            case "+":
                result = sum(nums)

        total += result

    print(total)


part1()
part2()
