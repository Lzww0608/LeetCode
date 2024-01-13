class Solution {
public:
    string repeatLimitedString(string s, int repeatLimit) {
       vector<int>cnt(26);
       for (auto c : s) {
           cnt[c-'a']++;
       }
       string ans;
       int m = 0;
       for (int i = 25, j = 24; i >= 0 && j >= 0;) {
           if (cnt[i] == 0) {
               m = 0;
               i--;
           } else if (m < repeatLimit) {
               cnt[i]--;
               ans.push_back('a'+i);
               m++;
           } else if (j >= i || cnt[j] == 0) {
               j--;
           } else {
               cnt[j]--;
               ans.push_back('a'+j);
               m = 0;
           }
       }
       return ans;
    }
};