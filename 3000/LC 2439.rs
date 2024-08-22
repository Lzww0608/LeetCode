impl Solution {
    pub fn minimize_array_value(nums: Vec<i32>) -> i32 {
        let mut s: i64 = 0;
        let mut ans: i64 = 0;
        for (i, &x) in nums.iter().enumerate() {
            s += x as i64;
            ans = ans.max((s + i as i64) / (i + 1) as i64)
        }

        ans as i32
    }
}