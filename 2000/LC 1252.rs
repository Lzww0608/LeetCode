impl Solution {
    pub fn odd_cells(m: i32, n: i32, indices: Vec<Vec<i32>>) -> i32 {
        let mut row = vec![0; m as usize];
        let mut col = vec![0; n as usize];
        let (mut odd, mut even) = (0, 0);
        for v in &indices {
            let a = v[0] as usize;
            let b = v[1] as usize;
            row[a] += 1;
            col[b] += 1;
        }

        for &x in &row {
            if x & 1 == 1 {
                odd += 1;
            } else {
                even += 1;
            }
        }

        let mut ans = 0;
        for &x in &col {
            if x & 1 == 1 {
                ans += even;
            } else {
                ans += odd;
            }
        }

        ans
    }
}