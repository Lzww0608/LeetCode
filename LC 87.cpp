class Solution {
    unordered_map<string, bool> m;
public:
    bool isScramble(string s1, string s2) {
        string key = s1 + ", " + s2;
        if (m.contains(key)) {
            return m[key];
        }
        if (s1 == s2) {
            m[key] = true;
            return true;
        } 
        auto a = s1, b = s2;
        sort(a.begin(), a.end());
        sort(b.begin(), b.end());
        if (a != b) {
            return false;
        }
        int n = s1.length();
        for (int i = 1; i < n; ++i) {
            if (isScramble(s1.substr(0, i), s2.substr(0, i)) && isScramble(s1.substr(i), s2.substr(i))) {
                m[key] = true;
                return true;
            } else if (isScramble(s1.substr(0, i), s2.substr(n-i)) && isScramble(s1.substr(i), s2.substr(0, n-i))) {
                m[key] = true;
                return true;
            }
        }
        m[key] = false;
        return false;
    }
};