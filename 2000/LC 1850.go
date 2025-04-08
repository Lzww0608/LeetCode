class Solution {
public:
    int getMinSwaps(string num, int k) {
        std::string num_k = num;
        for (int i = 0; i < k; i++) {
            std::next_permutation(num_k.begin(), num_k.end());
        }

        int n = num.size(), ans = 0;
        for (int i = 0; i < n; i++) {
            if (num[i] != num_k[i]) {
                for (int j = i + 1; j < n; j++) {
                    if (num[j] == num_k[i]) {
                        for (int t = j - 1; t >= i; t--) {
                            ans++;
                            std::swap(num[t], num[t+1]);
                        }
                        break;
                    }
                }
            }
        }

        return ans;
    }
};