impl Solution {
    pub fn min_operations(k: i32) -> i32 {
        let mut ans = k;
        for i in 1..=k {
            ans = std::cmp::min(ans, i - 1 + (k - 1) / i);
        }

        ans
    }
}