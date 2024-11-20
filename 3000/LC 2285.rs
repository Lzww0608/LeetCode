impl Solution {
    pub fn maximum_importance(n: i32,roads: Vec<Vec<i32>>) -> i64 {
        let mut ans: i64 = 0;
        let n = n as usize;
        let mut cnt = vec![0; n];
        for v in &roads {
            cnt[v[0] as usize] += 1;
            cnt[v[1] as usize] += 1;
        }

        cnt.sort_unstable();
        for (i, &x) in cnt.iter().enumerate() {
            ans += ((i + 1) as i64 * x) as i64;
        }

        ans
    }
}