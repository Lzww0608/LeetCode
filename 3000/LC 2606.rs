const N: usize = 26;
impl Solution {
    pub fn maximum_cost_substring(s: String, chars: String, vals: Vec<i32>) -> i32 {
        let s = s.as_bytes();
        let chars = chars.as_bytes();
        let mut m: Vec<i32> = (1..=26).collect();
        let n = chars.len();

        for i in 0..n {
            let x = chars[i] - b'a';
            m[x as usize] = vals[i];
        }

        let mut ans = 0;
        let mut cur = 0;
        for &c in s {
            let x = c - b'a';
            cur = std::cmp::max(0, cur + m[x as usize]);
            ans = std::cmp::max(cur, ans);
        }

        ans
    }
}