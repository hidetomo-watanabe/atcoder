N, K = list(map(int, input().split()))

trees = {}
for _ in range(N-1):
    a, b = list(map(int, input().split()))
    if a not in trees.keys():
        trees[a] = set()
    if b not in trees.keys():
        trees[b] = set()
    trees[a].add(b)
    trees[b].add(a)

ans = K
if N > 1:
    memo = {1}
    targets = {1}
    while targets:
        next_targets = set()
        for target in targets:
            if target == 1:
                base = K-1
            else:
                base = K-2
            i = 0
            for child in trees[target]:
                if child not in memo:
                    memo.add(child)
                    ans = (ans*(base-i)) % (10**9+7)
                    next_targets.add(child)
                    i += 1
        targets = next_targets
print(ans)
