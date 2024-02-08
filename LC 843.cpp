/**
 * // This is the Master's API interface.
 * // You should not implement it, or speculate about its implementation
 * class Master {
 *   public:
 *     int guess(string word);
 * };
 */
class Solution {
    vector<vector<int>> sim;
    vector<int> sel;
    void cal(vector<string>& words) {
        int n = words.size();
        for (int i = 0; i < n; ++i) {
            sim[i][i] = 6;
            for (int j = i + 1; j < n; ++j) {
                int cnt = 0;
                for (int k = 0; k < 6; ++k) {
                    if (words[i][k] == words[j][k]) cnt++;
                }
                sim[i][j] = sim[j][i] = cnt;
            }
        }
    }

    int find() {
        int next = -1, maxVal = INT_MAX;
        int n = sel.size();
        for (int i = 0; i < n; ++i) {
            if (!sel[i]) continue;
            vector<int> cnt(6, 0);
            for (int j = 0; j < n; ++j) {
                if (j == i || !sel[j]) continue;
                ++cnt[sim[i][j]];
            }
            int cur = *max_element(cnt.begin(), cnt.end());
            if (cur < maxVal) {
                maxVal = cur;
                next = i;
            }
        }
        return next;
    }
public:
    void findSecretWord(vector<string>& words, Master& master) {
        int n = words.size();
        sim.resize(n, vector<int> (n, 0));
        sel.resize(n, 1);
        cal(words);
        while (words.size() > 0) {
            int next = find();
            auto s = words[next];
            int pos = master.guess(s);
            if (pos == 6) {
                return;
            }
            for (int i = 0; i < n; ++i) {
                if (sim[next][i] != pos) {
                    sel[i] = false;
                }
            }
        }
    }
};