A, B = list(map(int, input().split()))

if (A - B) >= 2:
    print(A + A - 1)
elif (B - A) >= 2:
    print(B + B - 1)
else:
    print(A + B)
