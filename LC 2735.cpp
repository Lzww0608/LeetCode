class Solution {
public:
    long long minCost(vector<int>& nums, int x) {
        long long ans = 0;
        vector<int>a(nums);
        for (int x : nums) ans += x;
        long long s = 0;
        for (int i = 0; i < nums.size(); i++) {
            s = 0;
            for (int j = 0; j < nums.size(); j++) {
                a[j] = min(a[j], nums[(j-i+nums.size())%nums.size()]);
                s += a[j];
            }
            ans = min(ans, (long long)x * i + s);
        }
        return ans;
    }
};