use std::cmp::Ordering;

impl Solution {
    pub fn min_operations(mut nums: Vec<i32>, queries: Vec<i32>) -> Vec<i64> {
        nums.sort();
        let n = nums.len();
        let mut pre = vec![0i64; n + 1];
        for (i, &x) in nums.iter().enumerate() {
            pre[i+1] = pre[i] + x as i64;
        }

        let mut ans = vec![0i64; queries.len()];
        
        for (i, &x) in queries.iter().enumerate() {
            let cnt = nums.binary_search_by(|&y| if y < x {Ordering::Less} else {Ordering::Greater}).unwrap_or_else(|e| e);
            let sum = pre[cnt];
            ans[i] = (x as i64 * cnt as i64 - sum) + (pre[n] as i64 - pre[cnt] as i64 - x as i64 * (n as i64 - cnt as i64));
        }

        ans
    }
}