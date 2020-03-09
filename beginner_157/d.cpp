#include <iostream>
#include <vector>
#define rep(i,n) for (int i = 0; i < (n); ++i)
using namespace std;

struct unionfind {
	vector<int> d;	// 親：連結のサイズ×(-1)、子：親のID（非負）
	unionfind(int n = 0): d(n, -1) {}
	int find(int x) {
		if(d[x] < 0) return x;
		return d[x] = find(d[x]);
	}
	bool unite(int x, int y) {
		x = find(x); y = find(y);
		if(x == y) return false;
		if(d[x] > d[y]) swap(x, y);
		d[x] += d[y];
		d[y] = x;
		return true;
	}
	bool same(int x, int y) { return find(x) == find(y); }
	int size(int x) { return -d[find(x)]; }
};

int main() {
	int n, m, k;
	cin >> n >> m >> k;
	unionfind uf(n);
	vector<vector<int>> friends(n), blocks(n);
	rep(i, m) {
		int a, b;
		cin >> a >> b;
		a--; b--;
		friends[a].push_back(b);
		friends[b].push_back(a);
		uf.unite(a, b);
	}
	rep(i, k) {
		int a, b;
		cin >> a >> b;
		a--; b--;
		blocks[a].push_back(b);
		blocks[b].push_back(a);
	}
	rep(i, n){
		int ans = uf.size(i) - 1 - friends[i].size();
		for(int u: blocks[i]) {
			if (uf.same(i, u)) ans--;
		}
		cout << ans << " ";
	}
	cout << endl;
	
	return 0;
}
