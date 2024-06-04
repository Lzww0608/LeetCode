class Solution {
public:
    int numberOfBoomerangs(vector<vector<int>>& points) {
        int ans = 0;
        auto dis = [](int a, int b, int c, int d) -> int {
            return (a - c) * (a - c) + (b - d) * (b - d);
        };
        for (int i = 0; i < points.size(); i++) {
            unordered_map<int,int>m;
            for (int j = 0; j < points.size(); j++) {
                if (i == j) continue;
                int d = dis(points[i][0], points[i][1], points[j][0], points[j][1]);
                ans += m[d] * 2;
                m[d]++;
            }
        }
        return ans;
    }
};
