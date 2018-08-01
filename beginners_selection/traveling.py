N = int(input())
# init point is 0 0 0
data = [[0, 0, 0]]
for _ in range(N):
    data.append(list(map(int, input().split())))


def can_travel(datum1, datum2):
    diff_time = datum2[0] - datum1[0]
    diff_point = abs((datum2[1] + datum2[2]) - (datum1[1] + datum1[2]))
    # too much diff_point
    if diff_point > diff_time:
        return False
    # different parity
    if diff_time % 2 != diff_point % 2:
        return False
    return True


for i in range(len(data)):
    if i == len(data) - 1:
        print('Yes')
        break
    if not can_travel(data[i], data[i+1]):
        print('No')
        break
