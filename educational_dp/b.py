import numpy as np

N, k = list(map(int, input().split()))
if k > N:
    k = N
hs = list(map(int, input().split()))
hs = np.array([0] + hs)

# init
dp = np.zeros(N + 1).astype(int)
dp[2] = abs(hs[2] - hs[1])

# loop
if N > 2:
    for i in range(3, N + 1):
        """
        for j in range(1, k + 1):
            if i <= j:
                continue
            if dp[i] == 0:
                dp[i] = dp[i-j] + abs(hs[i] - hs[i-j])
            else:
                dp[i] = min(dp[i-j] + abs(hs[i] - hs[i-j]), dp[i])
        """
        x = max(1, i - k)
        dp[i] = np.min(dp[x:i] + np.abs(hs[i] - hs[x:i]))

# output
print(dp[N])
