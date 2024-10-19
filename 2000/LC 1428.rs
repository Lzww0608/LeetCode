/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 *  struct BinaryMatrix;
 *  impl BinaryMatrix {
 *      fn get(&self, row: i32, col: i32) -> i32;
 *     fn dimensions() -> Vec<i32>;
 * };
 */

impl Solution {
    pub fn left_most_column_with_one(c: &BinaryMatrix) -> i32 {
        let t = c.dimensions();
        let m = t[0];
        let n = t[1];
        let (mut i, mut j) = (0, n - 1);

        let mut ans = n;
        while i < m && j >= 0 {
            if c.get(i, j) == 1 {
                ans = ans.min(j);
                j -= 1;
            } else {
                i += 1;
            }
        }

        if ans == n {
            return -1;
        }

        ans
    }
}