import numpy as np

N = int(input())
nums = list(map(int, list(input().split())))

cross_list = []
for index in range(N):
    cross_list.append((-1) ** index)
result = np.sum(np.sort(np.array(nums))[::-1] * (np.array(cross_list)))
print(result)
