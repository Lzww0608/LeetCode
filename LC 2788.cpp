class Solution {
public:
    vector<string> splitWordsBySeparator(vector<string>& words, char separator) {
        vector<string> ans;
        for (auto& s : words) {
            string t;
            for (auto& c : s) {
                if (c == separator) {
                    if (t != "") {
                        ans.push_back(t);
                        t = "";
                    }
                } else {
                    t += c;
                }
            }
            if (t != "") {
                ans.push_back(t);
            }
        }
        return ans;
    }
};