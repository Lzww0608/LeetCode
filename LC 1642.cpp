class Solution {
public:
    int furthestBuilding(vector<int>& heights, int bricks, int ladders) {
        int n = heights.size();
        priority_queue<int, vector<int>, greater<int>> pq;
        int sum = 0;

        for (int i = 1; i < n; ++i) {
                if (heights[i] <= heights[i-1]) 
                    continue;
                pq.push(heights[i] - heights[i-1]);
                if (pq.size() > ladders) {
                    sum += pq.top();
                    pq.pop();
                }
                if (sum > bricks) return i - 1;
        
        }

        return n - 1;        
    }
};