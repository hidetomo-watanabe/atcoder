N = int(input())
H_LIST = list(map(int, input().split()))

count = 0
for i in range(N):
    j = 0
    while (i - j) >= 0:
        if i == j:
            count += 1
            break
        if H_LIST[i] < H_LIST[j]:
            break
        j += 1
print(count)
