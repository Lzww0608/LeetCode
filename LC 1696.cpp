class Solution {
public:
    int maxResult(vector<int>& nums, int k) {
        int n = nums.size();
        int ans = 0;
        vector<int> dp(n, INT_MIN);
        dp[0] = nums[0];
        deque<int> dq{0};
        for (int i = 1; i < n; ++i) {
            while (i - dq.front() > k) {
                dq.pop_front();
            }
            dp[i] = dp[dq.front()] + nums[i];
            while (!dq.empty() && dp[dq.back()] <= dp[i]) {
                dq.pop_back();
            }
            dq.push_back(i);
        }
        return dp[n-1];
    }
};