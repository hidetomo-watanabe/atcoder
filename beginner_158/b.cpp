#include <iostream>
using namespace std;

int main() {
    long long n, a, b;
    cin >> n >> a >> b;

    long long q, r, ans;
    q = n / (a + b);
    r = n % (a + b);
    ans = a * q + min(a, r);
    cout << ans << endl;

    return 0;
}
