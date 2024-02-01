class Solution {
    int mod = 1e9 + 7;
    priority_queue<int> pq1;
    priority_queue<int, vector<int>, greater<int>> pq2;
public:
    vector<int> numsGame(vector<int>& nums) {
        int n = nums.size();
        vector<int> ans(n, 0);
        long long sum1 = 0, sum2= 0;
        for (int i = 0; i < n; ++i) {
            nums[i] -= i;
            int x = nums[i];
            if (pq1.empty() || x < pq1.top()) {
                pq1.push(x);
                sum1 += x;
            } else {
                pq2.push(x);
                sum2 += x;
            }
            while (pq1.size() > pq2.size() + 1) {
                auto t = pq1.top();
                pq1.pop();
                pq2.push(t);
                sum1 -= t;
                sum2 += t;
            }
            while (pq1.size() < pq2.size()) {
                auto t = pq2.top();
                pq2.pop();
                pq1.push(t);
                sum2 -= t;
                sum1 += t;
            }
            if (i & 1) {
                ans[i] = abs(sum2 - sum1) % mod;
            } else {
                ans[i] = abs(sum2 - sum1 + pq1.top()) % mod;
            }
        }
        return ans;
    }
};