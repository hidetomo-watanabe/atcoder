from collections import deque

N = int(input())

colors = [-1] * N
trees = [{} for _ in range(N)]
for _ in range(N-1):
    u, v, w = list(map(int, input().split()))
    trees[u-1][v-1] = (w % 2 == 0)
    trees[v-1][u-1] = (w % 2 == 0)

targets = deque([0])
colors[0] = 0
while len(targets) > 0:
    target = targets.popleft()
    next_targets = trees[target]
    for next_target, is_even in next_targets.items():
        if colors[next_target] != -1:
            continue
        if is_even:
            colors[next_target] = colors[target]
        else:
            colors[next_target] = 1 - colors[target]
        targets.append(next_target)

for c in colors:
    print(c)
