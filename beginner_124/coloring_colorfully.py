S = input()

answer1 = '10' * (len(S) // 2) + '1' * (len(S) % 2)
answer2 = '01' * (len(S) // 2) + '0' * (len(S) % 2)
count1 = bin(int(answer1, 2) ^ int(S, 2))[2:].count('1')
count2 = bin(int(answer2, 2) ^ int(S, 2))[2:].count('1')
print(min(count1, count2))
