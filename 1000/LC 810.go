impl Solution {
    pub fn xor_game(nums: Vec<i32>) -> bool {
        let mut n = nums.len();
        if n & 1 == 0 {
            return true;
        }
        let mut x: i32 = 0;
        for i in 0..n {
            x ^= nums[i];
        }

        x == 0
    }
}