A, B, C = list(map(int, input().split()))

if (A < C and C < B) or (B < C and C < A):
    print('Yes')
else:
    print('No')
