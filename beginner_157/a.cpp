#include <iostream>
using namespace std;

int main() {
    int n;
    cin >> n;

    int ans;
    ans = n / 2;
    if (n % 2 == 1) {
        ans += 1;
    }
    cout << ans << endl;
}
