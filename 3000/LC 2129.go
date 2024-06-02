class Solution {
public:
    string capitalizeTitle(string title) {
        istringstream iss(title);

        string ans, s;

        while(iss >> s) {
            if (s.length() <= 2) {
                for (char& c : s) {
                    c = tolower(c);
                }
            } else {
                s[0] = toupper(s[0]);

                for (int i = 1; i < s.length(); i++) {
                    s[i] = tolower(s[i]);
                }
            }

            ans += s + " ";
        }
        ans.pop_back();
        return ans;
    }
};
