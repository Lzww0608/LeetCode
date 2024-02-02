class Solution {
public:
    int consecutiveNumbersSum(int n) {
        int count = 0;
        // The maximum possible value of k can be found when 2x + k - 1 = k, which simplifies to k(k+1)/2 <= n
        // Thus, iterate through all possible values of k to find valid combinations
        for (int k = 1; k <= std::sqrt(2 * n + 0.25) - 0.5; ++k) {
            // Check if 2*n is divisible by k
            if ((2 * n) % k == 0) {
                // (2*n/k - k + 1) must be even for x to be an integer
                if (((2 * n / k) - k + 1) % 2 == 0) {
                    count++;
                }
            }
        }
        return count;
    }
};