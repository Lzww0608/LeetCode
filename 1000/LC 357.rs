impl Solution {
    pub fn count_numbers_with_unique_digits(n: i32) -> i32 {
        if n == 0 {
            return 1
        }

        let mut ans = 10;
        let mut cur = 9;
        for i in 2..=n {
            cur *= 10 - i + 1;
            ans += cur;
        }

        ans
    }
}