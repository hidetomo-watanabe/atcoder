N, Y = list(map(int, input().split()))


def _calc_target(n, y):
    # Nが少なすぎる場合
    if n < y // 10000:
        return [-1] * 3

    for num_10 in range((y // 10000), -1, -1):
        n_tmp10 = n
        y_tmp10 = y
        n_tmp10 -= num_10
        y_tmp10 -= 10000 * num_10
        if n_tmp10 < y_tmp10 // 5000:
            return [-1] * 3
        for num_5 in range((y_tmp10 // 5000), -1, -1):
            n_tmp5 = n_tmp10
            y_tmp5 = y_tmp10
            n_tmp5 -= num_5
            y_tmp5 -= 5000 * num_5
            # check
            if 1000 * n_tmp5 == y_tmp5:
                num_1 = n_tmp5
                return [num_10, num_5, num_1]
    # default
    return [-1] * 3


print(' '.join(map(str, _calc_target(N, Y))))
