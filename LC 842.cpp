class Solution {
public:
    vector<int> splitIntoFibonacci(string num) {
        vector<int> ans, path;
        int n = num.length();
        function<bool(long long,long long,int)> dfs = [&] (long long a, long long b, int idx) -> bool {
            if (idx >= n) {
                if (path.size() >= 3) {
                    ans.swap(path);
                    return true;
                }
                return false;
            }
            long long t = 0;
            for (int i = idx; i < n; ++i) {
                if (i > idx && num[idx] == '0') break;
                t = t * 10 + int(num[i] - '0');
                if (t > INT_MAX) break;
                if (path.size() < 2 || a + b == t) {
                    path.push_back(t);
                    if (dfs(b, t, i + 1))
                        return true;
                    path.pop_back();
                } 
            }
            return false;
        };
        dfs(0, 0, 0);
        return ans;
    }
};