impl Solution {
    pub fn make_integer_beautiful(mut n: i64, target: i32) -> i64 {
        let check = |n: i64| -> bool {
            let mut cnt = 0;
            let mut t = n;
            while t > 0 {
                cnt += (t % 10) as i32;
                t /= 10
            }

            cnt <= target
        };

        let mut ans = 0;
        let mut i = 10i64;
        while !check(n) {
            ans += (10 - n % 10) * (i / 10) as i64;
            n = n / 10 + 1;
            i *= 10;
        }

        ans
    }
}