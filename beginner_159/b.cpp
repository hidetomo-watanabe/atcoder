#include <iostream>
#define rep(i, n) for (int i=0; i<n; ++i)
using namespace std;

int main() {
    string s;
    cin >> s;
    int n = s.size();
    string pre = s.substr(0, (n-1)/2);
    string post = s.substr((n+1)/2, (n-1)/2);

    string ans = "Yes";
    if (pre == post) {
        rep(i, pre.size()/2) {
            if (pre[i] != pre[pre.size()-i-1]) {
                ans = "No";
                break;
            }
        }
    } else {
        ans = "No";
    }
    cout << ans << endl;

    return 0;
}
