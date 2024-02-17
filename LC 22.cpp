class Solution {
public:
    vector<string> generateParenthesis(int n) {
        vector<string> ans;
        vector<vector<string>> dp(n + 1);
        dp[0] = {""};
        dp[1] = {"()"};
        for (int i = 2; i <= n; ++i) {
            for (int j = 0; j < i; ++j) {
                for (auto& p : dp[j]) {
                    for (auto& q : dp[i-j-1]) {
                        string s = "(" + p + ")" + q;
                        dp[i].emplace_back(s);
                    }
                }
            }
        }
        return dp[n];
    }
};