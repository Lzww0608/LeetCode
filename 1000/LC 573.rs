impl Solution {
    pub fn min_distance(height: i32, width: i32, tree: Vec<i32>, squirrel: Vec<i32>, nuts: Vec<Vec<i32>>) -> i32 {
        let mut ans = 0;
        
        let getDis = |x1: i32, y1: i32, x2: i32, y2: i32| -> i32 {
            (x1 - x2).abs() + (y1 - y2).abs()
        };

        let mut mn = i32::MAX;
        for v in &nuts {
            let t = getDis(v[0], v[1], tree[0], tree[1]);
            ans += t * 2;
            mn = mn.min(getDis(v[0], v[1], squirrel[0], squirrel[1]) - t);
        }

        ans + mn
    }
}

