class Solution {
public:
    vector<string> restoreIpAddresses(string s) {
        vector<string> ans;
        int n = s.length();
        string ip = "";
        function<void(int, int)> dfs = [&] (int idx, int num) {
            if (idx >= n) {
                if (num == 4) {
                    ip.pop_back();
                    ans.push_back(ip);
                }
                return;
            }
            int t = 0;
            for (int i = idx; i < n; i++) {
                t = t * 10 + int(s[i] - '0');
                if (num > 3 || t > 255 || (i > idx && s[idx] =='0')) break;
                string ss = ip;
                ip += to_string(t) + ".";
                dfs(i + 1, num + 1);
                ip = ss;
            }
        };
        dfs(0, 0);
        return ans;
    }
};