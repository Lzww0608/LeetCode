class Solution {
    int dir[4][2] = {{0,1},{0,-1},{1,0},{-1,0}};
public:
    int shortestPathAllKeys(vector<string>& grid) {
        int m = grid.size(), n = grid[0].size();
        unordered_map<char, int> mp;
        queue<tuple<int,int,int>> q;
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (grid[i][j] == '@') {
                    q.push({i, j, 0});
                } else if (islower(grid[i][j])) {
                    if (!mp.contains(grid[i][j])) {
                        int t = mp.size();
                        mp[grid[i][j]] = t;
                    }
                }
            }
        }

        int k = mp.size();
        vector<vector<vector<int>>> dist(m, vector<vector<int>>(n, vector<int>(1 << k, -1)));
        auto [a, b, _] = q.front();
        dist[a][b][0] = 0;
        while (!q.empty()) {
            auto [i, j, mask] = q.front();
            q.pop();
            for (int t = 0; t < 4; t++) {
                int x = i + dir[t][0], y = j + dir[t][1];
                if (x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == '#')
                    continue;
                if (grid[x][y] == '.' || grid[x][y] == '@') {
                    if (dist[x][y][mask] == -1) {
                        dist[x][y][mask] = dist[i][j][mask] + 1;
                        q.emplace(x, y, mask);
                    }
                } else if (islower(grid[x][y])) {
                    int idx = mp[grid[x][y]];
                    if (dist[x][y][mask | (1 << idx)] == -1) {
                        dist[x][y][mask | (1 << idx)] = dist[i][j][mask] + 1;
                        if ((mask | (1 << idx)) == (1 << k) - 1) {
                            return dist[x][y][mask | (1 << idx)];
                        }
                        q.emplace(x, y, mask | (1 << idx));
                    }
                } else {
                    int idx = mp[tolower(grid[x][y])];
                    if ((mask & (1 << idx)) && dist[x][y][mask] == -1) {
                        dist[x][y][mask] = dist[i][j][mask] + 1;
                        q.emplace(x, y, mask);
                    }
                }
                
            }
        }

        
        
        return -1;
    }
};