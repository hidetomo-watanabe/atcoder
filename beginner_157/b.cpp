#include <iostream>
using namespace std;

int main() {
    int a[3][3];
    int l;
    for (int i=0; i<3; i++) {
        for (int j=0; j<3; j++) {
            cin >> l;
            a[i][j] = l;
        }
    }

    int n;
    cin >> n;
    int b;
    for (int i=0; i<n; i++) {
        cin >> b;
        for (int i=0; i<3; i++) {
            for (int j=0; j<3; j++) {
                if (a[i][j] == b) {
                    a[i][j] = 0;
                }
            }
        }
    }

    string ans = "No";
    for (int i=0; i<3; i++) {
        if (a[i][0] + a[i][1] + a[i][2] == 0) {
            ans = "Yes";
            break;
        } else if (a[0][i] + a[1][i] + a[2][i] == 0) {
            ans = "Yes";
            break;
        } else if (a[0][0] + a[1][1] + a[2][2] == 0) {
            ans = "Yes";
            break;
        } else if (a[0][2] + a[1][1] + a[2][0] == 0) {
            ans = "Yes";
            break;
        }
    }
    cout << ans << endl;
}
