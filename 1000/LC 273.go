func numberToWords(num int) string {
    
    check1 := func(x int) string {
        switch x {
            case 1:
                return "One"
            case 2:
                return "Two"
            case 3:
                return "Three"
            case 4:
                return "Four"
            case 5:
                return "Five"
            case 6:
                return "Six" 
            case 7:
                return "Seven"
            case 8:
                return "Eight"
            case 9:
                return "Nine"
            case 10:
                return "Ten"
        }
        return ""
    }

    check2 := func(x int) string {
        switch x {
            case 2:
                return "Twenty"
            case 3:
                return "Thirty"
            case 4:
                return "Forty"
            case 5:
                return "Fifty"
            case 6:
                return "Sixty" 
            case 7:
                return "Seventy"
            case 8:
                return "Eighty"
            case 9:
                return  "Ninety"
        }
        return ""
    }

    check0 := func(x int) string {
        switch x {
            case 11:
                return "Eleven"
            case 12:
                return "Twelve"
            case 13:
                return "Thirteen"
            case 14:
                return "Fourteen"
            case 15:
                return "Fifteen"
            case 16:
                return "Sixteen" 
            case 17:
                return "Seventeen"
            case 18:
                return "Eighteen"
            case 19:
                return  "Nineteen"
        }
        return ""
    }


    var dfs func(x, d int) string 
    dfs = func(x, d int) string {
        if x == 0 {
            return ""
        }
        res := dfs(x / 1000, d + 1)
        x %= 1000
        builder := strings.Builder{}
        if x >= 100 {
            builder.WriteString(check1(x / 100) + " Hundred ")
            x %= 100
        }
        if x > 10 {
            if x < 20 {
                builder.WriteString(check0(x))
            } else {
                builder.WriteString(check2(x / 10))
                builder.WriteString(" ")
                x %= 10
            }
        }
        if x > 0 {
            builder.WriteString(check1(x))
            builder.WriteString(" ")
        }

        if builder.Len() == 0 {
            return res
        }

        if d == 1 {
            builder.WriteString("Thousand ")
        } else if d == 2 {
            builder.WriteString("Million ")
        } else if d == 3 {
            builder.WriteString("Billion ")
        }

        return res + builder.String()
    }
    ans := dfs(num, 0)
    if ans == "" {
        return "Zero"
    }
    return ans[:len(ans)-1]
}
