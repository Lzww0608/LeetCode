class Solution {
public:
    int minRefuelStops(int target, int startFuel, vector<vector<int>>& stations) {
        if (target <= startFuel) return 0;
        int ans = 0;
        priority_queue<int> pq;
        int n = stations.size(), dis = startFuel;
        int i = 0;
        while (dis < target) {
            while (i < n && stations[i][0] <= dis) {
                pq.push(stations[i][1]);
                i++;
            }
            if (pq.empty()) 
                return -1;
            dis += pq.top();
            pq.pop();
            ans++;
        }
        return ans;
    }
};