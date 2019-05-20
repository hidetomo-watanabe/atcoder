S = input()
front = int(S[:2])
back = int(S[2:])


def is_month(num):
    if num >= 1 and num <= 12:
        return True


if is_month(front) and is_month(back):
    print('AMBIGUOUS')
elif is_month(front) and not is_month(back):
    print('MMYY')
elif not is_month(front) and is_month(back):
    print('YYMM')
else:
    print('NA')
