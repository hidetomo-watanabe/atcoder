import numpy as np

N = int(input())
A_LIST = np.array(list(map(int, input().split())))

x_list = np.zeros(N)
start_index = 0
x_list[start_index] = (A_LIST.sum() - 2 * A_LIST[1::2].sum()) / 2
tmp_index = start_index
for _ in range(N-1):
    next_index = (tmp_index+1) % N
    x_list[next_index] = A_LIST[tmp_index] - x_list[tmp_index]
    tmp_index = next_index
ans = ' '.join((x_list * 2).astype(int).astype(str))
print(ans)
