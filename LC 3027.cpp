class Solution {
public:
    int numberOfPairs(vector<vector<int>>& points) {
        ranges::sort(points, [](const auto& a, const auto& b) {
            if (a[0] == b[0]) return a[1] > b[1];
            return a[0] < b[0];
        }); 
        int ans = 0;
        int n = points.size();
        for (int i = 0; i < n; ++i) {
            int a = points[i][1];
            int max_y = INT_MIN;
            for (int j = i + 1; j < n; ++j) {
                int b = points[j][1];
                if (b <= a && b > max_y) {
                    ans++;
                    max_y = b;
                }
            }
        }
        return ans;
    }
};