class Solution {
public:
    int atMostNGivenDigitSet(vector<string>& digits, int n) {
        auto s = to_string(n);
        int m = s.length(), dp[m];
        memset(dp, -1, sizeof(dp));
        function<int(int, bool, bool)> f = [&] (int i, bool is_limit, bool is_num) -> int {
            if (i >= m) return is_num;
            if (!is_limit && is_num && dp[i] >= 0) return dp[i];
            int res = 0;
            if (!is_num) {
                res = f(i + 1, false, false);
            }
            int up = is_limit ? s[i] - '0' : 9;
            for (const auto& t : digits) {
                int x = atoi(t.c_str());
                if (x <= up) {
                    res += f(i + 1, is_limit && x == up, true);
                } else {
                    break;
                }
            }
            if (!is_limit && is_num) {
                dp[i] = res;
            }
            return res;
        };
        return f(0, true, false);
    }
};