const MOD: i64 = 1_000_000_007;
impl Solution {
    pub fn subarrays_with_more_zeros_than_ones(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut cnts = vec![0; n * 2 + 1];
        cnts[n] = 1;
        let mut pre = n as i32;
        let mut cnt = 0;
        let mut ans: i64 = 0;
        for i in 0..n {
            if nums[i] == 1 {
                cnt += cnts[pre as usize];
                pre += 1;
            } else {
                cnt -= cnts[pre as usize - 1];
                pre -= 1;
            }
            cnts[pre as usize] += 1;
            ans = (ans + cnt as i64) % MOD
        }

        ans as _
    }
}