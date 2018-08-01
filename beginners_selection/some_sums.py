import itertools
N, A, B = list(map(int, list(input().split())))


def calc_targets(digit_num, digit_sum, N):
    output = []
    all_num_products = list(itertools.product(range(10), repeat=digit_num))
    for product in all_num_products:
        if product[0] == 0:
            continue
        if sum(product) == digit_sum:
            num = int(''.join(map(str, product)))
            if num <= N:
                output.append(num)
    return output


nums = []
for digit_num in range(1, len(str(N)) + 1):
    for digit_sum in range(A, B + 1):
        nums.extend(calc_targets(digit_num, digit_sum, N))
print(sum(nums))
