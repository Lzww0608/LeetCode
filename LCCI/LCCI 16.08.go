func numberToWords(num int) string {
    if num == 0 {
        return "Zero"
    }
    a := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
    b := []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
    c := []string{"", " ", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

    var partition func(num int) string 
    partition = func(num int) (ans string) {
        if num < 10 {
            ans = a[num]
        } else if num < 20 {
            ans = b[num-10]
        } else if num < 100 {
            ans = c[num/10] + " " + partition(num%10)
        } else if num < 1_000 {
            ans = a[num/100] + " Hundred " + partition(num%100)
        } else if num < 1_000_000 {
            ans = partition(num/1_000) + " Thousand " + partition(num%1_000)
        } else if num < 1_000_000_000 {
            ans = partition(num/1_000_000)  + " Million " + partition(num%1_000_000)
        } else {
            ans = partition(num/1_000_000_000)  + " Billion " + partition(num%1_000_000_000)
        }
        return strings.TrimSpace(ans)
    }

    return partition(num)
}