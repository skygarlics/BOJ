#include <stdio.h>
#include <vector>
#include <map>
using namespace std;



enum Direction {
	RIGHT = 0,
	DOWN = 1,
	LEFT = 2,
	UP = 3
};
Direction rotate(Direction dir, int delta) {
    return static_cast<Direction>((dir + delta + 4) % 4);
}

class Tuner {
public:
	int x;
	int y;
	unsigned char visited;
	Tuner() : x(0), y(0), visited(0) {}
	Tuner(int x, int y) : x(x), y(y) {
		this->visited = 0;
	}
	void setVisited(Direction dir) {
		this->visited |= (1 << dir);
	}
	bool isVisited(Direction dir) {
		return (this->visited & (1 << dir)) != 0;
	}
};

const int MAX_TUNER = 50;

int solve() {
	int N, R;
	scanf("%d %d", &N, &R);

	vector<Tuner> tuners(R);
	map<int, Tuner*> xMap[51];
	map<int, Tuner*> yMap[51];
	for (int i = 0; i < R; i++) {
		int x, y;
		scanf("%d %d", &x, &y);
		tuners[i] = Tuner(x, y);
		Tuner* ptr = &tuners[i];
		xMap[x][y] = ptr;
		yMap[y][x] = ptr;
	}
	int lX, lY;
	Direction dir;
	scanf("%d %d", &lX, &lY);

	if (lX < 1) {
		// laser from left border
		dir = RIGHT;
	} else if (lX > N) {
		dir = LEFT;
	} else if (lY < 1) {
		dir = UP;
	} else if (lY > N) {
		dir = DOWN;
	}

	int cnt = MAX_TUNER;
	while (cnt-- > 0) {
		if (dir == RIGHT) {
			auto it = yMap[lY].upper_bound(lX);
			if (it == yMap[lY].end()) {
				// laser scape to right
				printf("%d %d\n", N+1, lY);
				return 0;
			}
			Tuner* next = it->second;
			if (next->isVisited(LEFT)) {
				// already visited edge = cycle
				printf("0 0\n");
				return 0;
			}
			next->setVisited(LEFT);

			lX = next->x;
			lY = next->y;
			dir = rotate(dir, 1);
		} else if (dir == LEFT) {
			auto it = yMap[lY].lower_bound(lX);
			if (it == yMap[lY].begin()) { 
				// laser scape to left
				printf("0 %d\n", lY);
				return 0;
			}
			it--;
			Tuner* next = it->second;
			if (next->isVisited(RIGHT)) {
				// already visited edge = cycle
				printf("0 0\n");
				return 0;
			}
			next->setVisited(RIGHT);

			lX = next->x;
			lY = next->y;
			dir = rotate(dir, 1);
		} else if (dir == UP) {
			auto it = xMap[lX].upper_bound(lY);
			if (it == xMap[lX].end()) {
				// laser scape to up
				printf("%d %d\n", lX, N+1);
				return 0;
			}
			Tuner* next = it->second;
			if (next->isVisited(DOWN)) {
				// already visited edge = cycle
				printf("0 0\n");
				return 0;
			}
			next->setVisited(DOWN);

			lX = next->x;
			lY = next->y;
			dir = rotate(dir, 1);
		} else if (dir == DOWN) {
			auto it = xMap[lX].lower_bound(lY);
			if (it == xMap[lX].begin()) {
				// laser scape to down
				printf("%d %d\n", lX, 0);
				return 0;
			}
			it--;
			Tuner* next = it->second;
			if (next->isVisited(UP)) {
				// already visited edge = cycle
				printf("0 0\n");
				return 0;
			}
			next->setVisited(UP);

			lX = next->x;
			lY = next->y;
			dir = rotate(dir, 1);
		} 
	}
	// if cnt == 0; may be infinite loop
	printf("0 0\n");

	return 0;
}

int main() {
	int T;
	scanf("%d", &T);
	while (T-- >0) {
		solve();
	}
	return 0;
}