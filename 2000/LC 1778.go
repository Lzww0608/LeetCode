/**
 * // This is the GridMaster's API interface.
 * // You should not implement it, or speculate about its implementation
 * class GridMaster {
 *   public:
 *     bool canMove(char direction);
 *     void move(char direction);
 *     boolean isTarget();
 * };
 */

class Solution {
    constexpr static int N = 1001;
public:
    int findShortestPath(GridMaster &master) {
        std::array<std::array<char, N>, N> canReach{};
        std::array<std::array<char, N>, N> vis{};
        std::array<std::array<int, 2>, 4> dirs = {{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}};
        std::array<char, 4> ch = {{'D', 'U', 'R', 'L'}};
        std::array<char, 4> antich = {{'U', 'D', 'L', 'R'}};
        int end_x = -1, end_y = -1;
        auto dfs = [&] (this auto&& dfs, GridMaster& master, int i, int j) {
            if (canReach[i][j]) {
                return;
            }
            canReach[i][j] = true;
            if (master.isTarget()) {
                end_x = i;
                end_y = j;
            }
            for (int k = 0; k < 4; k++) {
                if (master.canMove(ch[k])) {
                    int x = i + dirs[k][0], y = j + dirs[k][1];
                    master.move(ch[k]);
                    dfs(master, x, y);
                    master.move(antich[k]);
                }
            }
        };
        dfs(master, 500, 500);

        if (end_x == -1) return -1;

        std::queue<std::pair<int, int>> q;
        q.push({500, 500});
        int d = 0;
        vis[500][500] = true;
        while (!q.empty()) {
            for (int sz = q.size(); sz > 0; sz--) {
                auto [i, j] = q.front();
                if (i == end_x && j == end_y) return d;
                q.pop();
                for (auto& dir : dirs) {
                    int x = i + dir[0], y = j + dir[1];
                    if (canReach[x][y] && !vis[x][y]) {
                        vis[x][y] = true;
                        q.emplace(x, y);
                    }
                }
            }
            d++;
        }

        return -1;
    }
};