class Solution {
    vector<vector<int>> rects;
public:
    Solution(vector<vector<int>>& rects) {
        this->rects = rects;
    }
    
    vector<int> pick() {
        int idx = -1, sum = 0, cur = 0, n = rects.size();

        for (int i = 0; i < n; ++i) {
            int a = rects[i][0], b = rects[i][1], x = rects[i][2], y = rects[i][3];
            cur = (x - a + 1) * (y - b + 1);
            sum += cur;
            if (rand() % sum < cur) {
                idx = i;
            }
        } 

        int a = rects[idx][0], b = rects[idx][1], x = rects[idx][2], y = rects[idx][3];
        return vector<int>{a + rand() % (x - a + 1), b + rand() % (y - b + 1)};
    }
};

/**
 * Your Solution object will be instantiated and called as such:
 * Solution* obj = new Solution(rects);
 * vector<int> param_1 = obj->pick();
 */