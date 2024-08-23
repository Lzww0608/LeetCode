impl Solution {
    pub fn rearrange_array(nums: Vec<i32>) -> Vec<i32> {
        let mut n = nums.len();
        let mut ans = vec![0; n];
        let mut i = 0;
        let mut j = 1;
        for &x in &nums {
            if x > 0 {
                ans[i] = x;
                i += 2;
            } else {
                ans[j] = x;
                j += 2;
            }
        }

        ans
    }
}