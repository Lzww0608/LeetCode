class Solution {
    int difs[10] = {0,0,1,-1,-1,1,1,-1,0,1};
public:
    //digital dp
    int rotatedDigits(int n) {
        auto s = to_string(n);
        int m = s.length(), dp[m][2];
        memset(dp, -1, sizeof(dp));
        function<int(int, bool, bool)> f = [&] (int i, bool has_dif, bool is_limit) ->int {
            if (i == m) return has_dif;
            if (!is_limit && dp[i][has_dif] >= 0) return dp[i][has_dif];
            int res = 0;
            int up = is_limit ? s[i] - '0' : 9;
            for (int d = 0; d <= up; ++d) {
                if (difs[d] != -1)
                    res += f(i + 1, has_dif || difs[d], is_limit && d == up);
            }
            if (!is_limit) {
                dp[i][has_dif] = res;
            }
            return res;
        };
        return f(0, false, true);
    }
};