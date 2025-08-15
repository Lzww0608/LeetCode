import (
	"strings"
)

func originalDigits(s string) string {
	cnt := make(map[rune]int)
	for _, c := range s {
		cnt[c]++
	}

	ansCnt := [10]int{}

	ansCnt[0] = cnt['z']
	ansCnt[2] = cnt['w']
	ansCnt[4] = cnt['u']
	ansCnt[6] = cnt['x']
	ansCnt[8] = cnt['g']

	ansCnt[3] = cnt['h'] - ansCnt[8]
	ansCnt[5] = cnt['f'] - ansCnt[4]
	ansCnt[7] = cnt['s'] - ansCnt[6]

	ansCnt[1] = cnt['o'] - ansCnt[0] - ansCnt[2] - ansCnt[4]
	ansCnt[9] = cnt['i'] - ansCnt[5] - ansCnt[6] - ansCnt[8]

	var builder strings.Builder
	for i := 0; i <= 9; i++ {
		builder.WriteString(strings.Repeat(string(byte(i+'0')), ansCnt[i]))
	}

	return builder.String()
}