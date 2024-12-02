impl Solution {
    pub fn check_two_chessboards(coordinate1: String, coordinate2: String) -> bool {
        let x = coordinate1.chars().nth(0).unwrap() as i32 - 'a' as i32 + coordinate1.chars().nth(1).unwrap() as i32 - '0' as i32;
        let y = coordinate2.chars().nth(0).unwrap() as i32 - 'a' as i32 + coordinate2.chars().nth(1).unwrap() as i32 - '0' as i32;

        (x & 1) == (y & 1)
    }
}