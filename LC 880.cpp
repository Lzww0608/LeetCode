class Solution {
public:
    string decodeAtIndex(string s, int k) {
        string ans;
        long long n = s.length(), m = 0;
        int j = -1;
        for (int i = 0; i < n; ++i) {
            if (isalpha(s[i])) {
                m++;
                if (m == k) {
                    cout << s[i];
                    ans = s[i];
                    break;
                }
            } else {
                int t = s[i] - '0';
                m *= t;
                if (m >= k) {
                    j = i;
                    break;
                }
            }
        }
        for (; j >= 0; j--) {
            k %= m;
            if (k == 0 && isalpha(s[j]))
                return (string)"" + s[j];
            if (isdigit(s[j])) {
                m /= s[j] - '0';
            } else {
                m--;
            }
        }
        return ans;
    }
};