func largestTimeFromDigits(arr []int) (ans string) {
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            for p := 0; p < 4; p++ {
                for q := 0; q < 4; q++ {
                    if i == j || i == p || i == q || j == p || j == q || p == q {
                        continue
                    }
                    if arr[i] > 2 || arr[p] > 5 {
                        continue
                    }
                    if arr[i] == 2 && arr[j] > 3 {
                        continue
                    }
                    tmp := string([]byte{byte('0' + arr[i]), byte('0' + arr[j]), ':', byte('0' + arr[p]), byte('0' + arr[q])})
                    ans = max(ans, tmp)
                }
            }
        }
    }

    return 
}