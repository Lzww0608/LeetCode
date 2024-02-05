class Solution {
public:
    int magicTower(vector<int>& nums) {
        long long sum = 1;
        priority_queue<int, vector<int>, greater<int>> pq;
        int ans = 0;
        if (accumulate(nums.begin(), nums.end(), (long long)0) < 0) return -1;
        for (int i = 0; i < nums.size(); ++i) {
            if (nums[i] < 0) {
                pq.push(nums[i]);
            }
            if (sum + nums[i] <= 0) {
                auto t = pq.top();
                pq.pop();
                sum += nums[i] - t;
                nums.push_back(t);
                ans++;
            } else {
                sum += nums[i];
            }
        }
        return ans;
    }
};