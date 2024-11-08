impl Solution {
    pub fn max_sum_range_query(mut nums: Vec<i32>, requests: Vec<Vec<i32>>) -> i32 {
        let MOD: i64 = 1_000_000_007;
        let mut ans: i64 = 0;
        let n = nums.len();
        let mut pre = vec![0; n + 1];
        let mut freq = vec![0; n];

        for v in &requests {
            let a = v[0] as usize;
            let b = v[1] as usize;
            pre[a] += 1;
            pre[b+1] -= 1;
        } 

        let mut cur = 0;
        for i in 0..n {
            cur += pre[i];
            freq[i] = cur;
        }

        freq.sort_unstable_by(|a, b| b.cmp(a));
        nums.sort_unstable_by(|a, b| b.cmp(a));

        for i in 0..n {
            if freq[i] == 0 {
                break;
            }
            ans = (ans + freq[i] as i64 * nums[i] as i64) % MOD;
        }

        ans as _

    }
}