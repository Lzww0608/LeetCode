class Solution {
public:
    int uniqueLetterString(string s) {
        int n = s.length();
        vector<int> last0(26, -1), last1(26, -1);
        int ans = 0, cnt = 0;
        for (int i = 0; i < n; ++i) {
            int c = s[i] - 'A';
            cnt += i - last0[c] - (last0[c] - last1[c]);
            ans += cnt;
            last1[c] = last0[c];
            last0[c] = i;
        }
        return ans;
    }
};