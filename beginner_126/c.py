import math

N, K = list(map(int, input().split()))

ans = 0.0
for i in range(1, N+1):
    count = max(0, math.ceil(math.log(K/i, 2)))
    ans += 1/(N * (2 ** count))
print(ans)
