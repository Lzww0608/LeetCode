impl Solution {
    pub fn two_egg_drop(n: i32) -> i32 {
        let mut ans = 0;
        let mut sum = 0;
        let mut i = 1;
        while sum < n {
            sum += i;
            ans += 1;
            i += 1;
        }

        ans
    }
}