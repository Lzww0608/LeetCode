func maximumEvenSplit(finalSum int64) (ans []int64) {
    if finalSum & 1 == 1 {
        return
    }

    sum := int64(0)
    for t := int64(2); sum < finalSum; t += 2 {
        if sum + t <= finalSum {
            sum += t
            ans = append(ans, t)
        } else {
            ans[len(ans)-1] += finalSum - sum
            break
        }
    }

    return
}
