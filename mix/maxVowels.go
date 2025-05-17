package mix

import (
	"fmt"
	"strings"
)

func MaxVowels(s string, k int) int {
	mp := map[string]int{"a": 1, "e": 1, "i": 1, "o": 1, "u": 1}
	arr := strings.Split(s, ",")
	count := 0
	glob := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j < i+k; j++ {
			fmt.Print(string(s[j]))
			if _, ok := mp[string(s[j])]; ok {

				count += 1
			}

		}

	}
	return glob
}
