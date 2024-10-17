impl Solution {
    pub fn min_operations(mut s: Vec<i32>) -> i32 {
        let n = s.len();
        let mut ans = 0;

        for i in 0..(n-2) {
            if s[i] == 0 {
                s[i+1] ^= 1;
                s[i+2] ^= 1;
                ans += 1;
            }
        }

        if s[n-1] == 0 || s[n-2] == 0 {
            return -1;
        }

        ans
    }
}