class Solution {
public:
    double mincostToHireWorkers(vector<int>& quality, vector<int>& wage, int k) {
        double ans = 0x7f7f7f7f;
        int n = quality.size();
        vector<pair<double, int>> workers;
        for (int i = 0; i < n; ++i) {
            workers.emplace_back(1.0 * wage[i] / quality[i], quality[i]);
        }
        sort(workers.begin(), workers.end());
        priority_queue<int> pq;
        int sum = 0;
        for (int i = 0; i < n; i++) {
            int t = workers[i].second;
            sum += t;
            pq.push(t);
            while (pq.size() > k) {
                sum -= pq.top();
                pq.pop();
            }
            if (pq.size() == k) {
                ans = min(ans, sum * workers[i].first);
            }
        }
        return ans;
    }
};