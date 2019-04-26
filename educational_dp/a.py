import numpy as np

N = int(input())
hs = list(map(int, input().split()))
hs = [0] + hs

# init
dp = np.zeros(N + 1).astype(int)
dp[2] = abs(hs[2] - hs[1])
# loop
if N > 2:
    for i in range(3, N + 1):
        dp[i] = min(
            dp[i-2] + abs(hs[i] - hs[i-2]),
            dp[i-1] + abs(hs[i] - hs[i-1]))

# output
print(dp[N])
