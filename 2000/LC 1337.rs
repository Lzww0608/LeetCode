impl Solution {
    pub fn k_weakest_rows(mat: Vec<Vec<i32>>, k: i32) -> Vec<i32> {
        let m = mat.len();
        let mut a: Vec<(i32, usize)> = Vec::with_capacity(m);

        for (i, row) in mat.iter().enumerate() {
            let cnt = row.iter().take_while(|&&x| x == 1).count() as i32;
            a.push((cnt, i));
        }

        a.sort_by(|a, b| {
            a.0.cmp(&b.0).then_with(|| a.1.cmp(&b.1))
        });

        let mut ans: Vec<i32> = Vec::with_capacity(k as usize);
        for i in 0..k {
            ans.push(a[i as usize].1 as i32);
        }

        ans
    }
}