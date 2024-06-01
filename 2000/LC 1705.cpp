
class Solution {
public:
    int eatenApples(vector<int>& apples, vector<int>& days) {
        priority_queue<pair<int,int>,vector<pair<int,int>>, greater<pair<int,int>>>pq;
        int n = apples.size(), ans = 0;
        for (int i = 0; i < n || !pq.empty(); i++) {
            if (i < n) {
                pq.push({i + days[i],apples[i]});
            }
            while (!pq.empty() && (pq.top().first <= i || pq.top().second < 1)) {
                pq.pop();
            }
            if (!pq.empty() && pq.top().first > i && pq.top().second >= 1) {
                auto t = pq.top();
                pq.pop();
                t.second -= 1;
                if (t.second > 0) {
                    pq.push(t);
                }
                ans++;
            }
        }
        return ans;
    }
};
