S = input()

words = ['dream', 'dreamer', 'erase', 'eraser']
words_reversed = list(map(lambda x: ''.join(reversed(x)), words))

S_reversed = ''.join(reversed(S))
tmp = S_reversed
while True:
    for word in words_reversed:
        if tmp[:len(word)] == word:
            tmp = tmp[len(word):]
            break
    else:
        print('NO')
        break

    if len(tmp) == 0:
        print('YES')
        break
