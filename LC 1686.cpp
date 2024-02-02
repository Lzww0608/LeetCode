class Solution {
public:
    int stoneGameVI(vector<int>& a, vector<int>& b) {
        int n = a.size();
        vector<int> v(n, 0), pos(n);
        for (int i = 0; i < n; ++i) {
            v[i] = a[i] + b[i];
        }
        iota(pos.begin(), pos.end(), 0);
        sort(pos.begin(), pos.end(), [&] (int i, int j) -> bool {
            return v[i] > v[j];
        });
        int sum_a = 0, sum_b = 0;
        for (int i = 0; i < n; ++i) {
            if (i & 1) {
                sum_b += b[pos[i]];
            } else {
                sum_a += a[pos[i]];
            }
        }
        if (sum_a > sum_b) {
            return 1;
        } else if (sum_a < sum_b) {
            return -1;
        }
        return 0;
    }
};