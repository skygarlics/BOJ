#include <iostream>
using namespace std;

int printTime(int h, int m, int s) {
	cout << h << " ";
	cout << m << " ";
	cout << s;
	cout << endl;
	return 0;
}
int printTime(int t) {
	int h = t / 3600;
	int m = (t % 3600) / 60;
	int s = t % 60;
	printTime(h, m, s);
	return 0;
}

int main() {
	int aH1, aM1, aS1, aH2, aM2, aS2;
	int bH1, bM1, bS1, bH2, bM2, bS2;
	int cH1, cM1, cS1, cH2, cM2, cS2;

	cin >> aH1 >> aM1 >> aS1 >> aH2 >> aM2 >> aS2;
	cin >> bH1 >> bM1 >> bS1 >> bH2 >> bM2 >> bS2;
	cin >> cH1 >> cM1 >> cS1 >> cH2 >> cM2 >> cS2;

	int a1 = aH1 * 3600 + aM1 * 60 + aS1;
	int b1 = bH1 * 3600 + bM1 * 60 + bS1;
	int c1 = cH1 * 3600 + cM1 * 60 + cS1;

	int a2 = aH2 * 3600 + aM2 * 60 + aS2;
	int b2 = bH2 * 3600 + bM2 * 60 + bS2;
	int c2 = cH2 * 3600 + cM2 * 60 + cS2;

	int a = a2 - a1;
	int b = b2 - b1;
	int c = c2 - c1;

	printTime(a);
	printTime(b);
	printTime(c);
	return 0;
}