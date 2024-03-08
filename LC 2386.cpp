class Solution {
    using ll = long long;
public:
    long long kSum(vector<int>& nums, int k) {
        int n = nums.size();
        ll sum = 0;
        for (int& x : nums) {
            if (x >= 0) {
                sum += x;
            } else {
                x = -x;
            }
        }

        ranges::sort(nums);
        priority_queue<pair<ll, int>, vector<pair<ll, int>>, greater<>> pq;
        pq.emplace(nums[0], 0);
        ll ans = sum;
        for (int i = 1; i < k; ++i) {
            auto [s, j] = pq.top();
            pq.pop();
            ans = sum - s;
            if (j + 1 == n)
                continue;
            pq.emplace(s + nums[j+1], j + 1);
            pq.emplace(s + nums[j+1] - nums[j], j + 1);
        }
        return ans;
    }
};