class Solution {
public:
    int kSimilarity(string a, string b) {
        string s1, s2;
        for (int i = 0; i < a.size(); i++) {
            if (a[i] != b[i]) {
                s1 += a[i];
                s2 += b[i];
            }
        }
        if (s1 == s2) return 0;
        unordered_map<string, bool> m;
        queue<pair<string, int>> q;
        q.push({s1, 0});
        int n = s1.size();
        while (!q.empty()) {
            for (int p = q.size(); p > 0; p--) {
                auto [str, k] = q.front();
                q.pop();
                if (str == s2) return k;
                int i = 0;
                while (i < str.size() && str[i] == s2[i]) i++;
                for (int j = i + 1; j < str.size(); j++) {
                    if (str[j] == s2[j] || str[j] != s2[i]) continue;
                    string next = str;
                    swap(next[i], next[j]);
                    if (!m[next]) {
                        m[next] = true;
                        q.push({next, k+1});
                    }
                }
            }
        }
        return -1;
    }
};

