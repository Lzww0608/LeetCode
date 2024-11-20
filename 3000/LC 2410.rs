impl Solution {
    pub fn match_players_and_trainers(mut a: Vec<i32>, mut b: Vec<i32>) -> i32 {
        a.sort_unstable();
        b.sort_unstable();

        let mut j = 0;
        let mut ans = 0;
        for i in 0..a.len() {
            while j < b.len() && b[j] < a[i] {
                j += 1;
            }

            if j >= b.len() {
                break;
            }
            j += 1;
            ans += 1;
        }

        ans
    }
}