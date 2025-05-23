from collections import defaultdict
import sys;input = sys.stdin.readline

class Player:
    _id:int
    _location:int
    @property
    def location(self):
        return self._location
    @location.setter
    def location(self, value:int):
        self._location = value

    _inventory:defaultdict

    def __init__(self, id_:int):
        self.id_ = id_
        self.location = 1
        self._inventory = defaultdict(int)

    def move(self, new_location:int):
        self.location = new_location
    
    def farm(self, item:int):
        self._inventory[item] += 1
        # cheating case #1
        return self.location == item

    def craft(self, item1:int, item2:int):
        item1_cnt = self._inventory[item1]
        item2_cnt = self._inventory[item2]

        self._inventory[item1] = max(0, item1_cnt - 1)
        self._inventory[item2] = max(0, item2_cnt - 1)

        # cheating case #2
        return item1_cnt != 0 and item2_cnt != 0

    def attack(self, target:'Player'):
        # cheating case #3
        return self._location == target.location
    

T, N = map(int, input().split())
players = [None] + [Player(i) for i in range(1, N + 1)] # 1-indexed
cheatings = []
banhammer = {}
for _ in range(T):
    log_idx, p_idx, command, *arguments = input().split()
    p_idx = int(p_idx)
    player = players[p_idx]
    if command == "M":
        player.move(int(arguments[0]))
    elif command == "F":
        item = int(arguments[0])
        if not player.farm(item):
            cheatings.append(log_idx)
    elif command == "C":
        item1, item2 = map(int, arguments)
        if not player.craft(item1, item2):
            cheatings.append(log_idx)
    elif command == "A":
        target = players[int(arguments[0])]
        if not player.attack(target):
            cheatings.append(log_idx)
            banhammer[p_idx] = None

print(len(cheatings))
if cheatings:
    print(*cheatings)
print(len(banhammer))
if banhammer:
    print(*sorted(banhammer.keys()))


