class Solution {
public:
    int shortestPathLength(vector<vector<int>>& graph) {
        int n = graph.size();
        vector<vector<int>> vis(n, vector<int> (1 << n, 0));
        queue<tuple<int, int, int>> q;
        for (int i = 0; i < n; ++i) {
            q.emplace(i, 1 << i, 0);
            vis[i][1<<i] = 1;
        }
        int ans = 0;
        while (!q.empty()) {
            auto [u, mask, dis] = q.front();
            q.pop();
            if (mask == (1 << n) - 1) {
                ans = dis;
                break;
            }
            for (auto v : graph[u]) {
                int mask_v = mask | (1 << v);
                if (!vis[v][mask_v]) {
                    vis[v][mask_v] = 1;
                    q.push({v, mask_v, dis + 1});
                }
            }
        }
        return ans;
    }
};