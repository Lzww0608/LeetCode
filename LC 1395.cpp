class Solution {
    const static int N = 1e5 + 5;
    int tr1[N], tr2[N];
public:
    void update (int tree[], int x, int v) {
        for (; x < N; x += -x & x) {
            tree[x] += v;
        }
    }

    int query (int tree[], int x) {
        int res = 0;
        for (; x; x -= -x & x) {
            res += tree[x];
        }
        return res;
    }

    int numTeams(vector<int>& rating) {
        int n = rating.size();
        int ans = 0;
        for (int i = 0; i < n; ++i) {
            update(tr2, rating[i], 1);
        }
        for (int i = 0; i < n; ++i) {
            int a = rating[i];
            update(tr2, a, -1);
            ans += query(tr1, a - 1) * (query(tr2, N - 1) - query(tr2, a));
            ans += (query(tr1, N - 1) - query(tr1, a)) * query(tr2, a - 1);
            update(tr1, a, 1);
        }
        return ans;
    }
};