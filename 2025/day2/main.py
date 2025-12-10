def part1():
    def is_invalid(num: str) -> bool:
        lenn = len(num)
        if lenn % 2 != 0:
            return False
        for i in range(lenn // 2):
            half = lenn // 2
            if num[i] != num[i + half]:
                return False
        return True

    total = 0
    file = open("input", encoding="utf-8")
    for num_range in file.readline().split(sep=","):
        nums = num_range.split(sep="-")
        start = int(nums[0])
        end = int(nums[1])
        for num in range(start, end + 1):
            if is_invalid(str(num)):
                # print(num)
                total += num
    print(total)


part1()
