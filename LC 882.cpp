using pii = pair<int, int>;
class Solution {
    vector<int> dijkstra(vector<vector<pii>>& g, int start) {
        int n = g.size();
        vector<int> dist(n, 0x7f7f7f7f), visited(n, 0);
        priority_queue<pii, vector<pii>, greater<pii>> pq;
        pq.push(make_pair(0, start));
        dist[start] = 0;
        while (!pq.empty()) {
            int u = pq.top().second;
            pq.pop();
            if (visited[u]) continue;
            visited[u] = 1;
            for (auto& edge : g[u]) {
                int v= edge.first, weight = edge.second;
                if (!visited[v] && dist[u] + weight < dist[v]) {
                    dist[v] = dist[u] + weight;
                    pq.push(make_pair(dist[v], v));
                }
            }
        }
        return dist;
    }
public:
    int reachableNodes(vector<vector<int>>& edges, int maxMoves, int n) {
        vector<vector<pii>> g(n);
        for (auto& edge : edges) {
            int a = edge[0], b = edge[1], c = edge[2];
            g[a].push_back(make_pair(b, c + 1));
            g[b].push_back(make_pair(a, c + 1));
        }
        auto dist = dijkstra(g, 0);
        int ans = 0;
        for (int x : dist) {
            if (x <= maxMoves)
                ++ans;
        }
        for (auto& e : edges) {
            int a = e[0], b = e[1], c = e[2];
            int x = max(0, maxMoves - dist[a]);
            int y = max(0, maxMoves - dist[b]);
            ans += min(x + y, c);
        }
        return ans;
    }
};