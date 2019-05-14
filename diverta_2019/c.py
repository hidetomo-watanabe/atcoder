N = int(input())

ans = 0
last_a_num = 0
first_b_num = 0
multi_ba_num = 0
for _ in range(N):
    s = input()
    ans += s.count('AB')
    if s[-1] == 'A':
        last_a_num += 1
    if s[0] == 'B':
        first_b_num += 1
    if s[-1] == 'A' and s[0] == 'B':
        multi_ba_num += 1

connect_num = min(last_a_num, first_b_num)
if last_a_num == first_b_num and last_a_num == multi_ba_num:
    if connect_num > 0:
        connect_num -= 1
ans += connect_num

print(ans)
