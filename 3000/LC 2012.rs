impl Solution {
    pub fn sum_of_beauties(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut pre = vec![false; n];
        let mut cur = 0;
        let mut ans = 0;

        for i in 0..n {
            if i > 0 && i < n - 1 && nums[i] > nums[i-1] && nums[i] < nums[i+1] {
                ans += 1;
            }

            if cur < nums[i] {
                pre[i] = true;
                cur = nums[i];
            }
        }

        cur = nums[n-1];
        for i in (1..n).rev() {
            if cur > nums[i] {
                cur = nums[i];
                if pre[i] {
                    ans += 1;
                }
            }
        }

        ans
    }
}