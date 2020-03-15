#include <iostream>
#include <string>
#include <vector>
#include <set>
#define rep(i, n) for (int i=0; i<n; ++i)
using namespace std;

int main() {
    int n, q;
    string s;
    cin >> n >> s >> q;
    vector< set< int > > is(26);
    for (int i = 0; i < n; i++) {
        is[s[i] - 'a'].insert(i);
    }

    for (int qi = 0; qi < q; qi++) {
        int t;
        cin >> t;
        if (t == 1) {
            int i;
            char c;
            cin >> i >> c;
            --i;
            is[s[i] - 'a'].erase(i);
            s.at(i) = c;
            is[s[i] - 'a'].insert(i);
        } else {
            int l, r;
            cin >> l >> r;
            --l;
            int ans = 0;
            for (int j = 0; j < 26; j++) {
                auto it = is[j].lower_bound(l);
                if (it != is[j].end() && *it < r) {
                    ++ans;
                }
            }
            cout << ans << endl;
        }
    }

    return 0;
}
