class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        int n = nums.size();
        function<int(int, int)> quickSort = [&] (int i, int j) -> int {
            int pos = i, pvt = nums[i];
            int l = i, r = j;
            while (l < r) {
                while (l < r && nums[r] > pvt) {
                    r--;
                }
                while (l < r && nums[l] <= pvt) {
                    l++;
                } 
                swap(nums[l], nums[r]);
            }
            swap(nums[pos], nums[l]);
            if (l == n - k) {
                return nums[l];
            } else if (l < n - k) {
                return quickSort(l+1, j);
            } else {
                return quickSort(i, l-1);
            }
        };
        return quickSort(0, n-1);
    }
};