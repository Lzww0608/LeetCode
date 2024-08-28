impl Solution {
    pub fn min_operations(mut n: i32) -> i32 {
        let mut ans: i32 = 0;
        while n > 0 {
            if n & 1 == 1 {
                ans += 1;
                if n & 2 == 2 {
                    n += 1;
                }
            }
            n >>= 1;
        }

        ans
    }
}