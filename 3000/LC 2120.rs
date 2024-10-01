impl Solution {
    pub fn execute_instructions(n: i32, start_pos: Vec<i32>, s: String) -> Vec<i32> {
        let mut ans = vec![0; s.len()];

        for i in 0..s.len() {
            let mut j = i;
            let (mut x, mut y) = (start_pos[0], start_pos[1]);
            while j < s.len() {
                match s.chars().nth(j).unwrap() {
                    'R' => y += 1,
                    'L' => y -= 1,
                    'U' => x -= 1,
                    'D' => x += 1,
                    _=>(),
                }

                if x < 0 || x >= n || y < 0 || y >= n {
                    break;
                }
                j += 1;
            }
            ans[i] = (j - i) as i32;
        }

        ans
    }
}