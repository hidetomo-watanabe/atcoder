import numpy as np

N = int(input())

ans = 0
for i in range(1, int(np.sqrt(N)) + 1):
    if N % i != 0:
        continue
    m = N//i - 1
    if m <= i:
        continue
    if N % m == 0:
        continue
    ans += m
print(ans)
