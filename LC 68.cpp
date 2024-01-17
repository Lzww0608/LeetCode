class Solution {
public:
    vector<string> fullJustify(vector<string>& words, int maxWidth) {
        int n = words.size();
        vector<string>ans, tmp;
        int d = 0;
        for (auto& s : words) {
            if (d + s.length() <= maxWidth) {
                d += s.length() + 1;
                tmp.push_back(s);
            } else {
                int m = tmp.size();
                d -= m;
                int vel = m == 1 ? 0 : (maxWidth - d) / (m - 1);
                int first = m == 1 ? 0 : (maxWidth - d) % (m - 1);
                string res = tmp[0];
                for (int j = 0; j < vel + min(1, first); ++j) 
                    res += " ";
                first = max(0, first - 1);
                for (size_t i = 1; i < m; ++i) {  // 从第二个单词开始
                    res += tmp[i];
                    if (i + 1 != m) {
                        for (int j = 0; j < vel + min(1, first); ++j)
                            res += " ";
                        first = max(0, first - 1);
                    }
                }
                while (res.length() < maxWidth) {
                    res += " ";
                }
                ans.push_back(res);
                d = s.length() + 1;
                tmp.clear();
                tmp.push_back(s);
            }
        }
        string res = "";
        for (auto& t : tmp) {
            res += t + " ";
        }
        while (res.length() > maxWidth) {
            res.pop_back();
        }
        while (res.length() < maxWidth) {
            res += " ";
        }
        ans.push_back(res);
        return ans;
    }
};