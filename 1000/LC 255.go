class Solution {
public:
    bool verifyPreorder(vector<int>& preorder) {
        std::ios::sync_with_stdio(false);
        std::cin.tie(nullptr);

        std::stack<int> st;
        int mx = INT_MIN;
        for (int x : preorder) {
            while (!st.empty() && x > st.top()) {
                mx = std::max(mx, st.top());
                st.pop();
            }
            if (x < mx) return false;
            st.push(x);
        }

        return true;
    }
};