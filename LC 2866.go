class Solution {
    using ll = long long;
public:
    long long maximumSumOfHeights(vector<int>& a) {
        int n = a.size();

        auto get = [&] (const vector<int>& arr) {
            vector<ll> f(n, 0);
            stack<int> st;

            for (int i = 0; i < n; i++) {
                while (!st.empty() && arr[st.top()] > arr[i]) {
                    st.pop();
                }

                if (!st.empty()) {
                    f[i] = f[st.top()] + ll(i - st.top()) * arr[i];
                } else {
                    f[i] = ll(i + 1) * arr[i];
                }
                st.push(i);
            }

            return f;
        };

        auto pre = get(a);
        reverse(a.begin(), a.end());
        auto suf = get(a);
        reverse(a.begin(), a.end());
        reverse(suf.begin(), suf.end());

        ll ans = 0;
        for (int i = 0; i < n; ++i) {
            ans = max(ans, pre[i] + suf[i] - a[i]);
        }
        return ans;
    }
};