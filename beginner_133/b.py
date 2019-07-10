import itertools
import numpy as np

N, D = list(map(int, input().split()))
X_LIST = []
for _ in range(N):
    X_LIST.append(list(map(int, input().split())))

ans = 0
for x1, x2 in itertools.combinations(X_LIST, 2):
    distance = np.linalg.norm(np.array(x1) - np.array(x2))
    if distance.is_integer():
        ans += 1
print(ans)
