impl Solution {
    pub fn count_components(n: i32, edges: Vec<Vec<i32>>) -> i32 {
        let mut fa = (0..n).collect::<Vec<i32>>();

        fn find(x: i32, fa: &mut Vec<i32>) -> i32 {
            if fa[x as usize] != x {
                fa[x as usize] = find(fa[x as usize], fa);
            }

            fa[x as usize]
        }

        let mut merge = |x: i32, y: i32, fa: &mut Vec<i32>| {
            let root_x = find(x, fa);
            let root_y = find(y, fa);
            if root_x != root_y {
                fa[root_x as usize] = root_y;
            }
        };

        for edge in edges {
            merge(edge[0], edge[1], &mut fa);
        }

        (0..n).filter(|&i| fa[i as usize] == i).count() as i32
    }
}