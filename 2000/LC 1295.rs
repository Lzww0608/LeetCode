impl Solution {
    pub fn find_numbers(nums: Vec<i32>) -> i32 {
        let mut ans: i32 = 0;

        for &x in &nums {
            if x.to_string().len() & 1 == 0 {
                ans += 1;
            }
        }

        ans
    }
}