#include <iostream>
#include <string>
#include <algorithm>
#define rep(i, n) for (int i=0; i<n; ++i)
using namespace std;

int main() {
    string s;
    int q;
    cin >> s >> q;

    bool is_r = false;
    string pre_s, post_s;
    rep(i, q) {
        int t;
        cin >> t;
        if (t == 1) {
            is_r = !is_r;
        } else if (t == 2) {
            int f;
            char c;
            cin >> f >> c;
            if (f == 1) {
                if (is_r) {
                    post_s += c;
                } else {
                    pre_s = c + pre_s;
                }
            } else if (f == 2) {
                if (is_r) {
                    pre_s = c + pre_s;
                } else {
                    post_s += c;
                }
            }
        }
    }

    if (is_r) {
        rep(i, post_s.size()) cout << post_s[post_s.size()-i-1];
        rep(i, s.size()) cout << s[s.size()-i-1];
        rep(i, pre_s.size()) cout << pre_s[pre_s.size()-i-1];
    } else {
        cout << pre_s;
        cout << s;
        cout << post_s;
    }
    cout << endl;

    return 0;
}
