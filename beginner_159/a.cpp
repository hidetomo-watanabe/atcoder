#include <iostream>
using namespace std;

int main() {
    int n, m;
    cin >> n >> m;
    int ans;
    ans = n * (n-1) / 2 + m * (m-1) / 2;
    cout << ans << endl;

    return 0;
}
