class Solution {
public:
    int minimumTimeRequired(vector<int>& jobs, int k) {
        int n = jobs.size();
        vector<int> sum(1 << n);
        for (int i = 0; i < n; ++i) {
            for (int j = 0, bit = (1 << i); j < bit; ++j) {
                sum[bit|j] = sum[j] + jobs[i];
            }
        }

        vector<int> dp(sum);
        for (int i = 1; i < k; ++i) {
            for (int j = (1 << n) - 1; j > 0; --j) {
                for (int s = j; s > 0; s = (s - 1) & j) {
                    dp[j] = min(dp[j], max(sum[s], dp[j ^ s]));
                }
            }
        }
        return dp.back();
    }
};