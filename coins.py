a = int(input())
b = int(input())
c = int(input())
x = int(input())

count = 0
for a_tmp in range((x // 500) + 1):
    if a_tmp > a:
        break
    x_wo_a = x - a_tmp * 500
    for b_tmp in range((x_wo_a // 100) + 1):
        if b_tmp > b:
            break
        x_wo_ab = x_wo_a - b_tmp * 100
        if c >= (x_wo_ab // 50):
            count += 1
print(count)
