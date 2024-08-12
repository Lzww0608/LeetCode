impl Solution {
    pub fn is_array_special(nums: Vec<i32>, queries: Vec<Vec<i32>>) -> Vec<bool> {
        let mut n = nums.len();
        let mut f = vec![0; n];

        for i in 0..n-1 {
            if (nums[i] + nums[i+1]) & 1 == 1 {
                f[i+1] = f[i] + 1;
            } else {
                f[i+1] = f[i];
            }
        }

        let mut ans = vec![false; queries.len()];
        for (i, q) in queries.iter().enumerate() {
            let (l, r) = (q[0] as usize, q[1] as usize);
            if f[r] - f[l] == r - l {
                ans[i] = true;
            }
        }

        ans
    }
}