class Solution {
public:
    vector<string> ambiguousCoordinates(string s) {
        int n = s.size() - 2;
        s = s.substr(1, n);
        vector<string>ans;
        auto getPos = [&] (string s) -> vector<string> {
            vector<string> res;
            if (s[0] != '0' || s == "0") res.push_back(s);
            for (int p = 1; p < s.size(); p++) {
                if ((p != 1 && s[0] == '0') || s.back() == '0') continue;
                res.push_back(s.substr(0,p) + "." + s.substr(p));
            }
            return res;
        };
        for (int i = 1; i < n; i++) {
            auto a = s.substr(0, i);
            auto l = getPos(a);
            auto b = s.substr(i);
            auto r = getPos(b);
            for (auto& p : l) {
                for (auto& q : r) {
                    ans.push_back("(" + p + ", " + q + ")");
                }
            }
        }
        return ans;
    }
};