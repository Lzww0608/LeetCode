class Solution {
public:
    int stoneGameVII(vector<int>& a) {
        int n = a.size();
        vector<int> sum(n + 1, 0);
        vector<int> dp(n);
        for (int i = 0; i < n; ++i) {
            sum[i+1] = sum[i] + a[i];
        }
        for (int i = n - 2; i >= 0; --i) {
            for (int j = i + 1; j < n; ++j) {
                dp[j] = max(sum[j] - sum[i] - dp[j-1], sum[j+1] - sum[i+1] - dp[j]);
            }
        }
        return dp[n-1];
    }
};
