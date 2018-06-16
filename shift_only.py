n = int(input())
nums = list(map(int, input().split()))


def calc_count(nums):
    budget = nums
    count = 0
    while True:
        for i, num in enumerate(nums):
            if num % 2 == 0:
                budget[i] = num / 2
            else:
                return count
        count += 1


print(calc_count(nums))
