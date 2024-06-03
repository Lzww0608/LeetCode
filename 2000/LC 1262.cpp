class Solution {
public:
    int maxSumDivThree(vector<int>& nums) {
        vector<int>dp(3,INT_MIN);
        dp[0] = 0;
        for (int x : nums) {
            vector<int>f{dp[0],dp[1],dp[2]};
            dp[x % 3] = max(x + f[0], f[x % 3]);
            dp[(1 + x) % 3] = max(x + f[1], f[(1 + x) % 3]);
            dp[(2 + x) % 3] = max(x + f[2], f[(2 + x) % 3]);
        }
        return dp[0];
    }
};
