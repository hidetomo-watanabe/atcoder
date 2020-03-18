#include <iostream>
#include <set>
#include <algorithm>
#include <cmath>
using namespace std;

int main() {
    float a, b;
    cin >> a >> b;

    set <int> froma, fromb, is;
    for (int i=ceil(a / 0.08); i < ceil((a+1) / 0.08); ++i) {
        froma.insert(i);
    }
    for (int i=ceil(b / 0.10); i < ceil((b+1) / 0.10); ++i) {
        fromb.insert(i);
    }
    set_intersection(froma.begin(), froma.end(), fromb.begin(), fromb.end(), inserter(is, is.end()));

    int ans;
    ans = *is.begin();
    if (ans == 0) ans = -1;
    cout << ans << endl;

    return 0;
}
