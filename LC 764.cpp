class Solution {
public:
    int orderOfLargestPlusSign(int n, vector<vector<int>>& mines) {
        vector<vector<int>> g(n + 1, vector<int>(n+1, 1)), l(n + 2, vector<int>(n+2, 0)), r(n + 2, vector<int>(n+2, 0)), u(n + 2, vector<int>(n+2, 0)), d(n + 2, vector<int>(n+2, 0));
        for (auto& mine : mines) {
            g[mine[0]+1][mine[1]+1] = 0;
        }
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= n; j++) {
                if (g[i][j] == 1) {
                    l[i][j] = l[i-1][j] + 1;
                    u[i][j] = u[i][j-1] + 1;
                }
                if (g[n-i+1][n-j+1] == 1) {
                    r[n-i+1][n-j+1] = r[n-i+2][n-j+1] + 1;
                    d[n-i+1][n-j+1] = d[n-i+1][n-j+2] + 1;
                }
            }
        }

        int ans = 0;
        for (int i = 1; i <= n; ++i) 
            for (int j = 1; j <= n; ++j)
            ans = max(ans, min({l[i][j], r[i][j], u[i][j], d[i][j]}));
        return ans;
    }
};