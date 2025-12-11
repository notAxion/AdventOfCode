# sorts and merges and ranges
def process_ranges(arr):
    def first_lower(sorting_arr):
        return sorting_arr[0]

    arr = sorted(arr, key=first_lower)

    # eg:
    # from: [[3, 5], [10, 14], [12, 18], [16, 20], [21, 25], [27, 35], [28, 39], [41, 43], [42, 48]]
    # to: [[3, 5], [10, 25], [27, 39], [41, 48]]
    i = 0
    while i < len(arr) - 1:
        nums = arr[i]
        next_nums = arr[i + 1]
        while nums[0] <= next_nums[0] <= nums[1] + 1:
            if nums[1] < next_nums[1]:
                nums[1] = next_nums[1]
            del arr[i + 1]
            if i >= len(arr) - 1:
                break
            next_nums = arr[i + 1]

        i += 1

    return arr


def part1():
    file = open("input", encoding="utf-8")

    num_ranges = []
    nums = []
    is_num = False

    lines = file.readlines()
    # first get the ranges
    for line in lines:
        line = line.strip()
        if line == "":
            is_num = True
            continue
        if not is_num:
            both_nums = line.split(sep="-")
            num_ranges.append([int(both_nums[0]), int(both_nums[1])])
        else:
            nums.append(int(line))

    # print(num_ranges)
    # print(nums)

    num_ranges = process_ranges(num_ranges)
    # print("after sort and merge")
    # print(num_ranges)

    nums.sort()

    range_iter = iter(num_ranges)
    curr_range = next(range_iter, None)

    fresh_count = 0

    for num in nums:
        # check if its more than current B; next(num_range)
        while curr_range is not None and num > curr_range[1]:
            curr_range = next(range_iter, None)

        # if curr_range is None; then False
        if curr_range is None:
            break

        # check if its within the current; then True
        if curr_range[0] <= num <= curr_range[1]:
            fresh_count += 1
            # print(f"{num} is fresh")

    print(fresh_count)


def part2():
    file = open("input", encoding="utf-8")

    num_ranges = []

    lines = file.readlines()
    # first get the ranges
    for line in lines:
        line = line.strip()
        if line == "":
            break
        both_nums = line.split(sep="-")
        num_ranges.append([int(both_nums[0]), int(both_nums[1])])

    num_ranges = process_ranges(num_ranges)
    # print("after sort and merge")
    # print(num_ranges)

    total_ids = 0

    for nums in num_ranges:
        total_ids += nums[1] - nums[0] + 1

    print(total_ids)


part1()
part2()
