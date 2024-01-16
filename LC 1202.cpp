class Solution {
public:
    string smallestStringWithSwaps(string s, vector<vector<int>>& pairs) {
        int m = s.length();
        vector<int>color(m,0);
        vector<vector<int>>g(m);
        for (const auto& v : pairs) {
            g[v[0]].push_back(v[1]);
            g[v[1]].push_back(v[0]);
        }

        auto sortString = [&] (vector<int>& v) {
            sort(v.begin(), v.end());
            vector<char>tmp;
            for (int x : v) {
                tmp.push_back(s[x]);
            }
            sort(tmp.begin(), tmp.end());
            for (size_t i = 0; i < v.size(); ++i) {
                s[v[i]] = tmp[i];
            }
        };

        function<void(int, vector<int>&)> dfs = [&] (int i, vector<int>& tmp) {
            if (color[i]) return;
            color[i] = 1;
            tmp.push_back(i);
            for (auto k : g[i]) {
                if (color[k] == 0)
                    dfs(k, tmp);
            } 
        };

        for (int i = 0; i < m; ++i) {
            if (color[i] == 0 && g[i].size() > 0) {
                vector<int>tmp;
                dfs(i, tmp);
                sortString(tmp);
            }
        }
        return s;
    }
};