line = input()
L = len(line)

if L == 2:
    print(int(line[1]) + int(line[0]))
elif L == 4:
    print(int(line[:2]) + int(line[2:]))
else:
    if line[1] == '0':
        print(int(line[:2]) + int(line[2:]))
    else:
        print(int(line[0]) + int(line[1:]))