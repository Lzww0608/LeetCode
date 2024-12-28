const N: usize = 101;
impl Solution {
    pub fn is_possible_to_split(nums: Vec<i32>) -> bool {
        let mut cnt = vec![0; N];
        for &x in &nums {
            cnt[x as usize] += 1;
            if cnt[x as usize] > 2 {
                return false;
            }
        }

        return true;
    }
}