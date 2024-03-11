class Solution {
public:
    vector<int> secondGreaterElement(vector<int>& a) {
        int n = a.size();
        vector<int> ans(n, -1);

        priority_queue<pair<int, int>, vector<pair<int, int>>, greater<pair<int, int>>> pq;
        stack<int> st;
        for (int i = 0; i < n; ++i) {
            while (pq.size() > 0 && pq.top().first < a[i]) {
                ans[pq.top().second] = a[i];
                pq.pop();
            }

            while (!st.empty() && a[st.top()] < a[i]) {
                pq.push({a[st.top()], st.top()});
                st.pop();
            }

            st.push(i);
        }
        return ans;
    }
};