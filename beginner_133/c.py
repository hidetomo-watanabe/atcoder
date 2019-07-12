import itertools

L, R = list(map(int, input().split()))

if (R - L) >= 2019:
    ans = 0
else:
    l_mod = L % 2019
    r_mod = R % 2019
    if l_mod > r_mod:
        ans = 0
    else:
        ans = 2018
        for i, j in itertools.combinations(list(range(L, R+1)), 2):
            ans = min(ans, (i * j) % 2019)
print(ans)
