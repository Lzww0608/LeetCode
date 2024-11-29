impl Solution {
    pub fn can_alice_win(nums: Vec<i32>) -> bool {
        let (mut single, mut sum) = (0, 0);
        
        for &x in &nums {
            sum += x;
            if x < 10 {
                single += x;
            }
        }

        single * 2 != sum
    }
}