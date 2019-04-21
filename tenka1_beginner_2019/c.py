N = int(input())
S = input()

all_black_num = S.count('#')
all_white_num = N - all_black_num

if all_black_num == 0 or all_white_num == 0:
    ans = 0
else:
    left_black_num = 0
    left_white_num = 0
    right_white_num = all_white_num - left_white_num
    right_black_num = all_black_num - left_black_num
    ans = left_black_num + right_white_num
    for i, c in enumerate(S):
        if c == '#':
            left_black_num += 1
        elif c == '.':
            left_white_num += 1
        right_white_num = all_white_num - left_white_num
        right_black_num = all_black_num - left_black_num
        ans = min(ans, left_black_num + right_white_num)
        if right_black_num == N - i - 1:
            break
print(ans)
