class Solution {
    std::vector<int> prefix;
public:
    Solution(vector<int>& w) {
        std::ios::sync_with_stdio(false);
        std::cin.tie(nullptr);

        prefix = w;
        for (int i = 1; i < prefix.size(); ++i) {
            prefix[i] += prefix[i-1];
        }
    }
    
    int pickIndex() {
        int x = std::rand() % (prefix.back()) + 1;
        return std::distance(prefix.begin(), std::lower_bound(prefix.begin(), prefix.end(), x));
    }
};

/**
 * Your Solution object will be instantiated and called as such:
 * Solution* obj = new Solution(w);
 * int param_1 = obj->pickIndex();
 */