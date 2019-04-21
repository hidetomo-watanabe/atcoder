N = int(input())
S = input()
K = int(input())

target = S[K - 1]
ans = ''
for i, c in enumerate(S):
    if c == target:
        ans += c
    else:
        ans += '*'
print(ans)
