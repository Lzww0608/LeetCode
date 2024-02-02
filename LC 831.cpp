class Solution {
    string email(string& s) {
        string ans;
        int i = 0, n = s.size();
        while (s[i] != '@') {
            if (i == 0 || s[i+1] == '@') {
                if (s[i] >= 'A' && s[i] <= 'Z') {
                    ans += s[i] + 32;
                } else {
                    ans += s[i];
                }
                if (ans.length() < 5) 
                    ans += "*****";
            } 
            i++;
        }
        i++;
        ans += '@';
        while (i < n) {
            if (s[i] >= 'A' && s[i] <= 'Z'){
                ans += s[i] + 32;
            } else {
                ans += s[i];
            }
            i++;
        }
        return ans;
    }

    string phoneNumber(string& s) {
        string ans, local;
        int n = s.length(), cnt = 0;
        for (int i = n - 1; i >= 0; i--) {
            if (isdigit(s[i])) {
                if (local.length() < 4)
                    local += s[i];
                cnt++;
            }
        }
        reverse(local.begin(), local.end());
        if (cnt == 10) {
            ans += "***-***-";
        } else if (cnt == 11) {
            ans += "+*-***-***-";
        } else if (cnt == 12) {
            ans += "+**-***-***-";
        } else {
            ans += "+***-***-***-";
        }
        ans += local;
        return ans;
    }
public:
    string maskPII(string s) {
        for (auto& c : s) {
            if (c == '@') {
                return email(s);
            }
        }
        return phoneNumber(s);
    }
};