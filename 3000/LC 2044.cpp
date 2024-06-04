class Solution {
public:
    int countMaxOrSubsets(vector<int>& nums) {
        int n = nums.size();
        int cnt = 0, maxVal = 0;
        function<void(int,int)> dfs = [&] (int idx, int orVal) {
            if (idx == n) {
                if (orVal > maxVal) {
                    maxVal = orVal;
                    cnt = 1;
                } else if (orVal == maxVal) {
                    cnt++;
                }
                return;
            }
            dfs(idx + 1, orVal | nums[idx]);
            dfs(idx + 1, orVal);
        };
        dfs(0,0);
        return cnt;
    }
};
