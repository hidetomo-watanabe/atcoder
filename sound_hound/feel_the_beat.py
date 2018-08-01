import math

C, D = list(map(int, list(input().split())))

# 問題に準じる
MIN_BPM = 140
MAX_BPM = 170


# 好きなBPM区間はX.0、好きじゃないBPM区間はX.5
# 2 ** X.0倍することで、区間の端を再現できるように
def calc_pos(n):
    n_min = n // MIN_BPM
    if n <= MAX_BPM * (2 ** (math.log2(n_min) // 1)):
        n_pos = math.log2(n_min) // 1
    else:
        n_pos = math.log2(n_min) // 1 + 0.5
    return n_pos


c_pos = calc_pos(C)
d_pos = calc_pos(D)

result = 0
# posを0.5ずつ増加
for i in range(int(2 * c_pos), int(2 * d_pos) + 1):
    tmp_pos = i * 0.5
    # 嫌いなBPM区間はskip
    if tmp_pos % 1 == 0.5:
        continue

    if tmp_pos == d_pos:
        end = D
    else:
        end = MAX_BPM * (2 ** tmp_pos)
    if tmp_pos == c_pos:
        start = C
    else:
        start = MIN_BPM * (2 ** tmp_pos)
    result += int(end - start)

print(result)
