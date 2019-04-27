import numpy as np

N = int(input())

# init
# dp[day][action]
# day: 1-N, action: 0-2(a-c)
dp = np.zeros((N + 1, 3)).astype(int)

# loop
for i in range(1, N + 1):
    abc = list(map(int, input().split()))
    dp[i][0] = max(dp[i-1][1], dp[i-1][2]) + abc[0]
    dp[i][1] = max(dp[i-1][0], dp[i-1][2]) + abc[1]
    dp[i][2] = max(dp[i-1][0], dp[i-1][1]) + abc[2]

# output
print(max(dp[N]))
