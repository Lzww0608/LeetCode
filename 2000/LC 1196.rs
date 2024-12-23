impl Solution {
    pub fn max_number_of_apples(mut weight: Vec<i32>) -> i32 {
        weight.sort_unstable();
        let mut ans = 0;
        let mut sum = 0;
        for &x in &weight {
            sum += x;
            if sum <= 5000 {
                ans += 1;
            } else {
                break;
            }
        }

        ans
    }
}