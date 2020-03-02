#include <stdio.h>
#include <map>
using namespace std;

int main() {
    int n, m;
    scanf("%d", &n);
    scanf("%d", &m);

    int s, c;
    map<int, int> b;
    for (int i=0; i<m; i++) {
        scanf("%d", &s);
        scanf("%d", &c);
        if (b.count(s)) {
            if (b[s] != c) {
                printf("-1\n");
                return 0;
            }
        }
        b[s] = c;
    }

    if (n > 1) {
        if (b.count(1) and b[1] == 0) {
            printf("-1\n");
            return 0;
        }
        if (not b.count(1)) {
            b[1] = 1;
        }
    }

    int ans[n] = {};
    for(auto i = b.begin(); i != b.end(); ++i) {
        ans[i->first - 1] = i->second;
    }
    for (int i=0; i<n; i++) {
        printf("%d", ans[i]);
    }
    printf("\n");
}
