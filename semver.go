// Package semver does minimal major.minor.patch parse & compare, dependency-free.
package semver

import (
	"strconv"
	"strings"
)

func parse(v string) [3]int {
	v = strings.TrimPrefix(v, "v")
	v = strings.SplitN(v, "-", 2)[0]
	v = strings.SplitN(v, "+", 2)[0]
	var out [3]int
	for i, p := range strings.SplitN(v, ".", 3) {
		if i > 2 {
			break
		}
		out[i], _ = strconv.Atoi(p)
	}
	return out
}

// Compare returns -1, 0, or 1.
func Compare(a, b string) int {
	pa, pb := parse(a), parse(b)
	for i := 0; i < 3; i++ {
		if pa[i] != pb[i] {
			if pa[i] < pb[i] {
				return -1
			}
			return 1
		}
	}
	return 0
}
