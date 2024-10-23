impl Solution {
    pub fn find_winning_player(skills: Vec<i32>, k: i32) -> i32 {
        let mut x = 0;
        let mut cnt = 0;

        for i in 1..skills.len() {
            if skills[x] < skills[i] {
                x = i;
                cnt = 1;
            } else {
                cnt += 1;
            }
            if cnt >= k {
                 break
            }
        }

        x as i32
    }
}