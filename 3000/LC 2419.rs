impl Solution {
    pub fn longest_subarray(nums: Vec<i32>) -> i32 {
        let mut ans = 1;
        let mut cnt = 0;
        let mut mx = 0;

        for &x in &nums {
            if x > mx {
                mx = x;
                cnt = 1;
                ans = 1;
            } else if x == mx {
                cnt += 1;
                ans = ans.max(cnt);
            } else {
                cnt = 0;

            }
        }

        ans
    }
}