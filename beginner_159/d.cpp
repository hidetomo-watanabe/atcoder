#include <iostream>
#define rep(i, n) for (int i=0; i<n; ++i)
using namespace std;

long long _calc_comb(long long n) {
    return (n * (n-1)) / 2;
}

int main() {
    int n;
    cin >> n;

    int a, as[n+1];
    long long counts[n+1]={0};
    rep(i, n) {
        cin >> a;
        as[i+1] = a;
        counts[a] += 1;
    }
    long long base = 0;
    rep(i, n) {
        base += _calc_comb(counts[i+1]);
    }
    rep(i, n) {
        cout << base - (counts[as[i+1]] - 1) << endl;
    }

    return 0;
}
