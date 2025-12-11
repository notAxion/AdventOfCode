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


part1()
