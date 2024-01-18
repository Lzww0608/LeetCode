class Solution {
    int dir[4][2] = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
public:
    //BFS
    int longestIncreasingPath(vector<vector<int>>& matrix) {
        int n = matrix.size(), m = matrix[0].size();
        vector<vector<int>> degree(n, vector<int>(m, 0));
        for (int i = 0; i < n; ++i) {
            for (int j = 0; j < m; ++j) {
                for (int k = 0; k < 4; ++k) {
                    int x = i + dir[k][0], y = j + dir[k][1];
                    if (x >= 0 && x < n && y >= 0 && y < m && matrix[x][y] > matrix[i][j]) {
                        degree[i][j]++;
                    }
                }
            }
        }

        queue<pair<int,int>>q;
        for (int i = 0; i < n; ++i) {
            for (int j = 0; j < m; ++j) {
                if (degree[i][j] == 0) {
                    q.push({i,j});
                }
            }
        }

        int ans = 0;
        while (!q.empty()) {
            ++ans;
            for (int t = q.size(); t > 0; --t) {
                int i = q.front().first, j = q.front().second;
                q.pop();
                for (int k = 0; k < 4; ++k) {
                    int x = i + dir[k][0], y = j + dir[k][1];
                    if (x >= 0 && x < n && y >= 0 && y < m && matrix[x][y] < matrix[i][j]) {
                        --degree[x][y];
                        if (degree[x][y] == 0) {
                            q.push({x, y});
                        }
                    }
                }
            }
        }

        return ans;
    }
};
