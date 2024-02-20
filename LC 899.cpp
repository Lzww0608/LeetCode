class Solution {
public:
    string orderlyQueue(string s, int k) {
        if (k >= 2) {
            sort(s.begin(), s.end());
            return s;
        }
        string ans;
        char minChar = 'z';
        int n = s.length();
        for (int i = 0; i < n; ++i) {
            if (s[i] <= minChar) {
                string tmp = s.substr(i) + s.substr(0, i);
                if (ans == "" || ans > tmp) {
                    ans = tmp;
                }
            }
        }
        return ans;
    }
};