class Solution {
public:
    vector<int> processQueries(vector<int>& queries, int m) {
        int n = queries.size();
        vector<int> pos(m), ans;
        iota(pos.begin(), pos.end(), 1);
        for (int x : queries) {
            auto it = find(pos.begin(), pos.end(), x);
            ans.push_back(it - pos.begin());
            int t = *it;
            pos.erase(it);
            pos.insert(pos.begin(), t);
        }
        return ans;
    }
};
